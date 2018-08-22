package models

import "github.com/astaxie/beego/orm"

type Tag struct {
	Id 		int		`json:"id" orm:"auto" orm:"column(id)"`
	Name 	string	`json:"name" orm:"column(name)"`
}

func init()  {
	orm.RegisterModel(new(Tag))
}

func TagAdd(o orm.Ormer, tag Tag) (int64, error) {
	res, err := o.Insert(&tag)
	return res, err
}

func TagAll(o orm.Ormer) ([]Tag, error) {
	var tags []Tag
	_, err := o.QueryTable("tag").All(&tags, "Id", "Name")
	return tags, err
}
