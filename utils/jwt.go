package utils

import (
	"errors"
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

	var secretKey string = os.Getenv("SECRET_KEY")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 3).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func IsTokenValid(token string) (int64, error) {
	var secretKey string = os.Getenv("SECRET_KEY")
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Unexpected sign in method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, errors.New("Could not parse the token")
	}

	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return 0, errors.New("Invalid token")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("Invalid token claims")
	}

	userId := int64(claims["userId"].(float64))

	return userId, nil
}
