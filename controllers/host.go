package controllers

import (
	"gosshtool/models"
	"gosshtool/utils/msgcrypt"
	"gosshtool/utils/validate"
	"strconv"
	"strings"

	"github.com/astaxie/beego/validation"
)

type HostController struct {
	FrontController
}

func (c *HostController) Get() {
	hostinfo, err := models.FindAllHostinfo()
	c.checkerr(err, "get hostinfo error")
	groups, err := models.FindAllGroups()
	c.checkerr(err, "find groups error")
	if c.isPost() {
		hid, _ := c.GetInt("id")
		hostname := strings.TrimSpace(c.GetString("hostname"))
		user := strings.TrimSpace(c.GetString("user"))
		ipaddr := strings.TrimSpace(c.GetString("ipaddr"))
		port := strings.TrimSpace(c.GetString("port"))
		group := strings.TrimSpace(c.GetString("group"))

		valid := validation.Validation{}
		var hostinfo = validate.HostEdit{hostname, ipaddr, user, port, group}
		b, err := valid.Valid(&hostinfo)
		if err != nil {
			c.Response(false, err.Error()+"valid hostinfo error", "")
		}
		if !b {
			for _, err := range valid.Errors {
				c.Response(false, err.Key+" validation error: "+err.Message, "")
			}
		}
		iport, err := strconv.Atoi(port)
		c.checkerr(err, "port to int error")
		igroup, err := strconv.Atoi(group)
		c.checkerr(err, "groupid to int error")
		_, err = models.EditHost(hostname, ipaddr, user, iport, igroup, hid)
		if err == nil {
			c.Response(true, "编辑主机成功", "")
		} else {
			c.Response(false, "修改失败", "")
		}
	}
	c.Data["Title"] = "主机列表"
	c.Data["Groups"] = groups
	c.Data["Hosts"] = hostinfo
	c.TplName = "host/index.html"
}

func (c *HostController) AddHost() {
	groups, err := models.FindAllGroups()
	c.checkerr(err, "find groups error")
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
			c.Response(false, err.Error()+"valid hostinfo error", "")
		}
		if !b {
			for _, err := range valid.Errors {
				c.Response(false, err.Key+" validation error: "+err.Message, "")
			}
		}
		iport, err := strconv.Atoi(port)
		c.checkerr(err, "port to int error")
		igroup, err := strconv.Atoi(group)
		c.checkerr(err, "groupid to int error")
		if models.NameExistCheck(hostname) {
			c.Response(false, "the hostname has been already existing", "")
		}
		encryptPass, err := msgcrypt.AesEncrypt(pass)
		c.checkerr(err, "encrypt pass error")
		_, err = models.AddHost(hostname, ipaddr, user, encryptPass, iport, igroup)
		if err == nil {
			c.Response(true, "添加主机成功", "")
		} else {
			c.Response(false, "添加主机失败", "")
		}

	}
	c.Data["Title"] = "新增主机"
	c.Data["Groups"] = groups
	c.TplName = "host/add.html"
}

/*
func (c *HostController) EditHost() {
	id := c.Ctx.Input.Param(":id")
	bid, err := strconv.Atoi(id)
	c.checkerr(err, "delete id error with a to i")
	hostinfo, err := models.FindHostById(bid)
	c.checkerr(err, "find hostinfo by id error")
	groups, err := models.FindAllGroups()
	c.checkerr(err, "find groups error")
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
			c.Response(false, err.Error()+"valid hostinfo error")
		}
		if !b {
			for _, err := range valid.Errors {
				c.Response(false, err.Key+" validation error: "+err.Message)
			}
		}
		iport, err := strconv.Atoi(port)
		c.checkerr(err, "port to int error")
		igroup, err := strconv.Atoi(group)
		c.checkerr(err, "groupid to int error")
		_, err = models.EditHost(hostname, ipaddr, user, iport, igroup, bid)
		if err == nil {
			c.Response(true, "编辑主机成功")
		} else {
			c.Response(false, "修改失败")
		}
	}
	c.Data["Title"] = "新增主机"
	c.Data["Groups"] = groups
	c.Data["Hosts"] = hostinfo
	c.TplName = "host/edit.html"
}
*/

func (c *HostController) ChangePass() {
	id := c.Ctx.Input.Param(":id")
	bid, err := strconv.Atoi(id)
	c.checkerr(err, "get change host id error with a to i")
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
		hostinfo, err := models.FindHostById(bid)
		c.checkerr(err, "get host info by id error")
		encryptoldpass, cerr := msgcrypt.AesEncrypt(oldpass)
		c.checkerr(cerr, "encrypt oldpass error")
		if encryptoldpass != hostinfo.Pass {
			c.Response(false, "旧密码不正确", "")
		}
		encryptPass, err := msgcrypt.AesEncrypt(newpass)
		c.checkerr(err, "encrypt newpass error")
		_, err = models.ChangeHostPass(encryptPass, bid)
		if err != nil {
			c.Response(false, "修改密码入库失败", "")
		}
		c.Response(true, "修改密码成功", "")
	}
	c.Data["Id"] = bid
	c.TplName = "host/changepass.html"
	c.Layout = "layout/layer.tpl"
}

func (c *HostController) DelHost() {
	id := c.Ctx.Input.Param(":id")
	bid, err := strconv.Atoi(id)
	c.checkerr(err, "delete id error with a to i")
	_, err = models.DeleteHost(bid)
	if err == nil {
		c.Redirect("/host", 302)
	}

}
