package role

import (
	"time"

	"github.com/yhkl-dev/turtle-dove-beego/gin-turtle-api/database/models"
	"github.com/yhkl-dev/turtle-dove-beego/gin-turtle-api/utils/serializers"
)

// RoleSerializer struct for return data
type RoleSerializer struct {
	ID           int        `form:"id" json:"id"`
	ParentRoleID int        `form:"parent_id" json:"parent_id"`
	RoleName     string     `form:"role_name" json:"role_name"`
	Description  string     `form:"description" json:"description"`
	CreateTime   *time.Time `form:"create_time" json:"create_time"`
	UpdateTime   *time.Time `form:"update_time" json:"update_time"`
}

// BuildRole single role serialzier
func BuildRole(role models.Role) RoleSerializer {
	return RoleSerializer{
		ID:           role.ID,
		ParentRoleID: role.ParentRoleID,
		RoleName:     role.RoleName,
		Description:  role.Description,
		CreateTime:   role.CreateTime,
		UpdateTime:   role.UpdateTime,
	}
}

// BuildRoleList role list serializer func
func BuildRoleList(items []models.Role) (roles []RoleSerializer) {
	for _, item := range items {
		role := BuildRole(item)
		roles = append(roles, role)

	}
	return roles

}

// BuildRoleResponse for return data
func BuildRoleResponse(role models.Role) serializers.Response {
	return serializers.Response{
		Data: BuildRole(role),
	}

}