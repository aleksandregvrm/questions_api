package routes

import (
	"example.com/questions/models"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(route *gin.Engine) {
	route.POST("/add", models.AddUser)
}
