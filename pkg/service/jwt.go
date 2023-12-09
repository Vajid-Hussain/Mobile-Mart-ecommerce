package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func TemperveryTokenForOtpVerification(securityKey string, phone string) (string, error) {
	key := []byte(securityKey)
	claims := jwt.MapClaims{
		"exp":   time.Now().Unix() + 3000,
		"Phone": phone,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		fmt.Println(err, "error at create token ")
	}
	return tokenString, err
}

func GenerateAcessToken(securityKey string, id string) (string, error) {
	key := []byte(securityKey)
	claims := jwt.MapClaims{
		"exp": time.Now().Unix() + 300,
		"id":  id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		fmt.Println(err, "error at create token ")
	}
	return tokenString, err
}

func GenerateRefreshToken(securityKey string) (string, error) {
	key := []byte(securityKey)
	clamis := jwt.MapClaims{
		"exp": time.Now().Unix() + 3600000,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clamis)
	signedToken, err := token.SignedString(key)
	if err != nil {
		return "", errors.New("making refresh token lead to error")
	}

	return signedToken, nil
}

func VerifyAcessToken(token string, secretkey string) (string, error) {
	key := []byte(secretkey)
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if parsedToken == nil {
		return "", err
	}

	claims := parsedToken.Claims.(jwt.MapClaims)
	id, ok := claims["id"].(string)

	// defer func() {
	// 	if err != nil {
	// 		fmt.Println("$$$$$$", id, err)
	// 	}
	// 	err = nil
	// }()

	if !ok {
		return "", errors.New("id is not in accessToken. access denied")
	}
	return id, err
}

func VerifyRefreshToken(token string, securityKey string) error {
	key := []byte(securityKey)

	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return errors.New(" token tamperd or expired")
	}

	return nil
}

// func FetchPhoneFromToken(tokenString string, secretkey string) (string, error) {
// 	secret := []byte(secretkey)
// 	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		return secret, nil
// 	})
// 	if err != nil || !parsedToken.Valid {
// 		fmt.Println(err, "wronge user with wrong token")
// 		return "", errors.New("wrong token or expired")
// 	}
// 	claims := parsedToken.Claims.(jwt.MapClaims)
// 	phone := claims["Phone"].(string)

// 	return phone, nil
// }

func FetchPhoneFromToken(tokenString string, secretkey string) (string, error) {

	secret := []byte(secretkey)

	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil || !parsedToken.Valid {
		return "", errors.New("wrong token or expired")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("could not parse claims")
	}

	phoneClaim, ok := claims["Phone"].(string)
	if !ok {
		return "", errors.New("phone claim not found or not a string")
	}

	return phoneClaim, nil
}
