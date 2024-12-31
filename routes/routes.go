package routes

import (
	"HumptyDumpy01/go-restful-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(app *gin.Engine) {
	authenticated := app.Group("/")

	app.GET("/events", getEvents)
	app.GET("/events/:id", getEvent)

	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	////////////////////////////////////////
	app.POST("/signup", signup)
	app.POST("/login", login)
}
