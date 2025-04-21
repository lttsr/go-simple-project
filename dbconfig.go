package main

import (
	"beego-app/audit"
	"beego-app/model"
	"fmt"
	"log"

	"github.com/beego/beego/v2/client/orm"
)

func init() {
	if AppProperties == nil {
		audit.Log.Error("AppProperties not init")
	}
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		AppProperties.Db.Datasource.UserName,
		AppProperties.Db.Datasource.UserPassword,
		AppProperties.Db.Datasource.DBName,
		AppProperties.Db.Datasource.Host,
		AppProperties.Db.Datasource.Port,
		AppProperties.Db.Datasource.SSLMode)
	// Register the PostgreSQL driver
	orm.RegisterDriver("postgres", orm.DRPostgres)
	// Register the database with the ORM
	err := orm.RegisterDataBase("default", AppProperties.Db.Datasource.DriverName, connStr)
	if err != nil {
		log.Fatalf("Failed to register database: %v", err)
		audit.Log.Error("Register DB", err.Error())
	}
	audit.Log.Info("DB init")
	orm.Debug = true
	orm.ResetModelCache()
	orm.RegisterModel(new(model.User))
	orm.RegisterModel(new(audit.AuditActor))
	if err := orm.RunSyncdb("default", false, true); err != nil {
		audit.Log.Error("Sync DB Resource", err.Error())
	}
}
