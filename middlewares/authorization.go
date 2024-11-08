package authorize

import (
	"net/http"

	"example.com/questions/utils"
	"github.com/gin-gonic/gin"
)

func AuthorizeUser(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "Not authorized"})
		return
	}

	userId, err := utils.IsTokenValid(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "Not authorized"})
	}

	context.Set("userId", userId)
	context.Next()
}
