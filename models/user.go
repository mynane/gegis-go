package models

import (
	"github.com/astaxie/beego/orm"
	)

type User struct {
	Id 		int		`json:"id" orm:"auto" orm:"column(id)"`
	Name 	string	`json:"name" orm:"column(name)"`
	Sex 	string	`json:"sex" orm:"column(sex)"`
	Age 	int		`json:"age" orm:"column(age)"`
	Email 	string	`json:"email" orm:"column(email)"`
	Phone 	string 	`json:"phone" orm:"column(phone)"`
	Pwd     string	`json:"pwd" orm:"column(pwd)"`
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
