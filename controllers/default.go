package controllers

import (
	"1225/models"
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type UserController struct {
	beego.Controller
}



// @router /messages [get]
func (this *UserController) GetMessages() {
	models.GetMessages()
}


// @router /:id/:nihao [get]
func (this *UserController) GetAllUser() {
	resp := this.Ctx.ResponseWriter
	//req := this.Ctx.Request

	id := this.Ctx.Input.Param(":id")

	fmt.Println(id)

	fmt.Println(resp.Started)

	resp.Write([]byte("你好.."))

	query := this.Ctx.Input.Query("nihao")

	fmt.Println(query)
}

// @router /:id/:name/:age [get]
func (this *UserController) Test1() {

	fmt.Println("Hello World...")
	age, _ := this.GetInt("age", 0)
	name := this.GetString("name", "")
	id := this.GetString("id")

	fmt.Println(id, name, age)
	//http://localhost:7788/v1/user/123/lisi/12
	//输出 123 lisi 12


	idQ :=this.Ctx.Input.Query("id")
	nameQ := this.Ctx.Input.Query("name")
	ageQ:= this.Ctx.Input.Query("age")

	fmt.Println(idQ,nameQ,ageQ)
	//http://localhost:7788/v1/user/123/lisi/12
	//输出 123 lisi 12


	idP := this.Ctx.Input.Param("id")
	nameP:=this.Ctx.Input.Param("name")
	ageP:=this.Ctx.Input.Param("age")

	fmt.Println(idP,nameP,ageP)
}
