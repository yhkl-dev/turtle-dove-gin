package permission

import "github.com/yhkl-dev/turtle-dove-gin/gin-turtle-api/database/models"

type PermissionSerializer struct {
	ID             int    `json:"id"`
	PermissionName string `json:"permission_name"`
	CodeName       string `json:"code_name"`
}

// BuildPermission func
func BuildPermission(permission models.Permission) PermissionSerializer {
	return PermissionSerializer{
		ID:             permission.ID,
		PermissionName: permission.PermissionName,
		CodeName:       permission.CodeName,
	}
}

// BuildPermissionList return permission list
func BuildPermissionList(items []models.Permission) (permissions []PermissionSerializer) {
	for _, item := range items {
		permission := BuildPermission(item)
		permissions = append(permissions, permission)
	}
	return permissions
}
