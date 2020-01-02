package role

import (
	"fmt"
	"time"

	"github.com/yhkl-dev/turtle-dove-beego/gin-turtle-api/database"
	"github.com/yhkl-dev/turtle-dove-beego/gin-turtle-api/database/models"
	"github.com/yhkl-dev/turtle-dove-beego/gin-turtle-api/utils/serializers"
)

// AddRoleService add role form
type AddRoleService struct {
	ParentRoleID int    `form:"parent_role_id" json:"parent_role_id"`
	RoleName     string `form:"role_name" json:"role_name" binding:"required,max=30"`
	Description  string `form:"description" json:"description" binding:"required,max=200"`
}

// valid function for form
func (ads *AddRoleService) valid() *serializers.Response {
	isExist := 0
	database.DB.Model(&models.Role{}).Where("role_name = ?", ads.RoleName).Count(&isExist)
	if isExist > 0 {
		return &serializers.Response{
			Code:    40001,
			Message: fmt.Sprintf("role_name %s has exist", ads.RoleName),
		}

	}
	return nil

}

// AddRoleService add a role
func (ars *AddRoleService) AddRole() serializers.Response {
	role := models.Role{
		ParentRoleID: ars.ParentRoleID,
		RoleName:     ars.RoleName,
		Description:  ars.Description,
	}

	if err := ars.valid(); err != nil {
		return *err
	}

	nowTime := time.Now()
	role.CreateTime = &nowTime
	if err := database.DB.Create(&role).Error; err != nil {
		return serializers.ParameterError("add role failed", err)
	}
	return BuildRoleResponse(role)

}
