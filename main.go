package main

import (
	_ "beego/BeeRobot/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {

	//init db
	initDB()

}

func main() {

	beego.Run()
}

func initDB() {

	// load config of db
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	db := beego.AppConfig.String("db")

	orm.Debug = true

	logs.Info("mysql connect start")

	// regist conn
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// connect
	conn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + db + "?charset=utf8"

	orm.RegisterDataBase("default", "mysql", conn)

	orm.RunSyncdb("default", false, true)

	logs.Info("mysql connect OK")
}
