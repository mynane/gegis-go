package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"gegis/models"
	"fmt"
	"strconv"
	"encoding/json"
	"crypto/md5"
	"io"
)

// 用户管理
type UserController struct {
	beego.Controller
}

func (c *UserController) URLMapping() {
	c.Mapping("Add", c.Add)
	c.Mapping("Get", c.Get)
	c.Mapping("FindById", c.FindById)
	c.Mapping("Remove", c.Remove)
	c.Mapping("Update", c.Update)
}

func encryption(password string) string {
	//假设用户名abc，密码123456
	h := md5.New()
	io.WriteString(h, password)

	//pwmd5等于e10adc3949ba59abbe56e057f20f883e
	pwmd5 :=fmt.Sprintf("%x", h.Sum(nil))

	//指定两个 salt： salt1 = @#$%   salt2 = ^&*()
	salt1 := "@#$%"
	salt2 := "^&*()"

	//salt1+用户名+salt2+MD5拼接
	io.WriteString(h, salt1)
	io.WriteString(h, "abc")
	io.WriteString(h, salt2)
	io.WriteString(h, pwmd5)

	return fmt.Sprintf("%x", h.Sum(nil))
}

// @Title 新增用户
// @Description 新增用户
// @Param body body models.user.User true "用户信息"
// @Success 200 {int} models.user.User.Id
// @Failure 400 Invalid email supplied
// @Failure 404 User not found
// @router / [post]
func (this *UserController) Add() {
	o := orm.NewOrm()
	var user models.User
	json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	user.Pwd = encryption(user.Pwd)
	res, err := models.Add(o, user)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = map[string]int64{"success": res}
	}
	this.ServeJSON()
}

// @Title 获取所有用户
// @Description 获取所有用户
// @Success 200 {object} models.user.User
// @Failure 400 Invalid email supplied
// @Failure 404 User not found
// @router / [get]
func (this *UserController) Get() {
	o := orm.NewOrm()
	users, err := models.FindAll(o)
	fmt.Println(users)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = users
	}
	this.ServeJSON()
}

// @Title 通过id获取用户
// @Description 通过id获取用户
// @Param   id     path    string  true        "用户id"
// @Success 200 {object} models.user.User
// @Failure 400 Invalid email supplied
// @Failure 404 User not found
// @router /:id [get]
func (this *UserController) FindById() {
	o := orm.NewOrm()
	sid := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(sid)
	user, err := models.FindById(o, id)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = user
	}
	this.ServeJSON()
}

// @Title 删除用户
// @Description 删除用户
// @Param   id     path    string  true        "用户id"
// @Success 200 {object} models.user.User
// @Failure 400 Invalid email supplied
// @Failure 404 User not found
// @router /:id [delete]
func (this *UserController) Remove() {
	o := orm.NewOrm()
	sid := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(sid)
	user, err := models.FindById(o, id)
	res, err := models.Remove(o, user)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = map[string]int64{"succuss": res}
	}
	this.ServeJSON()
}

// @Title 更新用户
// @Description 更新用户
// @Param body body models.user.User true "用户信息"
// @Success 200 {int} models.user.User.Id
// @Failure 400 Invalid email supplied
// @Failure 404 User not found
// @router / [put]
func (this *UserController) Update() {
	o := orm.NewOrm()
	var user models.User
	json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	res, err := models.Update(o, user)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = map[string]int64{"success": res}
	}
	this.ServeJSON()
}

// @Title 用户登录
// @Description 用户登录
// @Param body body models.user.Login true "用户信息"
// @Success 200 {int} models.user.User.Id
// @Failure 400 Invalid email supplied
// @Failure 404 User not found
// @router / [put]

