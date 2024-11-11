package controller

import (
	"fmt"
	"net/http"

	"example.com/questions/models"
	"github.com/gin-gonic/gin"
)

func RegisterUser(context *gin.Context) {

	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		fmt.Println("1", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.AddUser()

	if err != nil {
		fmt.Println("2", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func LoginUser(context *gin.Context) {
	var loggedInUser models.User

	err := context.ShouldBindJSON(&loggedInUser)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Data provided"})
		return
	}
	err = loggedInUser.LoginUser()

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Internal server error"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"msg": "successful login"})
}
