package main

import (
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
	context.JSON(http.StatusOK, gin.H{
		"message": "pong!",
	})
}
