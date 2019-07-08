package controllers

import (
	"gosshtool/models"

	"github.com/gogather/com"
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
		c.Response(false, "用户和密码不能为空", "")
	}
	//need to verify with mysql user table
	user, err := models.FindUserByName(username)
	if err != nil {
		c.Response(false, "用户不存在", "")

	} else {
		pass := com.Md5(password)
		if pass == user.Pass {

			c.SetSession("username", username)
			c.SetSession("userid", user.Id)
			c.Response(true, "success", "")
		} else {
			c.Response(false, "passwd invalid", "")
		}
	}
}

func (this *LoginController) Logout() {
	this.SetSession("username", nil)
	this.DelSession("username")
	this.Ctx.Redirect(302, "/login")
}
