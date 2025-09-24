package main

import (
	"log"
	"os"

	"github.com/Rokas1212/supynes/internal/database"
	"github.com/Rokas1212/supynes/internal/handlers"
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8085"
	}

	// Create Gin router
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	router.POST("/login", func(c *gin.Context) {
		handlers.Login(c, database.DB)
	})
	router.GET("/swings", func(c *gin.Context) {
		handlers.GetAllSwings(c, database.DB)
	})
	router.POST("/swing", func(c *gin.Context) {
		handlers.AddSwing(c, database.DB)
	})
	router.POST("/register", func(c *gin.Context) {
		handlers.Register(c, database.DB)
	})

	log.Printf("Server running on http://localhost:%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
