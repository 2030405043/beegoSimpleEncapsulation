package DBconf

import (
	"1225/entities"
	"github.com/astaxie/beego"
)

var dbconf *entities.DBconf

func init() {
	driverName := beego.AppConfig.String("driverName")
	account := beego.AppConfig.String("account")
	password := beego.AppConfig.String("password")
	DB := beego.AppConfig.String("DB")
	enCoding := beego.AppConfig.String("enCoding")
	ip := beego.AppConfig.String("ip")
	port := beego.AppConfig.String("port")

	maxIdle, _ := beego.AppConfig.Int("maxIdle")
	maxConn, _ := beego.AppConfig.Int("maxConn")

	dbconf = &entities.DBconf{DriverName: driverName, Account: account, Password: password,
		DB: DB, EnCoding: enCoding, Ip: ip, Port: port, MaxIdle: maxIdle, MaxConn: maxConn}

}

func GetDBconf() entities.DBconf {
	return *dbconf
}
