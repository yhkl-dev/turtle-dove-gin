package permission

import (
	"github.com/yhkl-dev/turtle-dove-gin/gin-turtle-api/database"
	"github.com/yhkl-dev/turtle-dove-gin/gin-turtle-api/database/models"
	"github.com/yhkl-dev/turtle-dove-gin/gin-turtle-api/utils/serializers"
)

// ListPermissionsService struct
type ListPermissionsService struct {
	PermissionName string `form:"permission_name"`
	Page           int    `form:"page"`
	PageSize       int    `form:"page_size"`
}

// ListAllPermissions func to show permissions list
func (lps *ListPermissionsService) ListAllPermissions() serializers.Response {
	var permissions []models.Permission
	total := 0
	queryString := "%" + lps.PermissionName + "%"

	if lps.PageSize == 0 {
		lps.PageSize = 10
	}

	if err := database.DB.Model(models.Permission{}).Where("permission_name like ?", queryString).Count(&total).Error; err != nil {
		return serializers.Response{
			Code:    50000,
			Message: "Query database error",
			Error:   err.Error(),
		}
	}
	if err := database.DB.Limit(lps.PageSize).Offset(lps.Page).Where("permission_name like ?", queryString).Find(&permissions).Error; err != nil {
		return serializers.Response{
			Code:    50000,
			Message: "Query database error",
			Error:   err.Error(),
		}
	}
	return serializers.BuildListResponse(BuildPermissionList(permissions), total)

}
