package main

import (
	"beego-app/audit"
	"beego-app/conf"
	"beego-app/conf/injector"
)

var AppProperties *conf.AppProperties

func init() {
	audit.Init()
	AppProperties = injector.InitializeAppProperties()
}
