package main

import (
	"github.com/ar-tiwari75/wealth-manager/routes"
	"github.com/ar-tiwari75/wealth-manager/storage"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	// Initialize the database connection
	storage.Connect()

	// Create a new Gin router
	r := gin.Default()

	//Register all routes
	routes.RegisterRoutes(r)

	// Start the server
	r.Run(":8080")
}
