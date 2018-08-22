package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"gegis/models"
	"encoding/json"
	"time"
)

type EmailController struct {
	beego.Controller
}

func (this *EmailController) URLMapping() {
	this.Mapping("Add", this.Add)
	this.Mapping("FindByEmail", this.FindByEmail)
}

// @Title 新增邮箱
// @Description 新增邮箱
// @Param body body models.email.Email true "邮箱"
// @Success 200 {int} models.user.User.Id
// @Failure 400 Invalid email supplied
// @Failure 404 User not found
// @router / [post]
func (this *EmailController) Add() {
	o := orm.NewOrm()
	var email models.Email
	json.Unmarshal(this.Ctx.Input.RequestBody, &email)
	email.Time = time.Now().UTC().Format(time.UnixDate)
	res, err := models.EmailAdd(o, email)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = map[string]int64{"success": res}
	}
	this.ServeJSON()
}

// @Title 查询邮箱
// @Description 查询邮箱
// @Param email path models.email.Email.email true "邮箱"
// @Success 200 {int} models.email.Email
// @Failure 400 Invalid email supplied
// @Failure 404 User not found
// @router /:email [get]
func (this *EmailController) FindByEmail()  {
	o := orm.NewOrm()
	e := this.Ctx.Input.Param(":email")
	email, err := models.EmailFindByEmail(o, e)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = email
	}
	this.ServeJSON()
}
