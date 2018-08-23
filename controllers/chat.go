package controllers

import (
	"github.com/astaxie/beego"
	"gegis/utils"
	"github.com/astaxie/beego/orm"
	"gegis/models"
	"strings"
	)

type ChatController struct {
	beego.Controller
}

var hub *utils.Hub

func init()  {
	//initWS()

}

func initWS()  {
	o := orm.NewOrm()
	user, _ := models.FindById(o, 2)
	users := strings.Split(user.Rooms, ",")
	hub = utils.NewHub(users)
	go hub.Run()
}

func (this *ChatController) Chat() {
	initWS()
	this.TplName = "ws.tpl"
}

func (this *ChatController) WS() {
	utils.ServeWs(hub, this.Ctx.ResponseWriter, this.Ctx.Request)
}
