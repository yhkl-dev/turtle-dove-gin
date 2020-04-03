package models

import "time"

// Resource
type Resource struct {
	ID          int        `gorm:"primary_key;column:id"`
	Host        string     `gorm:"type:varchar(40);not null;unique_index:idx_host_unique;column:host"`
	ManageIP    string     `gorm:"type:varchar(20);not null;column:ip"`
	IntraNetIP  string     `gorm:"type:varchar(20);not null;column:ip"`
	Description string     `gorm:"type:varchar(200);not null;column:description"`
	CreateTime  *time.Time `gorm:"column:create_time"`
}
