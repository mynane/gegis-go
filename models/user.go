package models

import (
	"github.com/astaxie/beego/orm"
	)

type User struct {
	Id 		int		`orm:"auto"`
	Name 	string
	Sex 	string
	Age 	int
	Email 	string
	Phone 	string
	Pwd     string
}

type Login struct {
	name 		string
	password 	string
}

func init()  {
	orm.RegisterModel(new(User))
}

func Add(o orm.Ormer, user User) (int64, error) {
	res, err := o.Insert(&user)
	return res, err
}

/**
 * 查询所有用户
 */
func FindAll(o orm.Ormer) ([]User, error) {
	var users []User
	_, err := o.QueryTable("user").All(&users, "Id", "Name")
	return users, err
}

func FindById(o orm.Ormer, id int) (User, error) {
	var user User
	err := o.QueryTable("user").Filter("Id", id).One(&user, "Id", "name")
	return user, err
}

func Remove(o orm.Ormer, user User) (int64, error) {
	res, err := o.Delete(&user, "id")
	return res, err
}

func Update(o orm.Ormer, user User) (int64, error) {
	res, err := o.Update(&user, "name")
	return res, err
}
