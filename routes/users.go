package routes

import (
	"HumptyDumpy01/go-restful-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "fail", "data": gin.H{"error": "Invalid input!"}})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"status": "error", "data": gin.H{"error": "Failed to create a user."}})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   user,
	})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"data":   gin.H{"error": "Failed to bind json!"},
		})
		return
	}
	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"status": "error",
			"data":   gin.H{"error": "Failed to validate credentials"},
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   "Authenticated",
	})
	return
}
