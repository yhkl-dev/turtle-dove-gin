package tables

import "time"

// User Table
type User struct {
	Id           int
	UserName     string    `orm:"unique;size(32)"`
	UserPassword string    `orm:"size(32)"`
	RealName     string    `orm:"size(30)"`
	IsActive     int       `orm:"default(0)"`
	LastLogin    time.Time `orm:"null;type(datetime)"`
	CreateTime   time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime   time.Time `orm:"auto_now:type(datetime)"`
}

// Role table
type Role struct {
	Id          int
	RoleName    string    `orm:"size(20)"`
	CreateTime  time.Time `orm:"auto_now_add; type(datetime)"`
	UpdateTime  time.Time `orm:"auto_now;type(datetime)"`
	Description string    `orm:"size(200)"`
}

// Permission Table
type Permission struct {
	Id             int
	PermissionName string    `orm:"size(32)"`
	CodeName       string    `orm:"size(32)"`
	CreateTime     time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime     time.Time `orm:"auto_now:type(datetime)"`
}

// User and Role Mapping table
type UserRole struct {
	Id     int
	UserId int
	RoleId int
}

// Role and Permission Mapping table
type RolePermission struct {
	Id           int
	RoleId       int
	PermissionId int
}
