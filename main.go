package main

import (
	"mall/models"
	_ "mall/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	models.RegisterDB()
}
func main() {
	orm.Debug = true
	/*orm.RunSyncdb("default", true, true)
	o := orm.NewOrm()
	configInfo := beego.AppConfig.String("mysql:dbname")
	o.Using(configInfo)
	*/

	beego.Run()
}
