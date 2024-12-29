package main

import (
	"HumptyDumpy01/go-restful-api/db"
	"HumptyDumpy01/go-restful-api/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
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
	context.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   models.GetAllEvents(),
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
	newEvent.ID = rand.ExpFloat64()
	newEvent.UserId = rand.ExpFloat64()
	newEvent.DateTime = time.Now()
	newEvent.Save()

	context.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   newEvent,
	})
}
