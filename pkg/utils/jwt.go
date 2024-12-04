package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(userID, clientIP, secret string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":   userID,
		"client_ip": clientIP,
		"exp":       time.Now().Add(time.Hour).Unix(),
		"issued_at": time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString([]byte(secret))
}
