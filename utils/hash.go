package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	fmt.Println("Password hashed")
	return string(bytes), err
}

func ComparePasswords(candidatePassword, realPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(realPassword), []byte(candidatePassword))
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
