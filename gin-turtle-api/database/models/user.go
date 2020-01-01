package models

import (
	"time"
)

// User struct for user table
type User struct {
	ID           int        `gorm:"primary_key;column:id"`
	UserName     string     `gorm:"type:varchar(32);not null;unique_index:idx_username_unique;column:user_name"`
	UserPassword string     `gorm:"type:varchar(40);not null;column:user_password"`
	RealName     string     `gorm:"type:varchar(32);not null;column:real_name"`
	Email        string     `gorm:"type:varchar(40);not null; column:email"`
	IsActive     int        `gorm:"default 0;column:is_active;comment:'是否禁用:0:否,1:是'"`
	IsDeleted    int        `gorm:"default 0;column:is_deleted; comment:'是否删除 0: 否, 1：是'"`
	LastLogin    *time.Time `gorm:"column:last_login"`
	CreateTime   *time.Time `gorm:"column:create_time"`
	UpdateTime   *time.Time `gorm:"column:update_time"`
	Roles        []Role     `gorm:"many2many:sys_user_role_mapping"`
}

// User table name
func (u User) TableName() string {
	return "sys_user"
}
