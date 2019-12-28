// @APIVersion 1.0.0
// @Title TurtleDove API
// @Description apis for turtle-dove
// @Contact kaiyang939325@gmail.com
package routers

import (
	"github.com/yhkl-dev/turtle-dove-beego/turtle-api/app/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/role",
			beego.NSInclude(
				&controllers.RoleController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
