package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/larryokubasu5460/gomart/user-service/config"
)

func GenerateJWT(userID uint, email string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":userID,
		"email":email,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Cfg.JWTSecret))
}