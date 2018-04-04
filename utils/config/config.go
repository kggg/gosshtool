package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"gosshtool/utils/db"
	"gosshtool/utils/msgcrypt"
	"os"
)

type AppConfig struct {
	Hostmode string `toml:"hostmode"`
	Workdir  string `toml:"workdir"`
	Mysql    Dbinfo
	File     Fileinfo
}

type Dbinfo struct {
	Host   string
	User   string
	Pass   string
	Dbname string
	Port   int
}

type Fileinfo struct {
	Hostpath  string
	Grouppath string
}

func getappConf() (AppConfig, error) {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	confpath := fmt.Sprintf("%s/conf/app.conf", path)
	var appconf AppConfig
	if _, err := toml.DecodeFile(confpath, &appconf); err != nil {
		return AppConfig{}, err
	}

	return appconf, nil
}

func GetHostInfo() (map[string]HostInfo, error) {
	appconf, err := getappConf()
	if err != nil {
		return nil, err
	}
	hostmap := make(map[string]HostInfo)
	if appconf.Hostmode == "file" {
		if appconf.File.Hostpath == "" {
			appconf.File.Hostpath = getHostPath()
		}
		hostinfo, err := GetHostsFromFile(appconf.File.Hostpath)
		if err != nil {
			return nil, err
		}
		for name, v := range hostinfo.Hosts {
			hostmap[name] = v
		}
	}
	if appconf.Hostmode == "mysql" {
		hostmap, err = getHostformMysql(appconf.Mysql.Host, appconf.Mysql.User, appconf.Mysql.Pass, appconf.Mysql.Port, appconf.Mysql.Dbname)
		if err != nil {
			return nil, err
		}

	}
	return hostmap, nil
}

type TomlConfig struct {
	Title string
	Hosts map[string]HostInfo
}

type HostInfo struct {
	Ip   string
	User string
	Pass string
	Port int
}

func getHostPath() string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	cfgfile := fmt.Sprintf("%s/conf/hosts.conf", path)
	return cfgfile
}

func GetHostsFromFile(cfgfile string) (TomlConfig, error) {
	var hostinfo TomlConfig

	if _, err := toml.DecodeFile(cfgfile, &hostinfo); err != nil {
		return TomlConfig{}, err
	}
	return hostinfo, nil
}

func getHostformMysql(host, user, pass string, port int, dbname string) (map[string]HostInfo, error) {
	dbcon, err := db.New(host, user, pass, port, dbname)
	if err != nil {
		return nil, err
	}
	hostmap := make(map[string]HostInfo)
	sqli := "select * from hostinfo"
	hostinfo, err := db.FetchRows(dbcon, sqli)
	for _, v := range hostinfo {
		dpass, err := msgcrypt.AesDecrypt(v["pass"])
		if err != nil {
			return nil, err
		}
		hostmap[v["name"]] = HostInfo{Ip: v["ip"], User: v["user"], Pass: dpass, Port: 22}
	}
	return hostmap, nil
}
