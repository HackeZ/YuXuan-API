// @APIVersion 1.0.0
// @Title YuXuan-Shop API
// @Description YuXuanAPI is a RESTful API for YuXuan-Shop
// @Contact a767110505@gmail.com
// @TermsOfServiceUrl http://hackez.github.io
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"YuXuanAPI/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),

		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),

		beego.NSNamespace("/yx_admin",
			beego.NSInclude(
				&controllers.YxAdminController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
