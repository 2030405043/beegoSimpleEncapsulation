package main

import (
	_ "1225/models"
	_ "1225/routers"
	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego"
)

func main() {

	o := orm.NewOrm()
	o.Using("default")

	beego.Run()
}
