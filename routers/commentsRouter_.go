package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["1225/controllers:UserController"] = append(beego.GlobalControllerRouter["1225/controllers:UserController"],
        beego.ControllerComments{
            Method: "Test1",
            Router: `/:id/:name/:age`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["1225/controllers:UserController"] = append(beego.GlobalControllerRouter["1225/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAllUser",
            Router: `/:id/:nihao`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["1225/controllers:UserController"] = append(beego.GlobalControllerRouter["1225/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetMessages",
            Router: `/messages`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
