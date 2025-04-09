package routers

import (
	"beego-app/controller"

	"github.com/beego/beego/v2/server/web"
)

func InitUserRoutes(userController *controller.UserController) {
	web.Router("/user/find", userController, "get:FindUser")
	web.Router("/user/register", userController, "get:RegisterUser")
}

func InitBookRoutes() {
	web.Router("/book/find", &controller.BookController{}, "get:Find")
	web.Router("/book/register", &controller.BookController{}, "post:Register")
}
