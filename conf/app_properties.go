package conf

import (
	"beego-app/conf/orm"
)

type AppProperties struct {
	Db orm.DbProps
	/* todo Mail MailProps */
}

func NewAppProperties(props *orm.DbProps) *AppProperties {
	return &AppProperties{
		Db: *props,
	}
}
