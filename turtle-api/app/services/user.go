package services

import (
	"errors"

	"github.com/yhkl-dev/turtle-dove-beego/turtle-api/app/tables"
)

type userService struct{}

func (us *userService) table() string {
	return tableName("user")
}

func (us *userService) GetUser(userId int) (*tables.User, error) {
	user := &tables.User{}
	user.Id = userId

	err := o.Read(user)
	return user, err
}

func (us *userService) GetTotal() (int64, error) {
	return o.QueryTable(us.table()).Count()
}

func (us *userService) GetUserList(page, pageSize int) ([]tables.User, error) {

	offset := (page - 1) * pageSize
	if offset < 0 {
		offset = 0
	}
	var users []tables.User
	queryset := o.QueryTable(us.table())
	_, err := queryset.OrderBy("id").Limit(pageSize, offset).All(&users)
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
	user.UserPassword = userPassword
	user.RealName = realName
	user.Email = email
	_, err := o.Insert(user)
	return user, err
}
