package jwtauth

import (
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/yhkl-dev/turtle-dove-beego/gin-turtle-api/conf"
)

const (
	secret     string = "yhkl" // encrypt salt
	expireTime int    = 3600
)

// JWTClaims
type JWTClaims struct {
	jwt.StandardClaims
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
}

// CustomClaims struct for jwt
type CustomClaims struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func getExpireTime() (confExpireTime int) {
	envExpireTime, err := conf.SysConfig.GetValue("Authentication", "token_expire_time")
	if err != nil {
		confExpireTime = expireTime
	}
	midTime, err := strconv.Atoi(envExpireTime)
	if err != nil {
		confExpireTime = expireTime
	}
	confExpireTime = midTime
	return
}

// GenerateToken return a token
func GenerateToken(userName, email string) (string, error) {
	claims := CustomClaims{
		userName,
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(getExpireTime()) * time.Second).Unix(),
			Issuer:    userName,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(secret))

	return token, err

}

// ParseToken parse jwt token; return a claims
func ParseToken(token string) (*CustomClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(
		token,
		&CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil

		},
	)

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
			return claims, nil

		}

	}
	return nil, err
}

// RefreshToken update expireAtTime and retunr a new string
func RefreshToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil

		},
	)
	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return "", err

	}

	nowTime := time.Now()
	newClaims := CustomClaims{
		claims.UserName,
		claims.Email,
		jwt.StandardClaims{
			ExpiresAt: nowTime.Add(time.Duration(getExpireTime()) * time.Second).Unix(),
			Issuer:    claims.UserName,
		},
	}
	newTokenClaims := jwt.NewWithClaims(jwt.SigningMethodES256, newClaims)
	newToken, err := newTokenClaims.SignedString([]byte(secret))
	return newToken, err
}
