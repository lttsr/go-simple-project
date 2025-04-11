package injector

import (
	"beego-app/conf"
	"beego-app/conf/datasource"
	"beego-app/controller"
	"beego-app/usecase"

	"github.com/beego/beego/v2/client/orm"
	"github.com/google/wire"
)

func InitializedAppProperties() *conf.AppProperties {
	wire.Build(
		datasource.NewDataSourceProperties,
		datasource.NewDbProps,
		conf.NewAppProperties,
	)
	return nil

}

func InitializedUserHandler(db orm.Ormer) *controller.UserController {
	wire.Build(
		usecase.NewUserService,
		controller.NewUserController,
	)
	return nil
}
