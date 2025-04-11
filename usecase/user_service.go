package usecase

import (
	"beego-app/conf/orm/txtemplate"
	"beego-app/model"

	"github.com/beego/beego/v2/client/orm"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) FindUserByID(id int) {

}

/** Register User*/
func (u *UserService) RegisterUser() (*model.User, error) {
	result, err := txtemplate.NewTxTemplate().Tx(func(rep orm.Ormer) (any, error) {
		return model.RegisterUser(rep)
	})
	if err != nil {
		return nil, err
	}
	return result.(*model.User), nil
}
