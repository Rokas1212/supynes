package handlers

import (
	"net/http"
	"time"

	"github.com/Rokas1212/supynes/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PlaygroundAddRequest struct {
	UserID  uint
	Name    string
	Address string
	Lat     *float64
	Lng     *float64
	City    string
}

func AddPlayground(c *gin.Context, db *gorm.DB) {
	var req PlaygroundAddRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	playground := models.Playground{
		UserID:    req.UserID,
		Name:      req.Name,
		Address:   req.Address,
		Lat:       req.Lat,
		Lng:       req.Lng,
		City:      req.City,
		CreatedAt: time.Now(),
	}

	result := db.Create(&playground)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create playground"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Playground created successfully",
		"id":      playground.ID,
	})
}

func GetAllPlaygrounds(c *gin.Context, db *gorm.DB) {
	var playgrounds []models.Playground
	result := db.Find(&playgrounds)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch playgrounds"})
		return
	}
	c.JSON(http.StatusOK, playgrounds)
}
