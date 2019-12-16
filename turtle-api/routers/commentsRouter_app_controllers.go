package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/yhkl-dev/turtle-dove-beego/turtle-api/app/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/yhkl-dev/turtle-dove-beego/turtle-api/app/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
