package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

// 用户管理
type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	fmt.Println("Test")
	this.TplName = "index.tpl"
}
