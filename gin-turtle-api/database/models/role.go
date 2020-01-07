package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Role for role table
type Role struct {
	ID           int          `gorm:"primary_key;column:id"`
	ParentRoleID int          `gorm:"column:parent_role_id; default(0)"`
	RoleName     string       `gorm:"type:varchar(32);column:role_name;"`
	Description  string       `gorm:"type:varchar(200);column:description"`
	CreateTime   *time.Time   `gorm:"column:create_time"`
	UpdateTime   *time.Time   `gorm:"column:update_time"`
	Permissions  []Permission `gorm:"many2many:sys_role_permission_mapping"`
}

func (r Role) TableName() string {
	return "sys_role"
}

// Permissions
func rolePermissions() map[string]string {
	var permissionList = make(map[string]string)
	permissionList["ListAllRoles"] = "GET:/api/v1/role"
	permissionList["AddRole"] = "GET:/api/v1/role/:id "
	permissionList["UpdateRole"] = "PUT:/api/v1/role/:id"
	permissionList["DeleteRole"] = "DELETE:/api/v1/role/:id"
	return permissionList
}

// RegisterPermission 权限注册
func (u Role) RegisterPermission(DB *gorm.DB) {
	for cn, pn := range rolePermissions() {
		var initList Permission
		initList.CodeName = cn
		initList.PermissionName = pn
		DB.Create(&initList)
	}

}
