package controllers

import (
	"github.com/astaxie/beego/validation"
	"github.com/gogather/com"
	"gosshtool/web/models"
	"gosshtool/web/utils/validate"
	"strconv"
	"strings"
)

type UserController struct {
	FrontController
}

func (c *UserController) Get() {
	users, err := models.FindAllUser()
	c.CheckErr(err, "find all user info error")

	c.Data["Users"] = users
	c.TplName = "user/index.html"
}

func (c *UserController) Add() {
	if c.isPost() {
		username := strings.TrimSpace(c.GetString("username"))
		password := strings.TrimSpace(c.GetString("password"))
		password2 := strings.TrimSpace(c.GetString("password2"))
		email := strings.TrimSpace(c.GetString("email"))
		valid := validation.Validation{}
		var userinfo = validate.UserAdd{username, password, password2, email}
		b, err := valid.Valid(&userinfo)
		if err != nil {
			c.Resp(false, err.Error()+"valid userinfo input error")
		}
		if !b {
			for _, err := range valid.Errors {
				c.Resp(false, err.Key+"; validation error: "+err.Message)
			}
		}
		check := models.UserExistCheck(username)
		if check {
			c.Resp(false, "用户已经存在")
		}
		pass := com.Md5(password)
		_, err = models.AddUser(username, pass, email)
		if err != nil {
			c.Resp(false, "添加用户信息入库失败")
		}
		c.Resp(true, "添加用户入库成功")

	}
	c.TplName = "user/add.html"
}

func (c *UserController) Edit() {
	id := c.Ctx.Input.Param(":id")
	bid, err := strconv.Atoi(id)
	rights, err := models.FindAllRights()
	c.CheckErr(err, "Find all rights info error")
	userinfo, err := models.FindUserById(bid)
	c.CheckErr(err, "edit userid error with a to i")
	if c.isPost() {
		email := strings.TrimSpace(c.GetString("email"))
		right := strings.TrimSpace(c.GetString("rights"))
		valid := validation.Validation{}
		var userinfo = validate.UserEdit{email, right}
		b, err := valid.Valid(&userinfo)
		if err != nil {
			c.Resp(false, err.Error()+"valid userinfo input error")
		}
		if !b {
			for _, err := range valid.Errors {
				c.Resp(false, err.Key+"; validation error: "+err.Message)
			}
		}
		iright, err := strconv.Atoi(right)
		c.CheckErr(err, "right atoi error")
		_, err = models.EditUser(email, iright, bid)
		if err != nil {
			c.Resp(false, "添加用户信息入库失败")
		}
		c.Resp(true, "添加用户入库成功")
	}
	c.Data["Userinfo"] = userinfo
	c.Data["Rights"] = rights
	c.TplName = "user/edit.html"
	c.Layout = "layout/layer.tpl"
}

func (c *UserController) Delete() {
	id := c.Ctx.Input.Param(":id")
	bid, err := strconv.Atoi(id)
	c.CheckErr(err, "delete id error with a to i")
	_, err = models.DeleteUser(bid)
	if err == nil {
		c.Resp(true, "success")
	}
}
