package services

import (
	"errors"

	"github.com/astaxie/beego/validation"
	"github.com/yhkl-dev/turtle-dove-beego/turtle-api/app/tables"
)

type authService struct{}

func (as *authService) table() string {
	return tableName("auth_token")
}
func (as *authService) valid(authToken *tables.AuthToken) error {
	valid := validation.Validation{}

	b, _ := valid.Valid(authToken)
	if !b {
		for _, err := range valid.Errors {
			return errors.New(err.Message)
		}
	}
	return nil
}

func (as *authService) IsExist(userId int) (*tables.AuthToken, error) {
	authToken := &tables.AuthToken{}
	authToken.UserId = userId
	err := o.Read(authToken, "UserId")
	return authToken, err
}

func (as *authService) AddToken(userId int, token string) (*tables.AuthToken, error) {
	// 后续可以使用redis 等
	// 前期先使用mysql 存储

	authToken, err := as.IsExist(userId)
	if err != nil {
		authToken = &tables.AuthToken{}
		authToken.UserId = userId
		authToken.Token = token
		_, err := o.Insert(authToken)
		return authToken, err
	}
	authToken.Token = token
	_, err = o.Update(authToken, "Token")
	return authToken, err
}
