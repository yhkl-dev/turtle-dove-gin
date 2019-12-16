package controllers

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/yhkl-dev/turtle-dove-beego/turtle-api/app/services"
)

// UserController: base controller
type UserController struct {
	beego.Controller
}

//@router / [get]
func (us *UserController) GetAll() {
	page, _ := strconv.Atoi(us.GetString("page"))

	if page < 1 {
		page = 1
	}

	count, _ := services.UserService.GetTotal()
	users, _ := services.UserService.GetUserList(page, 10)

	us.Data["json"] = users
	fmt.Println(count, users)
	us.ServeJSON()

}
