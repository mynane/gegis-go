package models

import "github.com/astaxie/beego/orm"

type Email struct {
	Id 		int		`json:"id" orm:"auto" orm:"column(id)"`
	Email 	string  `json:"email" orm:"column(email)"`
	Time 	string  `json:"time" orm:"column(time)"`
}

func init()  {
	orm.RegisterModel(new(Email))
}

func EmailAdd(o orm.Ormer, email Email) (int64, error) {
	res, err := o.Insert(&email)
	return res, err
}

func EmailAll(o orm.Ormer) ([]Email, error) {
	var emails []Email
	_, err := o.QueryTable("email").All(&emails, "id", "name")
	return emails, err
}

func EmailFindByEmail(o orm.Ormer, id string) (Email, error) {
	var email Email
	err := o.QueryTable("email").Filter("email", email).One(&email, "id")
	return email, err
}