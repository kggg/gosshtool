package command

import (
	"flag"
	"fmt"
	"gosshtool/bin/conf"
	"gosshtool/models"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

type Command struct {
	Host   []models.Hostinfo
	Module string
	Action string
	Model  string
}

func (this *Command) AddHost(h string) {
	hosts := strings.Split(h, ",")
	for _, host := range hosts {
		host = strings.TrimSpace(host)
		if this.Model == "mysql" {
			hostinfo, err := models.FindHostByName(host)
			if err != nil {
				fmt.Printf("get host by name error from db: %v\n", err)
				os.Exit(1)
			}
			if hostinfo.Name == "" {
				continue
			}
			this.Host = append(this.Host, hostinfo)
		}
		if this.Model == "file" {

		}

	}
	if len(this.Host) < 1 {
		fmt.Printf("No hostname match your input [%s]\n", h)
		os.Exit(1)
	}
}

func (this *Command) AddGroup(g string) {
	groups, err := models.FindHostByGroupname(g)
	if err != nil {
		fmt.Errorf("get host info error from db: %v\n", err)
	}
	for _, v := range groups {
		this.Host = append(this.Host, v.Hostinfo)
	}
	if len(this.Host) < 1 {
		fmt.Printf("No groupname match your input [%s]\n", g)
		os.Exit(1)
	}
}

func (this *Command) AddModule(c string) {
	this.Module = c
}

func (this *Command) AddRegular(str string) {
	allhost, err := models.FindAllHostinfo()
	if err != nil {
		fmt.Println("database error")
		os.Exit(1)
	}
	m, _ := regexp.Compile(str)

	for _, v := range allhost {
		if str == "*" || str == "all" {
			this.Host = append(this.Host, v.Hostinfo)
		} else {
			if m.MatchString(v.Name) {
				this.Host = append(this.Host, v.Hostinfo)
			}
		}
	}
	if len(this.Host) < 1 {
		fmt.Printf("No hostname match your input [%s]\n", str)
		os.Exit(1)
	}

}

func (c *Command) Listhostinfo() {
	allhost, err := models.FindAllHostinfo()
	if err != nil {
		fmt.Println("get all host info error")
		os.Exit(1)
	}
	//fmt.Printf("## IP:port\tHostname\tUser  ##\n")
	color.Cyan("## IP:port\tHostname\tUser  ##\n")
	for _, v := range allhost {
		fmt.Printf("%s:%d\t%s\t%s\n", v.Hostinfo.Ip, v.Hostinfo.Port, v.Hostinfo.Name, v.Hostinfo.User)
	}
	os.Exit(0)
}

var (
	host   string
	group  string
	reg    string
	list   bool
	module string
)

func init() {
	flag.StringVar(&host, "h", "", "remote hostname")
	flag.StringVar(&group, "g", "", "group name")
	flag.StringVar(&reg, "e", "", "regrex match host name")
	flag.BoolVar(&list, "list", false, "list host info")
	flag.StringVar(&module, "m", "", "module name")
}

func Getworkdir() (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return pwd, nil
}

func ParseCommand() *Command {
	flag.Usage = func() {
		fmt.Printf("Usage: %s host [host|group] options [cmd|copyfile]\n", os.Args[0])
		fmt.Printf("\t  -h : specified a remote host, use , split one or more  host\n")
		fmt.Printf("\t  -g : specified a remote hostgroup\n")
		fmt.Printf("\t  -e : Regrex match a remote host name default\n")
		fmt.Printf("\t  -list : list host info\n")
		fmt.Printf("\t  -m : select a module, -m [cmd|copy]\n")
		fmt.Printf("\t\t   copy : [src, dest,mode,force,backup,user,owner]\n\n")
		fmt.Printf("e.g.:   gosshtoll -h steven -m cmd 'uptime'\n")
		os.Exit(0)
	}
	flag.Parse()
	workdir, err := Getworkdir()
	if err != nil {
		log.Fatal(err)
	}
	filename := workdir + "/conf/conf.ini"
	cfg, err := conf.Newconfig(filename)
	if err != nil {
		log.Fatal(err)
	}
	mode := cfg.Getmode()
	cmdname := &Command{}
	cmdname.Model = mode
	if host != "" {
		cmdname.AddHost(host)
	}
	if group != "" {
		cmdname.AddGroup(group)
	}
	if reg != "" {
		cmdname.AddRegular(reg)
	}

	if module != "" {
		cmdname.AddModule(module)
		if flag.NArg() > 0 {
			cmdname.Action = strings.TrimSpace(flag.Args()[0])
		} else {
			fmt.Println("The NArg() empty")
			flag.Usage()
		}
	} else if module == "" && list {
		cmdname.Listhostinfo()
	} else {
		fmt.Println("Error: must to specified a module, like as: cmd, sendfile, getfile")
		flag.Usage()
	}
	return cmdname
}
