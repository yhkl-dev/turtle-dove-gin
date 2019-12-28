package services

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/yhkl-dev/turtle-dove-beego/turtle-api/app/tables"
)

type roleService struct{}

func (rs *roleService) table() string {
	return tableName("role")
}

func (rs *roleService) valid(role *tables.Role) error {
	valid := validation.Validation{}
	b, _ := valid.Valid(role)
	if !b {
		for _, err := range valid.Errors {
			return errors.New(err.Message)
		}
	}
	return nil
}

func (rs *roleService) GetRoleByID(roleID int) (*tables.Role, error) {
	role := &tables.Role{}
	role.Id = roleID

	err := o.Read(role)
	return role, err
}

func (rs *roleService) GetTotal() (int64, error) {
	return o.QueryTable(rs.table()).Count()
}

func (rs *roleService) GetRoleList(page, pageSize int) ([]orm.Params, error) {
	offset := (page - 1) * pageSize
	if offset < 0 {
		offset = 0
	}

	var roles []orm.Params
	queryset := o.QueryTable(rs.table())
	_, err := queryset.OrderBy("Id").Limit(pageSize, offset).Values(&roles, "Id", "ParentRoleId", "RoleName", "CreateTime", "UpdateTime", "Description")

	return roles, err
}

func (rs *roleService) AddRole(roleName, description string, parentRoleId int) (*tables.Role, error) {
	role := &tables.Role{}
	role.RoleName = roleName
	role.Description = description
	role.ParentRoleId = parentRoleId
	role.UpdateTime = time.Now()

	if err := rs.valid(role); err != nil {
		return nil, err
	}
	_, err := o.Insert(role)
	return role, err
}
