package model

import (
	"beego-app/model/types"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type User struct {
	ID            int                 `orm:"auto"`
	Name          string              `orm:"size(30)"`
	NameKana      string              `orm:"size(30)"`
	Profile       string              `orm:"size(250)"`
	Email         string              `orm:"size(100)"`
	AuthorityType types.AuthorityType `orm:"default(0)"`
	StatusType    types.StatusType    `orm:"default(0)"`
	RegisterId    string              `orm:"auto_now_add"`
	RegisterDate  time.Time           `orm:"auto_now_add"`
	UpdateId      string              `orm:"auto_now_add"`
	UpdateDate    time.Time           `orm:"auto_now_add"`
}

func FindUser(id int, orm orm.Ormer) (*User, error) {
	user := &User{}
	// err := orm.QueryTable("user").Filter("id", id).One(user)
	// if err != nil {
	// 	return nil, err
	// }
	return user, nil
}

/** Register User*/
func RegisterUser(rep orm.Ormer) (*User, error) {
	user := &User{
		Name:          "testuser",
		Email:         "txuser@gmail.com",
		AuthorityType: types.User,
	}
	rep.Insert(user)

	return user, nil
}

// ポインタレシーバはfindUserで取得したUserになる
func (u *User) ChangeUser(newName string, orm orm.Ormer) error {
	u.Name = newName
	_, err := orm.Update(u)
	return err
}
