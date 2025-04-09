package main

import (
	"beego-app/audit"
	"beego-app/conf/injector"
	"beego-app/routers"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
)

func main() {

	// ORM を初期化して接続確認
	orm.Debug = true
	o := orm.NewOrm()
	if o == nil {
		audit.Log.Error("ORM init")
		return
	}
	audit.Log.Info("ORM init")

	//Userコントローラ構造体を初期化
	var handler = injector.InitializeUserHandler(o)
	audit.Log.Info("Service init:", handler.UserService)

	routers.InitUserRoutes(handler)
	//routers.InitBookRoutes()
	web.Run()
}
