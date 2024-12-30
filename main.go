package main

import (
	"HumptyDumpy01/go-restful-api/db"
	"HumptyDumpy01/go-restful-api/routes"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	// initialize the database.
	db.InitDB()

	app := gin.Default()
	routes.RegisterRoutes(app)

	err := app.Run()
	if err != nil {
		fmt.Println("Failed to run a server!")
		fmt.Println(err)
		return
	}
}
