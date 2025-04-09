package controller

import (
	beego "github.com/beego/beego/v2/server/web"
)

type BookController struct {
	beego.Controller
}

func (u *BookController) FindUser() {
}
func (u *BookController) RegisterUser() {}
