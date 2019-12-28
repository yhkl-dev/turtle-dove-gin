package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/yhkl-dev/turtle-dove-beego/turtle-api/app/services"
	"github.com/yhkl-dev/turtle-dove-beego/turtle-api/app/tables"
	"github.com/yhkl-dev/turtle-dove-beego/turtle-api/utils"
)

// UserController struct, Operation for user api
type UserController struct {
	beego.Controller
}

// GetAll func
// @Title GetAllUser
// @Description Return all users
// @Failure 500 internal error
// @Param page query int false page
// @Param page_size query int false page_size
// @Param UserName query string false user_name
// @Param Email query string false user_email
// @Param RealName query string false user_real_name
// @Param IsActive query int false 0: no or 1:yes
// @Success 200 success
// @Failure 404 data not found
// @router / [get]
func (us *UserController) GetAll() {

	page, _ := strconv.Atoi(us.GetString("page"))
	if page < 1 {
		page = 1
	}
	pageSize, err := strconv.Atoi(us.GetString("page_size"))
	if err != nil {
		pageSize = 10
	}

	realName := us.GetString("realName")
	email := us.GetString("email")
	userName := us.GetString("userName")
	isActive, _ := us.GetInt("isActive")

	remoteAddr := us.Ctx.Request.Host

	users, count, _ := services.UserService.GetUserList(page, pageSize, userName, realName, email, isActive)

	data := new(utils.JSONData)
	data.Data = users
	data.Count = int(count)
	data.Pager = utils.NewPager(page,
		pageSize,
		data.Count,
		remoteAddr+beego.URLFor("UserController.GetAll"))

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
// @Description Delete user by userid
// @Param userID query int true userid
// @Failure 400 bad request
// @Success 204 success
// @router / [delete]
func (us *UserController) DeleteUser() {

	id, _ := us.GetInt("userID")
	err := services.UserService.DeleteUser(id)

	if err != nil {
		us.CustomAbort(400, err.Error())
	}
	us.Data["json"] = "ok"
	us.ServeJSON()
}

// @Title update user profile
// @Description update user profile and password
// @Param userID query int true userid
// @Param Email body string false
// @Param UserPassword body string false
// @Param RealName body string false
// @Failure 500 parse json error
// @Failure 404 uer not found
// @Success 204 ok
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

// @Title Login interface
// @Description return a jwt token
// @Param UserName body string true  user_name
// @Param UserPassword body string true user_password
// @Failure 500 internal error
// @Failure 401 authentication error
// @Success 201 login success
// @router /login [post]
func (us *UserController) Login() {
	var userParse tables.User
	infos := us.Ctx.Input.RequestBody
	if err := json.Unmarshal(infos, &userParse); err != nil {
		us.CustomAbort(500, err.Error())
	}

	token, err := services.UserService.Login(
		userParse.UserName,
		userParse.UserPassword,
	)

	if err != nil {
		us.CustomAbort(401, err.Error())
	}

	err = services.RedisService.Set(userParse.UserName, token)
	if err != nil {
		us.CustomAbort(500, err.Error())
	}
	us.Ctx.ResponseWriter.Header().Set("Authorization", token)
	us.Data["json"] = token
	us.ServeJSON()
}
