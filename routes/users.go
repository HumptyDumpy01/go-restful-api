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
