package controllers

import (
	"github.com/astaxie/beego/validation"
	"gosshtool/lib/sshclient"
	"gosshtool/models"
	"gosshtool/utils/msgcrypt"
	"gosshtool/utils/validate"
	"log"
	"strconv"
	"strings"
	"sync"
)

type GroupController struct {
	FrontController
}

func (c *GroupController) Get() {
	groups, err := models.FindAllGroups()
	c.CheckErr(err, "get groups info error")
	if c.isPost() {
		group := strings.TrimSpace(c.GetString("gname"))
		info := c.GetString("info")
		gid, _ := c.GetInt("id")

		valid := validation.Validation{}
		var groupv = validate.Groups{group, info}
		b, err := valid.Valid(&groupv)
		if err != nil {
			c.Resp(false, err.Error()+"valid groups info error")
		}
		if !b {
			for _, err := range valid.Errors {
				c.Resp(false, err.Key+" validation error: "+err.Message)
			}
		}
		_, err = models.EditGroups(group, info, gid)
		if err == nil {
			c.Resp(true, "编辑组名成功")
		} else {
			c.Resp(false, "编辑组名失败")
		}
	}
	c.Data["Title"] = "组别列表"
	c.Data["Groups"] = groups
	c.TplName = "group/index.html"
}

func (c *GroupController) AddGroups() {

	if c.isPost() {
		group := strings.TrimSpace(c.GetString("gname"))
		info := c.GetString("info")

		valid := validation.Validation{}
		var groupv = validate.Groups{group, info}
		b, err := valid.Valid(&groupv)
		if err != nil {
			c.Resp(false, err.Error()+"valid groups info error")
		}
		if !b {
			for _, err := range valid.Errors {
				c.Resp(false, err.Key+" validation error: "+err.Message)
			}
		}
		if models.GroupsExistCheck(group) {
			c.Resp(false, "the groupname has been already existing")
		}
		_, err = models.AddGroups(group, info)
		if err == nil {
			c.Resp(true, "添加组名成功")
		} else {
			c.Resp(false, "添加组名失败")
		}
	}

	c.Data["Title"] = "新增组别"
	c.TplName = "group/add.html"
}

/*
func (c *GroupController) EditGroups() {
	id := c.Ctx.Input.Param(":id")
	bid, err := strconv.Atoi(id)
	c.CheckErr(err, "delete id error with a to i")
	groupinfo, err := models.FindGroupsById(bid)
	c.CheckErr(err, "get groupid info error")
	if c.isPost() {
		group := strings.TrimSpace(c.GetString("gname"))
		info := c.GetString("info")

		valid := validation.Validation{}
		var groupv = validate.Groups{group, info}
		b, err := valid.Valid(&groupv)
		if err != nil {
			c.Resp(false, err.Error()+"valid groups info error")
		}
		if !b {
			for _, err := range valid.Errors {
				c.Resp(false, err.Key+" validation error: "+err.Message)
			}
		}
		_, err = models.EditGroups(group, info, bid)
		if err == nil {
			c.Resp(true, "编辑组名成功")
		} else {
			c.Resp(false, "编辑组名失败")
		}
	}
	c.Data["Title"] = "编辑组别"
	c.Data["Groups"] = groupinfo
	c.TplName = "group/edit.html"
}
*/

func (c *GroupController) DeleteGroups() {
	id := c.Ctx.Input.Param(":id")
	bid, err := strconv.Atoi(id)
	c.CheckErr(err, "delete id error with a to i")
	_, err = models.DeleteGroups(bid)
	if err == nil {
		c.Redirect("/group", 302)
	}
}

func (c *GroupController) Execute() {
	groupname := c.Ctx.Input.Param(":group")
	groups, err := models.FindHostByGroupname(groupname)
	c.CheckErr(err, "get group by name error")

	if c.isPost() {
		cc := c.GetString("command")
		if cc == "" {
			c.Resp(false, "the command empty")
		}
		var result = make([]map[string]interface{}, 0, len(groups))
		var wg sync.WaitGroup
		for _, hostname := range groups {
			wg.Add(1)
			go func(ip, user, pass string, port int, name string) {
				out := make(map[string]interface{})
				decryptPass, err := msgcrypt.AesDecrypt(pass)
				c.CheckErr(err, "decrypt pass error")
				client := sshclient.New(ip, user, decryptPass, port, name)
				res, err := client.Exec(cc)
				if err != nil {
					log.Println("Error: execute remote cmd, ", err)
				}
				out["ip"] = ip
				out["hostname"] = name
				out["res"] = string(res)
				result = append(result, out)
				wg.Done()
			}(hostname.Ip, hostname.User, hostname.Pass, hostname.Port, hostname.Name)
		}
		wg.Wait()
		c.Data["json"] = result
		c.ServeJSON()
	}
	c.Data["Groups"] = groups
	c.Data["Gname"] = groupname
	c.Data["Title"] = "远程命令"
	c.TplName = "group/command.html"
}
