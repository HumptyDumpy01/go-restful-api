package routes

import (
	"HumptyDumpy01/go-restful-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "fail", "data": gin.H{"error": "Failed to parse the float"}})
		return
	}

	event, err := models.GetEventById(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"status": "error", "data": gin.H{"error": "Failed to find the event."}})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"status": "error", "data": gin.H{"error": "Failed to register an event."}})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"status": "success", "data": event})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	var event models.Event
	event.ID = id

	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"status": "error", "data": gin.H{"error": "Failed to cancel an event."}})
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": "success", "data": event})
}
