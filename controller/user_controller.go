package controller

import (
	"beego-app/usecase"

	beego "github.com/beego/beego/v2/server/web"
)

type UserController struct {
	beego.Controller
	UserService *usecase.UserService
}

func NewUserController(userService *usecase.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (u *UserController) FindUser() {
}
func (u *UserController) RegisterUser() {
	u.UserService.RegisterUser()
}
