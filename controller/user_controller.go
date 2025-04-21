package controller

import (
	"beego-app/usecase"
)

type UserController struct {
	DefaultController
	UserService *usecase.UserService
}

func NewUserController(userService *usecase.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (u *UserController) FindUser() {
	var req UserFind
	if err := u.ParseForm(&req); err != nil {
		u.CustomAbort(400, "Invalid query parameters")
		return
	}

	user, err := u.UserService.FindUserById(req.UserId)
	if err != nil {
		u.CustomAbort(500, "User not found")
		return
	}

	u.Data["json"] = user
	u.ServeJSON()

}

type UserFind struct {
	UserId int `form:"user_id" validate:"required"`
}

func (u *UserController) RegisterUser() {
	// DTO u.BindRequest(&RegisterUser{})
	// Params
	u.UserService.RegisterUser()

}

type RegisterUser struct {
	UserID   int    `json:"user_id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	NameKana string `json:"name_kana" validate:"required"`
	Profile  string `json:"profile"`
	Email    string `json:"email"`
}
