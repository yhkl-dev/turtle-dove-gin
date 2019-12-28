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
	Roles        []*Role   `orm:"rel(m2m)"`
}
