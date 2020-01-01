package user

import (
	"time"

	"github.com/yhkl-dev/turtle-dove-beego/gin-turtle-api/database/models"
	"github.com/yhkl-dev/turtle-dove-beego/gin-turtle-api/utils/serializers"
)

// UserSerializer  struct for serialize data
type UserSerializer struct {
	ID         int        `json:"id"`
	UserName   string     `json:"user_name"`
	RealName   string     `json:"real_name"`
	IsActive   int        `json:"is_active"`
	Email      string     `json:"email"`
	LastLogin  *time.Time `json:"last_login"`
	CreateTime *time.Time `json:"create_time"`
	UpdateTime *time.Time `json:"update_time"`
}

// BuildUser single user serialzier
func BuildUser(user models.User) UserSerializer {
	return UserSerializer{
		ID:         user.ID,
		UserName:   user.UserName,
		RealName:   user.RealName,
		IsActive:   user.IsActive,
		Email:      user.Email,
		LastLogin:  user.LastLogin,
		CreateTime: user.CreateTime,
		UpdateTime: user.UpdateTime,
	}
}

// BuildUserList user list serializer func
func BuildUserList(items []models.User) (users []UserSerializer) {
	for _, item := range items {
		user := BuildUser(item)
		users = append(users, user)
	}
	return users
}

// BuildUserResponse for return data
func BuildUserResponse(user models.User) serializers.Response {
	return serializers.Response{
		Data: BuildUser(user),
	}
}
