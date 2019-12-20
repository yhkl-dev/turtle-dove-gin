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
	_, err := queryset.OrderBy("Id").Limit(pageSize, offset).Values(&users, "Id", "UserName", "Email", "LastLogin", "RealName")
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
	valid := validation.Validation{}

	if isExist, _ := us.GetUserByName(userName); isExist.Id > 0 {
		return nil, errors.New("用户名已存在")
	}

	user := &tables.User{}
	user.UserName = userName
	user.UserPassword = utils.StringToMd5(userPassword)
	user.RealName = realName
	user.Email = email
	user.UpdateTime = time.Now()

	if _, err := valid.Valid(user); err != nil {
		return nil, err
	}
	_, err := o.Insert(user)
	return user, err
}

// 删除用户
func (this *userService) DeleteUser(userID int) error {
	tableUser := &tables.User{Id: userID}
	id, err := o.Delete(tableUser)
	if id == 0 {
		return errors.New(fmt.Sprintf("userID %d does not exist", userID))
	}
	return err
}
