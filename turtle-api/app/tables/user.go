package tables

import "time"

// User Table
type User struct {
	Id           int
	UserName     string    `orm:"unique;size(32)" valid:"Required"`
	UserPassword string    `orm:"size(32)" valid:"MinSize(8)"`
	RealName     string    `orm:"size(30)" valid:"Required"`
	Email        string    `orm:"size(50)" valid:"Email"`
	IsActive     int       `orm:"default(0)"`
	LastLogin    time.Time `orm:"null;type(datetime)"`
	CreateTime   time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime   time.Time `orm:"auto_now:type(datetime)"`
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
