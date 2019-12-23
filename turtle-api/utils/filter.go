package utils

import (
	"fmt"

	beego_context "github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"github.com/yhkl-dev/turtle-dove-beego/turtle-api/app/tables"
)

// FilterToken func to validate token
var FilterToken = func(ctx *beego_context.Context) {

	if ctx.Request.RequestURI != "/v1/user/login" && ctx.Input.Header("Authorization") == "" {
		ctx.ResponseWriter.WriteHeader(401)
		ctx.ResponseWriter.Write([]byte("please login first"))
	}
	if ctx.Request.RequestURI != "/v1/user/login" && ctx.Input.Header("Authorization") != "" {
		token := ctx.Input.Header("Authorization")
		o := orm.NewOrm()

		authToken := &tables.AuthToken{}
		authToken.Token = token
		err := o.Read(authToken, "Token")
		if err != nil {
			ctx.ResponseWriter.WriteHeader(401)
			ctx.ResponseWriter.Write([]byte("Authorization has been expired"))
		}
		originClaims, _ := ParseToken(token)
		fmt.Println(originClaims.UserName)

	}
}
