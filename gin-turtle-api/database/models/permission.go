package models

import "time"

// Permission for permission table
type Permission struct {
	ID             int       `gorm:"primary_key;column:id"`
	PermissionName string    `gorm:"type:varchar(32);not null;column:permission_name"`
	CodeName       string    `gorm:"type:varchar(32);not null;column:code_name"`
	CreateTime     time.Time `gorm:"auto_now_add;column:create_time"`
	UpdateTime     time.Time `gorm:"auto_now;column:update_time"`
}

// TableName Permission Table Name
func (p Permission) TableName() string {
	return "sys_permission"
}
