package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	app := gin.Default()
	app.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong!",
		})
	})
	app.Run()
}
