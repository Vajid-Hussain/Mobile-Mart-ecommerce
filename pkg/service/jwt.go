package service

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
)



func TemperveryTokenForOtpVerification(securityKey string, phone string) (string, error) {
	key := []byte(securityKey)
	claims := jwt.MapClaims{
		"Phone": phone,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		fmt.Println(err, "error at create token ")
	}
	return tokenString, err
}

func GenerateToken(securityKey string, id string) (string, error) {
	Kye := []byte(securityKey)
	claims := jwt.MapClaims{
		"id": id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(Kye)
	if err != nil {
		fmt.Println(err, "error at create token ")
	}
	return tokenString, err
}

func VerifyToken(token string, secretkey string) (string, error){
	key:= []byte(secretkey)
	parsedToken, err:=jwt.Parse(token , func(token *jwt.Token) (interface{}, error){
		return key, nil
	})
	if err!=nil{
		return "", errors.New("wrong token structure")
	}
	claims:= parsedToken.Claims.(jwt.MapClaims)
	id:=claims["id"].(string)

	return id, nil
}

func FetchPhoneFromToken(tokenString string, secretkey string) (string, error) {
	secret := []byte(secretkey)
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil || !parsedToken.Valid {
		fmt.Println(err, "wronge user with wrong token")
		return "", errors.New("wrong token or expired")
	}
	claims := parsedToken.Claims.(jwt.MapClaims)
	phone := claims["Phone"].(string)

	return phone, nil
}
