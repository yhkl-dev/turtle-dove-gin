package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/yhkl-dev/turtle-dove-beego/turtle-api/app/services"
	"github.com/yhkl-dev/turtle-dove-beego/turtle-api/app/tables"
	"github.com/yhkl-dev/turtle-dove-beego/turtle-api/utils"
)

// UserController struct
// Operations for User
type UserController struct {
	beego.Controller
}

// GetAll func
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

// AddUser function
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

	infos := us.Ctx.Input.RequestBody

	if err := json.Unmarshal(infos, &userParse); err != nil {
		us.CustomAbort(500, err.Error())
	}

	user, err := services.UserService.AddUser(
		userParse.UserName,
		userParse.UserPassword,
		userParse.RealName,
		userParse.Email)

	if err != nil {
		us.CustomAbort(400, err.Error())
	}
	us.Data["json"] = user
	us.ServeJSON()
}

// @Title Delete User By userID
// @Param userID query int required
// @success 200
// @router / [delete]
func (us *UserController) DeleteUser() {

	id, _ := us.GetInt("userID")
	err := services.UserService.DeleteUser(id)

	if err != nil {
		us.Ctx.ResponseWriter.WriteHeader(400)
		us.Ctx.ResponseWriter.Write([]byte(err.Error()))
		us.StopRun()
	}
	us.Data["json"] = "ok"
	us.ServeJSON()
}

// @Title update user profile
// @router / [put]
func (us *UserController) UpdateUserProfile() {
	id, _ := us.GetInt("userID")
	user, err := services.UserService.GetUser(id)
	if err != nil {
		us.CustomAbort(404, err.Error())
	}

	var userParse tables.User

	infos := us.Ctx.Input.RequestBody
	err = json.Unmarshal(infos, &userParse)
	if err != nil {
		us.CustomAbort(500, err.Error())
	}

	if len(userParse.RealName) != 0 {
		user.RealName = userParse.RealName
	}
	if len(userParse.Email) != 0 {
		user.Email = userParse.Email
	}

	err = services.UserService.UpdateUser(user, "Email", "RealName")

	if err != nil {
		us.CustomAbort(500, err.Error())
	}
	// update password
	if len(userParse.UserPassword) > 0 {
		err = services.UserService.ChangePassword(user, userParse.UserPassword)
		if err != nil {
			us.CustomAbort(500, err.Error())
		}
	}
	us.Data["json"] = "ok"
	us.ServeJSON()
}

// @router /login [post]
func (us *UserController) Login() {
	var userParse tables.User
	infos := us.Ctx.Input.RequestBody
	if err := json.Unmarshal(infos, &userParse); err != nil {
		us.CustomAbort(500, err.Error())
	}

	id, token, err := services.UserService.Login(
		userParse.UserName,
		userParse.UserPassword,
	)

	if err != nil {
		us.CustomAbort(401, err.Error())
	}
	_, err = services.AuthService.AddToken(id, token)
	if err != nil {
		us.CustomAbort(500, err.Error())
	}

	us.Ctx.ResponseWriter.Header().Set("Authorization", token)
	us.Data["json"] = token
	us.ServeJSON()
}
