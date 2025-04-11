package main

import (
	"beego-app/audit"
	"beego-app/conf/injector"
	"beego-app/routers"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
)

func main() {
	//Userコントローラ構造体を初期化
	var handler = injector.InitializeUserHandler()
	audit.Log.Info("Service init:", handler.UserService)
	routers.InitUserRoutes(handler)
	//routers.InitBookRoutes()
	web.Run()
}
