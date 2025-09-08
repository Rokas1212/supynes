package main

import (
	"log"
	"os"

	"github.com/Rokas1212/supynes/internal/database"
	"github.com/gin-gonic/gin"
)

func main() {

	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}
	log.Print("Connected to DB")
	if err := database.AutoMigrate(); err != nil {
		log.Fatal(err)
	}
	log.Print("Migration Complete")
	// Read PORT from environment, default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create Gin router
	router := gin.Default()

	// Simple GET endpoint
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World! 🚀",
		})
	})

	router.GET("/migrate", func(c *gin.Context) {
		if err := database.AutoMigrate(); err != nil {
			log.Fatal(err)
		}
		log.Print("Migration Complete")
	})

	log.Printf("Server running on http://localhost:%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
