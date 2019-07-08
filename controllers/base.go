package controllers

import (
	"strings"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Response(status bool, str string, data interface{}) {
	this.Data["json"] = &map[string]interface{}{"status": status, "info": str, "data": data}
	this.ServeJSON()
	this.StopRun()
}

func (this *BaseController) checkerr(err error, str string) {
	if err != nil {
		this.Response(false, str, err.Error())
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
