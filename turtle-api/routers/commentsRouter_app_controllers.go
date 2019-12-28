package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/yhkl-dev/turtle-dove-beego/turtle-api/app/controllers:RoleController"] = append(beego.GlobalControllerRouter["github.com/yhkl-dev/turtle-dove-beego/turtle-api/app/controllers:RoleController"],
        beego.ControllerComments{
            Method: "GetAlliRoles",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yhkl-dev/turtle-dove-beego/turtle-api/app/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/yhkl-dev/turtle-dove-beego/turtle-api/app/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yhkl-dev/turtle-dove-beego/turtle-api/app/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/yhkl-dev/turtle-dove-beego/turtle-api/app/controllers:UserController"],
        beego.ControllerComments{
            Method: "AddUser",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yhkl-dev/turtle-dove-beego/turtle-api/app/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/yhkl-dev/turtle-dove-beego/turtle-api/app/controllers:UserController"],
        beego.ControllerComments{
            Method: "DeleteUser",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yhkl-dev/turtle-dove-beego/turtle-api/app/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/yhkl-dev/turtle-dove-beego/turtle-api/app/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdateUserProfile",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/yhkl-dev/turtle-dove-beego/turtle-api/app/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/yhkl-dev/turtle-dove-beego/turtle-api/app/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
