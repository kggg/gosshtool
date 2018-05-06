package command

import (
	"flag"
	"fmt"
	"gosshtool/models"
	"os"
	//"regexp"
	"strings"
)

type Command struct {
	Host   []models.Hostinfo
	Module string
	Action string
}

func (this *Command) AddHost(h string) {
	hosts := strings.Split(h, ",")
	for _, host := range hosts {
		host = strings.TrimSpace(host)
		hostinfo, err := models.FindHostByName(host)
		if err != nil {
			fmt.Errorf("get host error from db: %v\n", err)
		}
		if hostinfo.Name == "" {
			continue
		}

		this.Host = append(this.Host, hostinfo)
	}
	if len(this.Host) < 1 {
		fmt.Printf("No hostname match your input [%s]\n", h)
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
	}
}

func (this *Command) AddModule(c string) {
	this.Module = c
}

/*
func (this *Command) AddRegular(str string) {
	allhost, err := readHost()
	if err != nil {
		panic(err)
	}
	this.Host = make(map[string]config.HostInfo)
	//for name, v := range allhost.Hosts {
	for name, v := range allhost {
		if str == "*" {
			this.Host[name] = v
		} else {
			match, _ := regexp.MatchString(str, name)
			if match {
				this.Host[name] = v
			}
		}
	}

}
*/

var (
	host   string
	group  string
	reg    string
	module string
)

func init() {
	flag.StringVar(&host, "h", "", "remote hostname")
	flag.StringVar(&group, "g", "", "group name")
	flag.StringVar(&reg, "e", "", "regrex match host name")
	flag.StringVar(&module, "m", "", "module name")
}

func ParseCommand() Command {
	flag.Usage = func() {
		fmt.Printf("Usage: %s host [host|group] options [cmd|copyfile]\n", os.Args[0])
		fmt.Printf("\t  -h : specified a remote host, use , split one or more  host\n")
		fmt.Printf("\t  -g : specified a remote hostgroup\n")
		fmt.Printf("\t  -E : Regrex match a remote host name default\n")
		fmt.Printf("\t  -m : select a module, -m [cmd|sendfile|getfile]\n")
		fmt.Printf("e.g.:   gosshtoll -h steven -m cmd 'uptime'\n")
		os.Exit(0)
	}
	flag.Parse()
	var cmdname Command
	if host != "" {
		cmdname.AddHost(host)
	}
	if group != "" {
		cmdname.AddGroup(group)
	}
	//if reg != "" {
	//	cmdname.AddRegular(reg)
	//}
	if module != "" {
		cmdname.AddModule(module)
		if flag.NArg() > 0 {
			cmdname.Action = strings.TrimSpace(flag.Args()[0])
		} else {
			fmt.Println("The NArg() empty")
			flag.Usage()
		}
	} else {
		fmt.Println("Error: must to specified a module, like as: cmd, sendfile, getfile")
		flag.Usage()
	}
	return cmdname
}
