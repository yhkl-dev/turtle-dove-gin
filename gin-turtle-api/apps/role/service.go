package role

import (
	"fmt"
	"time"

	"github.com/yhkl-dev/turtle-dove-beego/gin-turtle-api/database"
	"github.com/yhkl-dev/turtle-dove-beego/gin-turtle-api/database/models"
	"github.com/yhkl-dev/turtle-dove-beego/gin-turtle-api/utils/serializers"
)

// RoleListService list all roles
type RoleListService struct {
	ParentRoleID int    `query:"parent_role_id" json:"parent_role_id"`
	RoleName     string `query:"role_name" json:"role_name"`
	PageSize     int    `query:"page_size"`
	Page         int    `query:"page"`
}

func (rls *RoleListService) ListAllRoles() serializers.Response {
	var roles []models.Role
	total := 0
	queryRole := "%" + rls.RoleName + "%"

	if rls.PageSize == 0 {
		rls.PageSize = 10
	}

	if err := database.DB.Model(models.Role{}).Where("role_name like ?", queryRole).Count(&total).Error; err != nil {
		return serializers.Response{
			Code:    50000,
			Message: "query database Error",
			Error:   err.Error(),
		}
	}
	if err := database.DB.Limit(rls.PageSize).Offset(rls.Page).Where("parent_role_id = ? and role_name like ?", rls.ParentRoleID, queryRole).Find(&roles).Error; err != nil {
		return serializers.Response{
			Code:    50000,
			Message: "error occurquery database Error",
			Error:   err.Error(),
		}
	}
	return serializers.BuildListResponse(BuildRoleList(roles), total)
}

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

// UpdateRoelService update user service struct
type UpdateRoelService struct {
	ParentRoleID int    `form:"parent_role_id" json:"parent_role_id"`
	RoleName     string `form:"role_name" json:"role_name"`
	Description  string `form:"description" json:"description"`
}

// UpdateRole update role info
func (urs *UpdateRoelService) UpdateRole(id string) serializers.Response {
	var role models.Role

	err := database.DB.First(&role, id).Error

	if err != nil {
		return serializers.Response{
			Code:    404,
			Message: "role does not exist",
			Error:   err.Error(),
		}
	}
	if urs.RoleName != "" {
		role.RoleName = urs.RoleName
	}
	if urs.Description != "" {
		role.Description = urs.Description
	}
	if urs.ParentRoleID != 0 {
		role.ParentRoleID = urs.ParentRoleID
	}

	updateTime := time.Now()
	role.UpdateTime = &updateTime

	err = database.DB.Save(&role).Error
	if err != nil {
		return serializers.Response{
			Code:    50002,
			Message: "update role failed",
			Error:   err.Error(),
		}
	}
	return serializers.Response{
		Data: BuildRole(role),
	}
}

// DeleteRoleService struct
type DeleteRoleService struct {
	RoleID int `form:"role_id" json:"role_id"`
}

// DeleteRole delete role and its relationship
func (drs *DeleteRoleService) DeleteRole(id string) serializers.Response {
	var role models.Role
	err := database.DB.First(&role, id).Error
	if err != nil {
		return serializers.Response{
			Code:    404,
			Message: "role does not exist.",
			Error:   err.Error(),
		}
	}
	err = database.DB.Model(models.User{}).Delete(&role).Error
	if err != nil {
		return serializers.Response{
			Code:    50001,
			Message: "error where delete user relation.",
			Error:   err.Error(),
		}
	}
	err = database.DB.Delete(&role).Error

	if err != nil {
		return serializers.Response{
			Code:    50001,
			Message: "error where delete user relation.",
			Error:   err.Error(),
		}
	}
	return serializers.Response{
		Data: "success",
	}
}
