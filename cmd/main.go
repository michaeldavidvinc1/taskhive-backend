package main

import (
	"backend/config"
	"backend/routes"
	"log"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Initialize database
	config.ConnectDB()

	// Setup router
	router := routes.SetupRouter()

	// Start server
	if err := router.Run(":" + config.GetEnv("PORT", "8000")); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
