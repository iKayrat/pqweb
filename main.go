package main

import (
	_ "pqweb/routers"

	"github.com/astaxie/beego/orm"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	orm.Debug = true
	beego.BConfig.WebConfig.Session.SessionOn = true

	beego.Run()
}
