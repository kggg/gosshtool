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

// 重定向
func (self *BaseController) redirect(url string) {
	self.Redirect(url, 302)
	self.StopRun()
}

//ajax返回 列表
func (self *BaseController) ajaxList(msg interface{}, msgno int, count int64, data interface{}) {
	out := make(map[string]interface{})
	out["code"] = msgno
	out["msg"] = msg
	out["count"] = count
	out["data"] = data
	self.Data["json"] = out
	self.ServeJSON()
	self.StopRun()
}

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	c.Data["Title"] = "主页"
	c.TplName = "index.tpl"
}
