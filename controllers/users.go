package controller

import (
	"encoding/json"
	"net/http"

	"example.com/questions/models"
	"example.com/questions/utils"
	"github.com/gin-gonic/gin"
)

func RegisterUser(context *gin.Context) {

	newUser := models.User{}
	err := context.ShouldBindJSON(&newUser)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Data provided"})
		return
	}

	err = newUser.AddUser()

	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			context.JSON(customErr.StatusCode, gin.H{"error": customErr.Message})
			return
		}
		context.JSON(500, gin.H{"error": "internal server error"})
		return
	}
}

func LoginUser(context *gin.Context) {
	loggedInUser := models.User{}
	data, err := context.GetRawData()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Data provided"})
		return
	}
	err = json.Unmarshal(data, loggedInUser)
	context.JSON(http.StatusOK, gin.H{"msg": "successful login"})
}
