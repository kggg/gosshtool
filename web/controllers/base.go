package controllers

import (
	"github.com/astaxie/beego"
	//"testblog/models"
	"strings"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare() {
	/*
			username := c.GetSession("username")
			if username == false || username == nil {
				c.Redirect("/login", 302)
			}
		c.Data["Username"] = username
	*/
	c.Layout = "layout/main.tpl"

}

func (this *BaseController) Resp(status bool, str interface{}) {
	this.Data["json"] = &map[string]interface{}{"status": status, "info": str}
	this.ServeJSON()
	this.StopRun()
}

func (this *BaseController) CheckErr(err error, str string) {
	if err != nil {
		this.Resp(false, str)
	}
}

func (this *BaseController) isPost() bool {
	return this.Ctx.Request.Method == "POST"
}

func (this *BaseController) getClientIp() string {
	s := strings.Split(this.Ctx.Request.RemoteAddr, ":")
	return s[0]
}

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	c.Data["Title"] = "主页"
	c.TplName = "index.tpl"
}
