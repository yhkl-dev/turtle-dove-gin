package tables

import "time"

// Permission Table
type Permission struct {
	Id             int
	PermissionName string    `orm:"size(32)"`
	CodeName       string    `orm:"size(32)"`
	CreateTime     time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime     time.Time `orm:"auto_now:type(datetime)"`
}
