package routers

import (
	"github.com/astaxie/beego"
	"gosshtool/web/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/host", &controllers.HostController{})
	beego.Router("/host/add", &controllers.HostController{}, "get,post:AddHost")
	beego.Router("/host/edit/:id([0-9])", &controllers.HostController{}, "get,post:EditHost")
	beego.Router("/host/delete/:id([0-9]+)", &controllers.HostController{}, "post:DelHost")
	beego.Router("/host/command/:hostname", &controllers.CommandController{}, "get,post:Execute")

	beego.Router("/group", &controllers.GroupController{})
	beego.Router("/group/add", &controllers.GroupController{}, "get,post:AddGroups")
	beego.Router("/group/edit/:id([0-9])", &controllers.GroupController{}, "get,post:EditGroups")
	beego.Router("/group/delete/:id([0-9])", &controllers.GroupController{}, "post:DeleteGroups")
	beego.Router("/group/command/:group", &controllers.GroupController{}, "get,post:Execute")
}
