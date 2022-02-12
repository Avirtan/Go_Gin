package jwtHandler

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserDate struct {
	Id    string `json:"id"`
	Login string `json:"login"`
}

// поля для записи в токен
type authCustomClaims struct {
	UserDate
	jwt.StandardClaims
}

//метод на создание токена из данных пользователей
func CreateToken(date UserDate) string {
	claims := &authCustomClaims{
		date,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 168).Unix(),
			Issuer:    "work_craft",
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("SECRET")
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}
	return t
}

// метод на проверку токена на волидность и возвращения мапы с данными
func ValidateToken(token string) (jwt.MapClaims, bool) {
	tokenParse, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		secret := os.Getenv("SECRET")
		return []byte(secret), nil
	})
	if err != nil {
		return nil, false
	}
	if claims, ok := tokenParse.Claims.(jwt.MapClaims); ok && tokenParse.Valid {
		return claims, true
	} else {
		return nil, false
	}
}
