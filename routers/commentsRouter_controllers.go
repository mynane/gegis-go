package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["gegis/controllers:TagController"] = append(beego.GlobalControllerRouter["gegis/controllers:TagController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gegis/controllers:TagController"] = append(beego.GlobalControllerRouter["gegis/controllers:TagController"],
		beego.ControllerComments{
			Method: "All",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gegis/controllers:UserController"] = append(beego.GlobalControllerRouter["gegis/controllers:UserController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gegis/controllers:UserController"] = append(beego.GlobalControllerRouter["gegis/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gegis/controllers:UserController"] = append(beego.GlobalControllerRouter["gegis/controllers:UserController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gegis/controllers:UserController"] = append(beego.GlobalControllerRouter["gegis/controllers:UserController"],
		beego.ControllerComments{
			Method: "FindById",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gegis/controllers:UserController"] = append(beego.GlobalControllerRouter["gegis/controllers:UserController"],
		beego.ControllerComments{
			Method: "Remove",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
