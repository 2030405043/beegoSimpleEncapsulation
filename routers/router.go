package routers

import (
	"1225/controllers"
	"fmt"
	"github.com/astaxie/beego"
)

func init() {

	fmt.Println("22222222222222222222")
	beego.Router("/", &controllers.MainController{})

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user", beego.NSInclude(&controllers.UserController{})),
		beego.NSNamespace("/active", beego.NSInclude(&controllers.UserController{})))

	beego.AddNamespace(ns)

}
