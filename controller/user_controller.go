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
}

type UserFindRequest struct {
	UserID int `json:"user_id" validate:"required"` // Required UserID
}

func (u *UserController) RegisterUser() {
	u.UserService.RegisterUser()
}
