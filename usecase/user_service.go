package usecase

import (
	"beego-app/model"

	"github.com/beego/beego/v2/client/orm"
)

type UserService struct {
	Orm orm.Ormer
}

func NewUserService(db orm.Ormer) *UserService {
	return &UserService{
		Orm: db,
	}
}

func (u *UserService) FindUserByID(id int) {

}
func (u *UserService) RegisterUser() (*model.User, error) {
	// todo authorization
	// todo transaction
	user, err := model.RegisterUser(u.Orm)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserService) ChangeUser(id int, newName string) (*model.User, error) {
	user, err := model.FindUser(id, u.Orm)
	if err != nil {
		return nil, err
	}

	err = user.ChangeUser(newName, u.Orm)
	if err != nil {
		return nil, err
	}

	return user, nil
}
