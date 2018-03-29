package controllers

type FrontController struct {
	BaseController
}

func (c *FrontController) Prepare() {
	username := c.GetSession("username")
	if username == nil {
		c.redirect("/login")
	}
	userid := c.GetSession("userid")
	c.Data["Userid"] = userid
	c.Data["Username"] = username
	//c.Layout = "layout/layui.tpl"
	c.Layout = "layout/main.tpl"

}
