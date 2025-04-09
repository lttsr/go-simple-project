package orm

import (
	"beego-app/audit"

	"github.com/beego/beego/v2/server/web"
)

type DataSourceProperties struct {
	DriverName   string
	DBName       string
	UserName     string
	UserPassword string
	Host         string
	Port         string
	SSLMode      string
}

func NewDataSourceProperties() *DataSourceProperties {
	props := &DataSourceProperties{}
	props.ReadAppConf()
	return props
}

func (d *DataSourceProperties) ReadAppConf() {
	web.LoadAppConfig("ini", "app.conf")

	driverName, err := web.AppConfig.String("db_driver_name")
	if err != nil {
		audit.Log.Error("drivername", err.Error())
	}
	d.DriverName = driverName

	dbName, err := web.AppConfig.String("db_name")
	if err != nil {
		audit.Log.Error("db_name", err.Error())
	}
	d.DBName = dbName

	userName, err := web.AppConfig.String("db_user_name")
	if err != nil {
		audit.Log.Error("db_user_name", err.Error())
	}
	d.UserName = userName

	password, err := web.AppConfig.String("db_user_password")
	if err != nil {
		audit.Log.Error("db_user_password", err.Error())
	}
	d.UserPassword = password

	dbHost, err := web.AppConfig.String("db_host")
	if err != nil {
		audit.Log.Error("db_host", err.Error())
	}
	d.Host = dbHost

	dbPort, err := web.AppConfig.String("db_port")
	if err != nil {
		audit.Log.Error("db_port", err.Error())
	}
	d.Port = dbPort

	dbSSLMode, err := web.AppConfig.String("db_sslmode")
	if err != nil {
		audit.Log.Error("db_sslmode", err.Error())
	}
	d.SSLMode = dbSSLMode
}
