package routes

import (
	controller "example.com/questions/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(route *gin.Engine) {
	route.POST("/register", controller.RegisterUser)
	route.POST("/login")
}
