package main

import (
	"log"
	"os"
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/Rokas1212/supynes/internal/database"
	"github.com/Rokas1212/supynes/internal/handlers"
	"github.com/Rokas1212/supynes/internal/middleware"
	"github.com/gin-gonic/gin"
)

func keyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func errorHandler(c *gin.Context, info ratelimit.Info) {
	c.String(429, "Too many requests. Try again in "+time.Until(info.ResetTime).String())
}

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

	handlers.AddDefaultMaterials(database.DB)

	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  5 * time.Second,
		Limit: 2,
	})
	mw := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc:      keyFunc,
	})

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
	router.GET("/favorites", func(c *gin.Context) {
		handlers.GetFavorites(c, database.DB)
	})
	router.GET("/materials", func(c *gin.Context) {
		handlers.GetAllMaterials(c, database.DB)
	})
	router.GET("/swings/:id", func(c *gin.Context) {
		handlers.GetSwingByID(c, database.DB)
	})
	router.GET("/photos/:id", func(c *gin.Context) {
		handlers.GetPhotosBySwingID(c, database.DB)
	})
	router.GET("/reviews/:id", func(c *gin.Context) {
		handlers.GetSwingReviews(c, database.DB)
	})
	router.GET("/average-ratings/:id", func(c *gin.Context) {
		handlers.GetSwingAverageRating(c, database.DB)
	})
	router.GET("/swing-name/:id", func(c *gin.Context) {
		handlers.GetSwingNameByID(c, database.DB)
	})

	protected := router.Group("/auth")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/favorite", func(c *gin.Context) {
			handlers.AddFavorite(c, database.DB)
		})
		protected.GET("/profile", mw, func(c *gin.Context) {
			handlers.GetProfile(c, database.DB)
		})
		protected.DELETE("/users/delete", func(c *gin.Context) {
			handlers.DeleteUser(c, database.DB)
		})
		protected.POST("/reviews/add", func(c *gin.Context) {
			handlers.AddReview(c, database.DB)
		})
		protected.DELETE("/swings/:id", func(c *gin.Context) {
			handlers.DeleteSwing(c, database.DB)
		})
		protected.DELETE("/reviews/:id", func(c *gin.Context) {
			handlers.RemoveReview(c, database.DB)
		})
		protected.PUT("/swings/:id", func(c *gin.Context) {
			handlers.UpdateSwing(c, database.DB)
		})
		protected.POST("/materials/add", func(c *gin.Context) {
			handlers.AddNewMaterial(c, database.DB)
		})
		protected.DELETE("/materials/remove/:id", func(c *gin.Context) {
			handlers.RemoveMaterial(c, database.DB)
		})
		protected.PUT("/materials/update/:id", func(c *gin.Context) {
			handlers.UpdateMaterial(c, database.DB)
		})
	}

	log.Printf("Server running on http://localhost:%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
