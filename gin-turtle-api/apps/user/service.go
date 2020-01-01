package user

import (
	"fmt"
	"time"

	"github.com/yhkl-dev/turtle-dove-beego/gin-turtle-api/database"
	"github.com/yhkl-dev/turtle-dove-beego/gin-turtle-api/database/models"
	"github.com/yhkl-dev/turtle-dove-beego/gin-turtle-api/utils/encrypt"
	"github.com/yhkl-dev/turtle-dove-beego/gin-turtle-api/utils/serializers"
)

// UserListService for return data for ListAllUsers
type UserListService struct {
	UserName string `form:"user_name" query:"user_name"`
	PageSize int    `form:"page_size"`
	Page     int    `form:"page"`
}

// ListAllUsers function return user list
func (uls *UserListService) ListAllUsers() serializers.Response {
	var users []models.User
	total := 0
	queryUser := "%" + uls.UserName + "%"

	if uls.PageSize == 0 {
		uls.PageSize = 10
	}

	if err := database.DB.Model(models.User{}).Where("is_deleted = 0 and user_name like  ?", queryUser).Count(&total).Error; err != nil {
		return serializers.Response{
			Code:    50000,
			Message: "Query database error",
			Error:   err.Error(),
		}
	}

	if err := database.DB.Limit(uls.PageSize).Offset(uls.Page).Where("is_deleted = 0 and user_name like  ?", queryUser).Find(&users).Error; err != nil {
		return serializers.Response{
			Code:    50000,
			Message: "Error connect database",
			Error:   err.Error(),
		}
	}

	return serializers.BuildListResponse(BuildUserList(users), total)
}

// UserRegieterService construct for user register
type UserRegieterService struct {
	UserName     string `form:"user_name" json:"user_name" binding:"required,min=8,max=32"`
	UserPassword string `form:"user_password" json:"user_password" binding:"required,min=8,max=32"`
	RealName     string `form:"real_name" json:"real_name" binding:"required,min=2,max=32"`
	Email        string `form:"email" json:"email" binding:"required"`
}

// UserRegieterService for user register
func (urs *UserRegieterService) valid() *serializers.Response {
	isExist := 0
	database.DB.Model(&models.User{}).Where("user_name = ?", urs.UserName).Count(&isExist)
	if isExist > 0 {
		return &serializers.Response{
			Code:    40001,
			Message: fmt.Sprintf("user_name %s has exist", urs.UserName),
		}
	}
	return nil
}

// UserRegieterService enroll service
func (urs *UserRegieterService) Register() serializers.Response {
	user := models.User{
		UserName: urs.UserName,
		RealName: urs.RealName,
		Email:    urs.Email,
	}

	if err := urs.valid(); err != nil {
		return *err
	}

	user.UserPassword = encrypt.StringToMd5(urs.UserPassword)

	nowTime := time.Now()
	user.CreateTime = &nowTime
	if err := database.DB.Create(&user).Error; err != nil {
		return serializers.ParameterError("Enroll Failed", err)
	}
	return BuildUserResponse(user)
}

// GetUserProfileService retuan user profile
type GetUserProfileService struct {
	UserName   string     `json:"user_name"`
	RealName   string     `json:"real_name"`
	Email      string     `json:"email"`
	LastLogin  *time.Time `json:"last_login"`
	CreateTime *time.Time `json:"create_time"`
	UpdateTime *time.Time `json:"update_time"`
}

// GetUserProfile return user profile
func (gups *GetUserProfileService) GetUserProfile(id string) serializers.Response {
	var user models.User
	err := database.DB.First(&user, id).Error
	if err != nil {
		return serializers.Response{
			Code:    404,
			Message: "user does not exist",
			Error:   err.Error(),
		}
	}
	return serializers.Response{
		Data: BuildUser(user),
	}
}

// UpdateUserService update user profile
type UpdateUserService struct {
	UserPassword string `form:"user_password" json:"user_password"`
	RealName     string `form:"real_name" json:"real_name"`
	Email        string `form:"email" json:"email"`
	IsActive     int    `form:"is_active" json:"is_active"`
}

func (ups *UpdateUserService) UpdateUser(id string) serializers.Response {

	var user models.User

	err := database.DB.First(&user, id).Error
	if err != nil {
		return serializers.Response{
			Code:    404,
			Message: "user does not exist",
			Error:   err.Error(),
		}
	}

	user.RealName = ups.RealName
	user.UserPassword = encrypt.StringToMd5(ups.UserPassword)
	user.Email = ups.Email
	user.IsActive = ups.IsActive

	updateTime := time.Now()
	user.UpdateTime = &updateTime

	err = database.DB.Save(&user).Error
	if err != nil {
		return serializers.Response{
			Code:    50002,
			Message: "update user profile failed",
			Error:   err.Error(),
		}
	}
	return serializers.Response{
		Data: BuildUser(user),
	}
}

// DeleteUserService struct for delete user
type DeleteUserService struct {
	IsDeleted int `json:"is_deleted"`
}

// DeleteUser set uset is_deleted true
func (dus *DeleteUserService) DeleteUser(id string) serializers.Response {

	var user models.User

	err := database.DB.First(&user, id).Error
	if err != nil {
		return serializers.Response{
			Code:    404,
			Message: "user does not exist",
			Error:   err.Error(),
		}
	}

	user.IsDeleted = dus.IsDeleted

	updateTime := time.Now()
	user.UpdateTime = &updateTime

	err = database.DB.Save(&user).Error
	if err != nil {
		return serializers.Response{
			Code:    50002,
			Message: "update user profile failed",
			Error:   err.Error(),
		}
	}
	return serializers.Response{
		Data: "success",
	}

}

type UserLoginService struct {
	UserName     string `form:"user_name" json:"user_name"`
	UserPassword string `form:"user_password" json:"user_password"`
}

func (uls *UserLoginService) Login() serializers.Response {

	var user models.User

	err := database.DB.Where("user_name = ?", uls.UserName).First(&user).Error
	if err != nil {
		return serializers.Response{
			Code:    404,
			Message: "user_name or password incorrect.",
			Error:   err.Error(),
		}

	}
	if user.UserPassword == encrypt.StringToMd5(uls.UserPassword) {
		fmt.Println("passed")
	}

	return serializers.Response{
		Data: "success",
	}
}