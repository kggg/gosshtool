package controllers

import (
	"gosshtool/lib/sshclient"
	"gosshtool/web/models"
)

type CommandController struct {
	BaseController
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
		client := sshclient.New(hostname.Ip, hostname.User, hostname.Pass, hostname.Port, hostname.Name)
		result, err := client.Exec(cc)
		var cmd Cmd
		cmd.Result = result
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
