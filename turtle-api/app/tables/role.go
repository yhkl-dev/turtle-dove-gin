package tables

import "time"

// Role table
type Role struct {
	Id           int
	ParentRoleId int       `orm:"default(0)" valid:"Required"`
	RoleName     string    `orm:"size(20)" valid:"Required;MinSize(3)"`
	CreateTime   time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime   time.Time `orm:"auto_now;type(datetime)"`
	Description  string    `orm:"size(200)" valid:"Required"`
}
