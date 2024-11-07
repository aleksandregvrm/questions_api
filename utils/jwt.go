package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func GenerateToken(username string) (string, error) {
	err := godotenv.Load()

	if err != nil {
		return "", err
	}

	var SecretKey string = os.Getenv("SECRET_KEY")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 3).Unix(),
	})
	return token.SignedString([]byte(SecretKey))
}
