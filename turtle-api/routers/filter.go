package routers

import (
	"fmt"

	beego_context "github.com/astaxie/beego/context"
	"github.com/yhkl-dev/turtle-dove-beego/turtle-api/app/services"
	"github.com/yhkl-dev/turtle-dove-beego/turtle-api/utils"
)

// FilterToken func to validate token
var FilterToken = func(ctx *beego_context.Context) {

	if ctx.Request.RequestURI != "/v1/user/login" && ctx.Input.Header("Authorization") == "" {
		ctx.ResponseWriter.WriteHeader(401)
		ctx.ResponseWriter.Write([]byte("please login first"))
	}
	if ctx.Request.RequestURI != "/v1/user/login" && ctx.Input.Header("Authorization") != "" {
		token := ctx.Input.Header("Authorization")
		originClaims, err := utils.ParseToken(token)
		if err != nil {
			fmt.Println("xxxxxx")
			ctx.ResponseWriter.WriteHeader(401)
			ctx.ResponseWriter.Write([]byte(err.Error()))
		}
		fmt.Println(err)
		if originClaims != nil {
			isExist := services.RedisService.IsExist(originClaims.UserName)
			if !isExist {
				ctx.ResponseWriter.WriteHeader(401)
				ctx.ResponseWriter.Write([]byte("Authorization has been expired"))
			}

			rToken, _ := services.RedisService.Get(originClaims.UserName)

			if rToken != token {
				ctx.ResponseWriter.WriteHeader(401)
				ctx.ResponseWriter.Write([]byte("Authorization has been expired"))
			}
		}
	}
}
