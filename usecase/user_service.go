package usecase

import (
	"beego-app/audit"
	txt "beego-app/context/orm"
	"beego-app/model"

	"github.com/beego/beego/v2/client/orm"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) FindUserById(id int) (*model.User, error) {
	result, err := txt.NewTxTemplate().Tx(func(rep orm.Ormer) (any, error) {
		return model.FindUserById(id, rep)
	})
	if err != nil {
		return nil, err
	}
	audit.Log.Info(result.(*model.User).Name)
	return result.(*model.User), nil

}

/** Register User*/
func (u *UserService) RegisterUser() (*model.User, error) {
	result, err := txt.NewTxTemplate().Tx(func(rep orm.Ormer) (any, error) {
		return model.RegisterUser(rep)
	})
	if err != nil {
		return nil, err
	}
	return result.(*model.User), nil
}
