package services

import "github.com/yhkl-dev/turtle-dove-beego/turtle-api/app/tables"

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
