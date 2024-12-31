package routes

import (
	"HumptyDumpy01/go-restful-api/models"
	"fmt"
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
	// this prop is available because of Authenticate middleware
	userId := context.GetInt64("userId")

	newEvent := models.Event{}
	err := context.ShouldBindJSON(&newEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "error", "data": gin.H{"error": "Invalid Input"}})
		return
	}

	// Set the UserId to the newEvent
	newEvent.UserId = userId

	err = newEvent.Save()
	fmt.Println(`Executing newEvent.UserId`, newEvent.UserId)

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

	userId := context.GetInt64("userId")
	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "data": gin.H{"error": "Failed to fetch the event!"}})
		return
	}

	if event.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"status": "error", "data": gin.H{"error": "Access denied."}})
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

	userId := context.GetInt64("userId")

	if event.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"status": "error", "data": gin.H{"error": "Access denied."}})
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
