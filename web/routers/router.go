package routers

import (
	"github.com/astaxie/beego"
	"gosshtool/web/controllers"
)

func init() {
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LoginController{}, "*:Logout")
	beego.Router("/", &controllers.MainController{})

	beego.Router("/user", &controllers.UserController{})
	beego.Router("/user/add", &controllers.UserController{}, "*:Add")
	beego.Router("/user/edit/:id([0-9]+)", &controllers.UserController{}, "get,post:Edit")
	beego.Router("/user/delete/:id([0-9]+)", &controllers.UserController{}, "*:Delete")

	beego.Router("/host", &controllers.HostController{}, "get,post:Get")
	beego.Router("/host/add", &controllers.HostController{}, "get,post:AddHost")
	beego.Router("/host/delete/:id([0-9]+)", &controllers.HostController{}, "post:DelHost")
	beego.Router("/host/command/:hostname", &controllers.CommandController{}, "get,post:Execute")

	beego.Router("/group", &controllers.GroupController{}, "get,post:Get")
	beego.Router("/group/add", &controllers.GroupController{}, "get,post:AddGroups")
	beego.Router("/group/delete/:id([0-9]+)", &controllers.GroupController{}, "post:DeleteGroups")
	beego.Router("/group/command/:group", &controllers.GroupController{}, "get,post:Execute")
}
