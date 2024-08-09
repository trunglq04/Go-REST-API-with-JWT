package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "yoursecretkey"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(), // 2 hours
	})

	return token.SignedString([]byte(secretKey)) 	// convert secretKey to slice of byte
}
