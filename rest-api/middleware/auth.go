package middleware

import (
	"example/rest-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {

	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Missing authorization token",
			"code":    "REAPI401",
		})
		return
	}

	userId, err := utils.ValidateJWTToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
			"code":    "REAPI401",
		})
		return
	}
	context.Set("userId", userId)
	context.Next()
}
