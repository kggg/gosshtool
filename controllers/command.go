package controllers

import (
	"gosshtool/models"
	"gosshtool/utils/msgcrypt"
	"gosshtool/utils/sshclient"
	"strings"
)

type CommandController struct {
	FrontController
}

func (c *CommandController) Execute() {
	host := c.Ctx.Input.Param(":hostname")
	hostname, err := models.FindHostByName(host)
	c.checkerr(err, "get host by name error")

	if c.isPost() {
		cc := c.GetString("command")
		if cc == "" {
			c.Response(false, "the command empty", "")
		}
		if strings.Contains(cc, ">") {
			c.Response(false, "命令行中不能包含重写向符号", "")
		}
		decryptPass, err := msgcrypt.AesDecrypt(hostname.Pass)
		c.checkerr(err, "decrypt pass error")
		var sskey bool
		if hostname.Skey == 1 {
			sskey = true
		} else {
			sskey = false
		}
		client, err := sshclient.NewClient(hostname.Ip, hostname.User, decryptPass, hostname.Port, sskey, hostname.Name)
		c.checkerr(err, "ssh client connect error")
		result, err := client.Run(cc)
		c.checkerr(err, "execute remote cmd error")
		c.Response(true, "success", string(result))
		/*
			var cmd Cmd
			cmd.Result = string(result)
			cmd.Err = err
			c.Data["json"] = &cmd
			c.ServeJSON()
		*/
	}
	c.Data["Title"] = "主机页面"
	c.Data["Hosts"] = hostname
	c.TplName = "host/command.html"

}

type Cmd struct {
	Result string
	Err    error
}
