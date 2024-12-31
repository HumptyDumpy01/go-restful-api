package routes

import (
	"HumptyDumpy01/go-restful-api/models"
	"HumptyDumpy01/go-restful-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"data": gin.H{
				"error": "Failed to parse the id",
			},
		})
		return
	}
	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"data": gin.H{
				"error": "Failed to fetch event by id!",
			},
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   event,
	})
}

func getEvents(context *gin.Context) {
	data, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"data":   gin.H{"error": "Failed to fetch the events."},
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   data,
	})
}

func createEvent(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"status": "error", "data": gin.H{"error": "Not Authorized"}})
		return
	}

	err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"status": "error", "data": gin.H{"error": "Malformed token or expired."}})
		return
	}

	newEvent := models.Event{}
	err = context.ShouldBindJSON(&newEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "error", "data": gin.H{"error": "Invalid Input"}})
		return
	}

	err = newEvent.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"status": "error", "data": gin.H{"error": "Failed to save event."}})
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": "success", "data": newEvent})
}

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "fail", "data": gin.H{"error": "Failed to parse the float"}})
		return
	}

	_, err = models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "data": gin.H{"error": "Failed to fetch the event!"}})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "fail", "data": gin.H{"error": "Failed to parse user data."}})
		return
	}

	updatedEvent.ID = id
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "data": gin.H{"error": "Failed to perform the update."}})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{}})
}

func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "fail", "data": gin.H{"error": "Failed to parse the id!"}})
		return
	}

	event, err := models.GetEventById(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"status": "error", "data": gin.H{"error": "Failed to fetch event by id."}})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"status": "error", "data": gin.H{"error": "Failed to delete event by id."}})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   gin.H{},
	})
}
