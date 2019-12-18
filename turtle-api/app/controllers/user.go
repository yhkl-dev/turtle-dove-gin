package controllers

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/yhkl-dev/turtle-dove-beego/turtle-api/app/services"
)

// Operations for User
type UserController struct {
	beego.Controller
}

// @Title GetAllUser
// @Description get all users
// @success 200
// @router / [get]
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

// @Title Add User
// @Description add a new user
// @Success 201 {object} services.UserService.User
// @Failure 400
// @router / [post]
func (us *UserController) AddUser() {

	valid := validation.Validation{}

	userName := us.GetString("userName")
	email := us.GetString("email")
	realName := us.GetString("realName")
	userPassword := us.GetString("userPassword")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			fmt.Printf("%s\n", err.Message)
		}
	}
	user, err := services.UserService.AddUser(userName, userPassword, realName, email)
	if err != nil {
		fmt.Println(err)
		us.Data["code"] = 404
		us.Data["error"] = err
		fmt.Println(us.Data)
	}
	fmt.Println(user)
	us.ServeJSON()
}
