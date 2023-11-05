package token

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

// type jwtAuth struct {
// 	adminSecurityKey string
// 	userSecurityKey  string
// }

// func NewTokenService(adminAuthKey string, userAuthKey string) TokenService {
// 	return &jwtAuth{
// 		adminSecurityKey: adminAuthKey,
// 		userSecurityKey:  userAuthKey,
// 	}
// }

func GenerateToken(securityKey string, id string) string{
	claims := jwt.MapClaims{
		"id": id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(securityKey)
	if err != nil {
		fmt.Println(err, "error at create token ")
	}
	return tokenString
}

func VerifyToken(token string) {

}
