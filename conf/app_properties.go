package conf

import (
	"beego-app/conf/datasource"
)

type AppProperties struct {
	Db datasource.DbProps
	/* todo Mail MailProps */
}

func NewAppProperties(props *datasource.DbProps) *AppProperties {
	return &AppProperties{
		Db: *props,
	}
}
