package model

import (
	"beego-app/model/types"
	"fmt"

	"github.com/beego/beego/v2/client/orm"
)

type User struct {
	ID            int                 `orm:"auto"`
	Name          string              `orm:"size(100)"`
	Email         string              `orm:"size(100)"`
	AuthorityType types.AuthorityType `orm:"default(0)"`
}

func FindUser(id int, orm orm.Ormer) (*User, error) {
	user := &User{}
	// err := orm.QueryTable("user").Filter("id", id).One(user)
	// if err != nil {
	// 	return nil, err
	// }
	return user, nil
}
func RegisterUser(orm orm.Ormer) (*User, error) {
	fmt.Println("accessedModel:register")
	user := &User{
		Name:          "test",
		Email:         "example.com",
		AuthorityType: types.User,
	}
	orm.Insert(user)

	return user, nil
}

// ポインタレシーバはfindUserで取得したUserになる
func (u *User) ChangeUser(newName string, orm orm.Ormer) error {
	u.Name = newName
	_, err := orm.Update(u)
	return err
}
