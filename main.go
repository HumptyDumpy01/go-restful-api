package main

import (
	"HumptyDumpy01/go-restful-api/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	app := gin.Default()
	app.GET("/events", getEvents)
	err := app.Run()
	if err != nil {
		fmt.Println("Failed to run a server!")
		fmt.Println(err)
		return
	}
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, models.GetAllEvents())
}
