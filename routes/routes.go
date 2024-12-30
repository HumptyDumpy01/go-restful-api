package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(app *gin.Engine) {
	app.GET("/events", getEvents)
	app.GET("/events/:id", getEvent)
	app.POST("/events", createEvent)
}
