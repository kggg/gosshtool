package controllers

import (
	"gosshtool/models"
	"gosshtool/utils/msgcrypt"
	"gosshtool/utils/sshclient"
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
		var sskey bool
		if hostname.Skey == 1 {
			sskey = true
		} else {
			sskey = false
		}
		client, err := sshclient.NewClient(hostname.Ip, hostname.User, decryptPass, hostname.Port, sskey, hostname.Name)
		c.CheckErr(err, "ssh client connect error")
		result, err := client.Run(cc)
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
