package controllers

type MainController struct {
	FrontController
}

func (c *MainController) Get() {
	c.Data["Title"] = "主页"
	c.TplName = "index.tpl"
}
