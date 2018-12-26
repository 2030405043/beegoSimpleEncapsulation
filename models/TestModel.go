package models

import (
	"1225/DBconf"
	"1225/entities"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//数据库信息配置
func init() {

	dBconf := DBconf.GetDBconf()

	orm.RegisterDriver(dBconf.DriverName, orm.DRMySQL)

	RegisterModel()

	dataSource := dBconf.Account + ":" + dBconf.Password + "@tcp(" + dBconf.Ip + ":" + dBconf.Port + ")/" +
		"" + dBconf.DB + "?" + dBconf.EnCoding

	fmt.Println(dataSource)
	err := orm.RegisterDataBase(`default`, dBconf.DriverName, dataSource, dBconf.MaxIdle, dBconf.MaxConn)
	if err != nil {
		//panic("连接数据库失败....",err.Error())
		fmt.Println(err.Error())
	}
}
func GetOmr() orm.Ormer {
	orm.Debug = true
	ormer := orm.NewOrm()
	ormer.Using("default")
	return ormer
}
func GetMessages() {
	ormer := GetOmr()

	var messages []entities.Message

	sqls := "select * from message"
	i, err := ormer.Raw(sqls).QueryRows(&messages)
	if err != nil {
		fmt.Println("查询message表失败..", err.Error())
		return
	}
	fmt.Println(i)

	for i, v := range messages {
		fmt.Println(i, v.Id, v.Username, v.Message)
	}

}
