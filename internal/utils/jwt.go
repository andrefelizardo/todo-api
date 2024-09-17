package utils

import (
	"time"

	"github.com/andrefelizardo/todo-api/internal/configs"
	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
)

var jwtSecret []byte

func GetJWTSecret() []byte {
	return jwtSecret
}

func InitJWTSecret(config *configs.JWTConfig) {
	jwtSecret = []byte(config.Secret)
}

type Claims struct {
	UserID string `json:"user_id"`
	IsVerified bool `json:"is_verified"`
	jwt.RegisteredClaims
}

func GenerateToken(userID string, isVerified bool) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		IsVerified: isVerified,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			Issuer: "todo-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		log.Error("Error generating token", err)
		return "", err
	}

	return tokenString, nil

}

func GenerateEmailConfirmationToken(userID string) (string, error) {
	return GenerateToken(userID, false)
}
	