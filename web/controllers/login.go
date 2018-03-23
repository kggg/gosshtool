package controllers

import (
	"github.com/gogather/com"
	"gosshtool/web/models"
)

type LoginController struct {
	BaseController
}

type Login struct {
	name     string
	password string
}

func (c *LoginController) Prepare() {
	c.EnableXSRF = false
}

func (c *LoginController) Get() {
	c.TplName = "login/index.html"
}

func (c *LoginController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	if username == "" || password == "" {
		c.Resp(false, "用户和密码不能为空")
	}
	//need to verify with mysql user table
	user, err := models.FindUserByName(username)
	if err != nil {
		c.Resp(false, "用户不存在")

	} else {
		pass := com.Md5(password)
		if pass == user.Pass {

			c.SetSession("username", username)
			c.SetSession("userid", user.Id)
			c.Resp(true, "success")
		} else {
			c.Resp(false, "passwd invalid")
		}
	}
}

func (this *LoginController) Logout() {
	this.SetSession("username", nil)
	this.DelSession("username")
	this.Ctx.Redirect(302, "/login")
}
