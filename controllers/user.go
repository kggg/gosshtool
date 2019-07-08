package controllers

import (
	"gosshtool/models"
	"gosshtool/utils/validate"
	"strconv"
	"strings"

	"github.com/astaxie/beego/validation"
	"github.com/gogather/com"
)

type UserController struct {
	FrontController
}

func (c *UserController) Get() {
	users, err := models.FindAllUser()
	c.checkerr(err, "find all user info error")

	c.Data["Title"] = "用户列表"
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
			c.Response(false, err.Error()+"valid userinfo input error", "")
		}
		if !b {
			for _, err := range valid.Errors {
				c.Response(false, err.Key+"; validation error: "+err.Message, "")
			}
		}
		check := models.UserExistCheck(username)
		if check {
			c.Response(false, "用户已经存在", "")
		}
		pass := com.Md5(password)
		_, err = models.AddUser(username, pass, email)
		if err != nil {
			c.Response(false, "添加用户信息入库失败", "")
		}
		c.Response(true, "添加用户入库成功", "")

	}
	c.Data["Title"] = "新增用户"
	c.TplName = "user/add.html"
}

func (c *UserController) Edit() {
	id := c.Ctx.Input.Param(":id")
	bid, err := strconv.Atoi(id)
	rights, err := models.FindAllRights()
	c.checkerr(err, "Find all rights info error")
	userinfo, err := models.FindUserById(bid)
	c.checkerr(err, "edit userid error with a to i")
	if c.isPost() {
		email := strings.TrimSpace(c.GetString("email"))
		right := strings.TrimSpace(c.GetString("rights"))
		valid := validation.Validation{}
		var userinfo = validate.UserEdit{email, right}
		b, err := valid.Valid(&userinfo)
		if err != nil {
			c.Response(false, err.Error()+"valid userinfo input error", "")
		}
		if !b {
			for _, err := range valid.Errors {
				c.Response(false, err.Key+"; validation error: "+err.Message, "")
			}
		}
		iright, err := strconv.Atoi(right)
		c.checkerr(err, "right atoi error")
		_, err = models.EditUser(email, iright, bid)
		if err != nil {
			c.Response(false, "添加用户信息入库失败", "")
		}
		c.Response(true, "添加用户入库成功", "")
	}
	c.Data["Userinfo"] = userinfo
	c.Data["Rights"] = rights
	c.TplName = "user/edit.html"
	c.Layout = "layout/layer.tpl"
}

func (c *UserController) Delete() {
	id := c.Ctx.Input.Param(":id")
	bid, err := strconv.Atoi(id)
	c.checkerr(err, "delete id error with a to i")
	_, err = models.DeleteUser(bid)
	if err == nil {
		c.Response(true, "success", "")
	}
}

func (c *UserController) ChangePass() {
	id := c.Ctx.Input.Param(":id")
	bid, err := strconv.Atoi(id)
	c.checkerr(err, "get change user id error with a to i")
	if c.isPost() {
		oldpass := strings.TrimSpace(c.GetString("oldpass"))
		newpass := strings.TrimSpace(c.GetString("newpass"))
		newpass2 := strings.TrimSpace(c.GetString("newpass2"))
		if oldpass == "" || newpass == "" || newpass2 == "" {
			c.Response(false, "密码不能为空", "")
		}
		if len(oldpass) > 50 || len(newpass) > 50 || len(newpass2) > 50 {
			c.Response(false, "密码长度不能超过50字符", "")
		}
		if newpass != newpass2 {
			c.Response(false, "新密码与确认密码不一致", "")
		}
		userinfo, err := models.FindUserById(bid)
		c.checkerr(err, "get user info by id error")
		encryptoldpass := com.Md5(oldpass)
		if encryptoldpass != userinfo.Pass {
			c.Response(false, "旧密码不正确", "")
		}
		encryptPass := com.Md5(newpass)
		_, err = models.ChangeUserPass(encryptPass, bid)
		if err != nil {
			c.Response(false, "修改密码入库失败", "")
		}
		c.Response(true, "修改密码成功", "")
	}
	c.Data["Id"] = bid
	c.TplName = "user/changepass.html"
	c.Layout = "layout/layer.tpl"
}
