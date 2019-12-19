package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/yhkl-dev/turtle-dove-beego/turtle-api/app/services"
	"github.com/yhkl-dev/turtle-dove-beego/turtle-api/app/tables"
	"github.com/yhkl-dev/turtle-dove-beego/turtle-api/utils"
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

	pageSize, err := strconv.Atoi(us.GetString("page_size"))
	if err != nil {
		pageSize = 10
	}

	remoteAddr := us.Ctx.Request.Host

	if page < 1 {
		page = 1
	}

	count, _ := services.UserService.GetTotal()
	users, _ := services.UserService.GetUserList(page, pageSize)

	data := new(utils.JSONData)

	data.Data = users
	data.Count = int(count)
	data.Pager = utils.NewPager(page, pageSize, data.Count, remoteAddr+beego.URLFor("UserController.GetAll"))

	us.Data["json"] = data

	us.ServeJSON()
}

// @Title Add User
// @Description add a new user
// @Param RealName     body string true
// @Param Email        body string true
// @Param RealName     body string true
// @Param UserPassword body string true
// @Success 201 {object} services.UserService.User
// @Failure 400
// @router / [post]
func (us *UserController) AddUser() {

	var userParse tables.User
	valid := validation.Validation{}

	infos := us.Ctx.Input.RequestBody

	fmt.Println(string(infos))

	err := json.Unmarshal(infos, &userParse)
	if err != nil {
		fmt.Println("json parse error", err.Error())
	}

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			fmt.Printf("参数错误 %s\n", err.Message)
		}
	}

	user, err := services.UserService.AddUser(
		userParse.UserName,
		userParse.UserPassword,
		userParse.RealName,
		userParse.Email)
	if err != nil {
		us.Ctx.ResponseWriter.WriteHeader(400)
		us.Ctx.ResponseWriter.Write([]byte(err.Error()))
		us.StopRun()
	}
	us.Data["json"] = user
	us.ServeJSON()
}
