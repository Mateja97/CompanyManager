package auth

import (
	"errors"
	"flag"
	"log"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecretKey = flag.String("secretKey", "", "JWT secret key for authroization")

//VerifyToken verfies token and return error
func VerifyToken(tokenString string) error {
	if tokenString == "" {
		return errors.New("token is missing")
	}
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(*jwtSecretKey), nil
	})
	if err != nil {
		log.Println("[Warn] jwt not authorized, token", tokenString, err)
		return err
	}
	return nil
}
