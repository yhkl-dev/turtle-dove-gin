package role

import (
	"time"

	"github.com/yhkl-dev/turtle-dove-beego/gin-turtle-api/apps/permission"
)

// Role for role table
type Role struct {
	ID           int                     `gorm:"primary_key;column:id"`
	ParentRoleID int                     `gorm:"column:parent_role_id; default(0)"`
	RoleName     string                  `gorm:"type:varchar(32);column:role_name;"`
	Description  string                  `gorm:"type:varchar(200);column:description"`
	CreateTime   time.Time               `gorm:"column:create_time"`
	UpdateTime   time.Time               `gorm:"column:update_time"`
	Permissions  []permission.Permission `gorm:"many2many:sys_role_permission_mapping"`
}

func (r Role) TableName() string {
	return "sys_role"
}
