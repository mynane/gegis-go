package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"gegis/models"
	"encoding/json"
)

// 标签管理
type TagController struct {
	beego.Controller
}

func (c *TagController) URLMapping() {
	c.Mapping("All", c.All)
}

// @Title 新增标签
// @Description 新增标签
// @Param body body models.tag.Tag true "用户信息"
// @Success 200 {int} models.user.User.Id
// @Failure 400 Invalid email supplied
// @Failure 404 User not found
// @router / [post]
func (this *TagController) Add() {
	o := orm.NewOrm()
	var tag models.Tag
	json.Unmarshal(this.Ctx.Input.RequestBody, &tag)
	res, err := models.TagAdd(o, tag)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = map[string]int64{"success": res}
	}
	this.ServeJSON()
}

// @Title 获取所有标签
// @Description 获取所有标签
// @Success 200 {object} models.tag.Tag
// @Failure 400 Invalid email supplied
// @Failure 404 User not found
// @router / [get]
func (this *TagController) All() {
	o := orm.NewOrm()
	tags, err := models.TagAll(o)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = tags
	}
	this.ServeJSON()
}


