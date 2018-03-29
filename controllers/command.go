package controllers

import (
	"gosshtool/lib/sshclient"
	"gosshtool/models"
	"gosshtool/utils/msgcrypt"
)

type CommandController struct {
	FrontController
}

func (c *CommandController) Execute() {
	host := c.Ctx.Input.Param(":hostname")
	hostname, err := models.FindHostByName(host)
	c.CheckErr(err, "get host by name error")

	if c.isPost() {
		cc := c.GetString("command")
		if cc == "" {
			c.Resp(false, "the command empty")
		}
		decryptPass, err := msgcrypt.AesDecrypt(hostname.Pass)
		c.CheckErr(err, "decrypt pass error")
		client := sshclient.New(hostname.Ip, hostname.User, decryptPass, hostname.Port, hostname.Name)
		result, err := client.Exec(cc)
		c.CheckErr(err, "execute remote cmd error")
		var cmd Cmd
		cmd.Result = string(result)
		cmd.Err = err
		c.Data["json"] = &cmd
		c.ServeJSON()
	}
	c.Data["Title"] = "主机页面"
	c.Data["Hosts"] = hostname
	c.TplName = "host/command.html"

}

type Cmd struct {
	Result string
	Err    error
}
