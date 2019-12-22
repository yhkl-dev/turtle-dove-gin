package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/yhkl-dev/turtle-dove-beego/turtle-api/app/tables"
	"github.com/yhkl-dev/turtle-dove-beego/turtle-api/utils"
)

type userService struct{}

func (us *userService) table() string {
	return tableName("user")
}

func (us *userService) valid(user *tables.User) error {
	valid := validation.Validation{}
	b, _ := valid.Valid(user)

	if !b {
		for _, err := range valid.Errors {
			return errors.New(err.Message)
		}
	}
	return nil
}

func (us *userService) GetUser(userID int) (*tables.User, error) {
	user := &tables.User{}
	user.Id = userID

	err := o.Read(user)
	return user, err
}

func (us *userService) GetTotal() (int64, error) {
	return o.QueryTable(us.table()).Count()
}

func (us *userService) GetUserList(page, pageSize int) ([]orm.Params, error) {

	offset := (page - 1) * pageSize
	if offset < 0 {
		offset = 0
	}

	var users []orm.Params
	queryset := o.QueryTable(us.table())
	_, err := queryset.OrderBy("Id").Limit(pageSize, offset).Values(&users, "Id", "UserName", "Email", "LastLogin", "RealName", "IsActive")
	return users, err
}

// 根据用户名称获取用户信息
func (us *userService) GetUserByName(userName string) (*tables.User, error) {
	user := &tables.User{}
	user.UserName = userName
	err := o.Read(user, "UserName")
	return user, err
}

// 创建用户
func (us *userService) AddUser(userName, userPassword, realName, email string) (*tables.User, error) {

	if isExist, _ := us.GetUserByName(userName); isExist.Id > 0 {
		return nil, errors.New("用户名已存在")
	}

	user := &tables.User{}
	user.UserName = userName
	user.UserPassword = utils.StringToMd5(userPassword)
	user.RealName = realName
	user.Email = email
	user.UpdateTime = time.Now()

	if err := us.valid(user); err != nil {
		return nil, err
	}
	_, err := o.Insert(user)
	return user, err
}

// 删除用户
func (us *userService) DeleteUser(userID int) error {
	tableUser := &tables.User{Id: userID}
	id, err := o.Delete(tableUser)
	if id == 0 {
		return errors.New(fmt.Sprintf("userID %d does not exist", tableUser.Id))
	}
	return err
}

// update user info
func (us *userService) UpdateUser(user *tables.User, fields ...string) error {

	if user.Id == 0 {
		return fmt.Errorf("userID %d does not exist", user.Id)
	}
	if err := us.valid(user); err != nil {
		return err
	}
	_, err := o.Update(user, fields...)
	return err
}

// update user password
func (us *userService) ChangePassword(user *tables.User, password string) error {

	user.UserPassword = utils.StringToMd5(password)
	_, err := o.Update(user, "UserPassword")

	return err
}

// Login
func (us *userService) Login(userName, userPassword string) (int, string, error) {
	user, err := us.GetUserByName(userName)
	if err != nil {
		return 0, "Wrong username", err
	}

	authResult := utils.CheckMd5Value(userPassword, user.UserPassword)
	if !authResult {

		return 0, "Authentication Failed, Wrong password", err
	}
	token, err := utils.GenerateToken(user.UserName, user.Email)
	if err != nil {
		return 0, "", err
	}
	return user.Id, token, nil
}
