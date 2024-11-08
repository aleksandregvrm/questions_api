package controller

import (
	"encoding/json"
	"net/http"

	"example.com/questions/models"
	"example.com/questions/utils"
	"github.com/gin-gonic/gin"
)

func RegisterUser(context *gin.Context) {

	newUser := &models.User{}
	data, err := context.GetRawData()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Data provided"})
	}
	err = json.Unmarshal(data, newUser)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error with processing data"})
	}

	err = models.AddUser(newUser)

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

}
