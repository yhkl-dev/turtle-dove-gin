package utils

import (
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
)

const (
	// KEY encrypt string
	KEY string = "JWT"
	// DEFAULT_EXPIRE_SECONDS expire seconds of token
	DEFAULT_EXPIRE_SECONDS int = 600
)

// CustomClaims struct for jwt
type CustomClaims struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

// RefreshToken update expireAtTime and retunr a new string
func RefreshToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(KEY), nil
		})
	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return "", err
	}

	nowTime := time.Now()
	confExpireTime := beego.AppConfig.String("EXPIRE_TIME")

	var expireAtTime time.Time
	if len(confExpireTime) == 0 {
		expireAtTime = nowTime.Add(time.Duration(DEFAULT_EXPIRE_SECONDS) * time.Second)
	}
	midTime, _ := strconv.Atoi(confExpireTime)
	expireAtTime = nowTime.Add(time.Duration(midTime) * time.Second)
	newClaims := CustomClaims{
		claims.UserName,
		claims.Email,
		jwt.StandardClaims{
			ExpiresAt: expireAtTime.Unix(),
			Issuer:    claims.UserName,
		},
	}
	newTokenClaims := jwt.NewWithClaims(jwt.SigningMethodES256, newClaims)
	newToken, err := newTokenClaims.SignedString([]byte(KEY))
	return newToken, err
}

// GenerateToken return a token
func GenerateToken(userName, email string) (string, error) {
	nowTime := time.Now()
	confExpireTime := beego.AppConfig.String("EXPIRE_TIME")
	var expireTime time.Time
	if len(confExpireTime) == 0 {
		expireTime = nowTime.Add(time.Duration(DEFAULT_EXPIRE_SECONDS) * time.Second)
	}
	midTime, _ := strconv.Atoi(confExpireTime)
	expireTime = nowTime.Add(time.Duration(midTime) * time.Second)

	claims := CustomClaims{
		userName,
		email,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    userName,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(KEY))

	return token, err
}

// ParseToken parse jwt token; return a claims
func ParseToken(token string) (*CustomClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(
		token,
		&CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(KEY), nil
		},
	)

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
