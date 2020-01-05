package models

// Permission for permission table
type Permission struct {
	ID             int    `gorm:"primary_key;column:id"`
	PermissionName string `gorm:"type:varchar(32);not null;column:permission_name"`
	CodeName       string `gorm:"type:varchar(32);not null;column:code_name"`
}

// TableName Permission Table Name
func (p Permission) TableName() string {
	return "sys_permission"
}
