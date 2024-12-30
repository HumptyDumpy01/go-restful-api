package main

import (
	"HumptyDumpy01/go-restful-api/db"
	"HumptyDumpy01/go-restful-api/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// initialize the database.
	db.InitDB()

	app := gin.Default()
	app.GET("/events", getEvents)
	app.POST("/events", createEvent)
	err := app.Run()
	if err != nil {
		fmt.Println("Failed to run a server!")
		fmt.Println(err)
		return
	}
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
	newEvent := models.Event{}
	err := context.ShouldBindJSON(&newEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"data":   gin.H{"error": "Invalid Input"},
		})
		return
	}

	err = newEvent.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"data":   gin.H{"error": "Failed to save event."},
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   newEvent,
	})
}
