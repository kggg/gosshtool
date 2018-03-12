package controllers

import (
	"github.com/astaxie/beego/validation"
	"gosshtool/web/models"
	"gosshtool/web/utils/msgcrypt"
	"gosshtool/web/utils/validate"
	"strconv"
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
	if c.isPost() {
		hostname := c.GetString("hostname")
		user := c.GetString("user")
		ipaddr := c.GetString("ipaddr")
		pass := c.GetString("pass")
		pass2 := c.GetString("pass2")
		port := c.GetString("port")
		group := c.GetString("group")

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
		if models.NameExistCheck(hostname) {
			c.Resp(false, "the hostname has been already existing")
		}
		encryptPass, err := msgcrypt.AesEncrypt(pass)
		c.CheckErr(err, "encrypt pass error")
		_, err = models.AddHost(hostname, ipaddr, user, encryptPass, iport, group)
		if err == nil {
			c.Resp(true, "添加主机成功")
		} else {
			result := make(map[string]interface{})
			//beego.Error("Insert into mysql error: ", err.Error())
			result["status"] = "false"
			result["message"] = "数据创建失败"
			result["debug"] = err.Error()
			c.Data["json"] = result
			c.ServeJSON()
			c.StopRun()
		}

	}
	c.Data["Title"] = "新增主机"
	c.TplName = "host/add.html"
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
