package controllers

import (
	"gosshtool/web/models"
	"strconv"
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

	}
	c.TplName = "user/add.html"
}

func (c *UserController) Edit() {
	if c.isPost() {

	}
	c.TplName = "user/edit.html"
}

func (c *UserController) Delete() {
	id := c.Ctx.Input.Param(":id")
	bid, err := strconv.Atoi(id)
	c.CheckErr(err, "delete id error with a to i")
	_, err = models.DeleteUser(bid)
	if err == nil {
		c.Redirect("/host", 302)
	}
}
