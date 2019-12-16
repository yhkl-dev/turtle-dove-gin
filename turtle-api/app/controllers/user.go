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
// @Param
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

// @router / [post]
func (us *UserController) AddUser() {
	valid := validation.Validation{}

	userName := us.GetString("userName")
	email := us.GetString("email")
	realName := us.GetString("realName")
	userPassword := us.GetString("userPassword")

	valid.Required(userName, "userName").Message("请输入用户名")
	valid.Required(email, "email").Message("请输入邮箱")
	valid.Email(email, "email").Message("Email无效")
	valid.MinSize(userPassword, 8, "userPassword").Message("密码长度不能小于6位")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			fmt.Printf("%s\n", err.Message)
		}
	}
	user, err := services.UserService.AddUser(userName, userPassword, realName, email)
	if err != nil {
		fmt.Println(err)

	}
	us.display()
}
