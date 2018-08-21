// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"gegis/controllers"
	"gegis/utils"
	"fmt"
	"net/http"
	)

func auth(ctx *context.Context)  {
	var token string
	cookie, err := ctx.Request.Cookie("Authorization")
	header := ctx.Request.Header.Get("Authorization")
	if header != "" {
		token = string(header)
	} else if cookie.Value != "" {
		token = string(cookie.Value)
	}
	fmt.Println(utils.GenToken())
	result, error := utils.CheckToken(token)
	if err != nil || !result {
		ctx.Output.Status = http.StatusUnauthorized
		ctx.Output.Body(utils.NotAuth(error))
	}
}

func init() {
	//beego.InsertFilter("/*", beego.BeforeRouter, func(ctx *context.Context) {
	//	cookie, err := ctx.Request.Cookie("Authorization")
	//	if err != nil || !utils.CheckToken(cookie.Value) {
	//		http.Redirect(ctx.ResponseWriter, ctx.Request, "/", http.StatusMovedPermanently)
	//	}
	//})
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user", beego.NSInclude(&controllers.UserController{})),
		beego.NSNamespace("/tag", beego.NSInclude(&controllers.TagController{})),
	)
	beego.AddNamespace(ns)
}