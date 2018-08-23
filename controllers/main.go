package controllers

import (
	"github.com/astaxie/beego"
	"gegis/utils"
)

// 用户管理
type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.TplName = "index.tpl"
}

func (this *MainController) Login() {
	token := utils.GenToken()
	this.Ctx.SetCookie("Authorization", token)
	//this.Data["xsrfdata"] = this.XSRFToken()
	this.TplName = "login.tpl"
}
