package middlewares

import (
	"net/http"

	"example.com/gin/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Authorization token is required",
		})
		return
	}

	userId, err := utils.ValidateJWT(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	context.Set("userId", userId)

	context.Next()
}
