package main

import (
	"gin-mongo-api/configs"
	"gin-mongo-api/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "6000" // Default port if not specified
	}
	router.Run("localhost:" + port)
}
