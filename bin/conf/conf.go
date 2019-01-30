package conf

import (
	"github.com/BurntSushi/toml"
)

const Filepath = "C:/Users/home/go/src/gosshtool/bin/conf/conf.ini"

type Mysql struct {
	Username string
	Password string
	Host     string
	Dbname   string
	Ipaddr   string
	Port     int64
}

type File struct {
	Hostpath  string
	Grouppath string
}

type Config struct {
	Mode     string
	Mysqldb  Mysql `toml:"mysql"`
	Fileinfo File  `toml:"file"`
	Filepath string
}

func Newconfig(filename string) (*Config, error) {
	config := &Config{}
	config.Filepath = filename
	_, err := toml.DecodeFile(config.Filepath, config)
	if err != nil {
		return &Config{}, err
	}
	return config, nil
}

func (c *Config) Getmode() string {
	return c.Mode
}

func (c *Config) GetDBparams() Mysql {
	return c.Mysqldb
}

func (c *Config) GetFileinfo() File {
	return c.Fileinfo
}
