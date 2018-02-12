package routers

import (
	"github.com/astaxie/beego"
	"gosshtool/web/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/host", &controllers.HostController{})
	beego.Router("/host/add", &controllers.HostController{}, "get,post:AddHost")
	beego.Router("/host/delete/:id([0-9]+)", &controllers.HostController{}, "post:DelHost")
	beego.Router("/command/:hostname", &controllers.CommandController{}, "get,post:Execute")
}
