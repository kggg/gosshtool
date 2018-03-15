package controllers

import (
	"github.com/astaxie/beego/validation"
	"gosshtool/web/models"
	"gosshtool/web/utils/msgcrypt"
	"gosshtool/web/utils/validate"
	"strconv"
	"strings"
)

type HostController struct {
	BaseController
}

func (c *HostController) Get() {
	hostinfo, err := models.FindAllHostinfo()
	c.CheckErr(err, "get hostinfo error")
	c.Data["Title"] = "主机列表"
	c.Data["Hosts"] = hostinfo
	c.TplName = "host/index.html"
}

func (c *HostController) AddHost() {
	groups, err := models.FindAllGroups()
	c.CheckErr(err, "find groups error")
	if c.isPost() {
		hostname := strings.TrimSpace(c.GetString("hostname"))
		user := strings.TrimSpace(c.GetString("user"))
		ipaddr := strings.TrimSpace(c.GetString("ipaddr"))
		pass := strings.TrimSpace(c.GetString("pass"))
		pass2 := strings.TrimSpace(c.GetString("pass2"))
		port := strings.TrimSpace(c.GetString("port"))
		group := strings.TrimSpace(c.GetString("group"))

		valid := validation.Validation{}
		var hostinfo = validate.HostInfo{hostname, ipaddr, user, pass, pass2, port, group}
		b, err := valid.Valid(&hostinfo)
		if err != nil {
			c.Resp(false, err.Error()+"valid hostinfo error")
		}
		if !b {
			for _, err := range valid.Errors {
				c.Resp(false, err.Key+" validation error: "+err.Message)
			}
		}
		iport, err := strconv.Atoi(port)
		c.CheckErr(err, "port to int error")
		igroup, err := strconv.Atoi(group)
		c.CheckErr(err, "groupid to int error")
		if models.NameExistCheck(hostname) {
			c.Resp(false, "the hostname has been already existing")
		}
		encryptPass, err := msgcrypt.AesEncrypt(pass)
		c.CheckErr(err, "encrypt pass error")
		_, err = models.AddHost(hostname, ipaddr, user, encryptPass, iport, igroup)
		if err == nil {
			c.Resp(true, "添加主机成功")
		} else {
			c.Resp(false, "添加主机失败")
		}

	}
	c.Data["Title"] = "新增主机"
	c.Data["Groups"] = groups
	c.TplName = "host/add.html"
}

func (c *HostController) EditHost() {
	id := c.Ctx.Input.Param(":id")
	bid, err := strconv.Atoi(id)
	c.CheckErr(err, "delete id error with a to i")
	hostinfo, err := models.FindHostById(bid)
	c.CheckErr(err, "find hostinfo by id error")
	groups, err := models.FindAllGroups()
	c.CheckErr(err, "find groups error")
	if c.isPost() {
		hostname := strings.TrimSpace(c.GetString("hostname"))
		user := strings.TrimSpace(c.GetString("user"))
		ipaddr := strings.TrimSpace(c.GetString("ipaddr"))
		port := strings.TrimSpace(c.GetString("port"))
		group := strings.TrimSpace(c.GetString("group"))

		valid := validation.Validation{}
		var hostinfo = validate.HostEdit{hostname, ipaddr, user, port, group}
		b, err := valid.Valid(&hostinfo)
		if err != nil {
			c.Resp(false, err.Error()+"valid hostinfo error")
		}
		if !b {
			for _, err := range valid.Errors {
				c.Resp(false, err.Key+" validation error: "+err.Message)
			}
		}
		iport, err := strconv.Atoi(port)
		c.CheckErr(err, "port to int error")
		igroup, err := strconv.Atoi(group)
		c.CheckErr(err, "groupid to int error")
		_, err = models.EditHost(hostname, ipaddr, user, iport, igroup, bid)
		if err == nil {
			c.Resp(true, "编辑主机成功")
		} else {
			c.Resp(false, "修改失败")
		}
	}
	c.Data["Title"] = "新增主机"
	c.Data["Groups"] = groups
	c.Data["Hosts"] = hostinfo
	c.TplName = "host/edit.html"
}

func (c *HostController) DelHost() {
	id := c.Ctx.Input.Param(":id")
	bid, err := strconv.Atoi(id)
	c.CheckErr(err, "delete id error with a to i")
	_, err = models.DeleteHost(bid)
	if err == nil {
		c.Redirect("/host", 302)
	}

}
