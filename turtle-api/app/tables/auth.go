package tables

import "github.com/dgrijalva/jwt-go"

const (
	// KEY define key name
	KEY string = "JWT"
	// DEFAULT_EXPIRE_SECOND token expire time
	DEFAULT_EXPIRE_SECOND int = 600
)

// CustomClaims -- json web token
// HEADER PAYLOAD SIGNATURE
// This struct is the PAYLOAD
type CustomClaims struct {
	User User
	jwt.StandardClaims
}

// AuthToken table
type AuthToken struct {
	Id     int
	UserId int    `orm:"unique" valid:"Rquired"`
	Token  string `orm:"size(250)" valid:"Required"`
}
