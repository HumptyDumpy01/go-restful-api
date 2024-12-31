package middlewares

import (
	"HumptyDumpy01/go-restful-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "error", "data": gin.H{"error": "Not Authorized"}})
		return
	}

	// Check if the token starts with "Bearer "
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	} else {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "error", "data": gin.H{"error": "Malformed token or expired 1."}})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "error", "data": gin.H{"error": "Malformed token or expired 2."}})
		return
	}

	context.Set("userId", userId)

	context.Next()
}
