package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Rokas1212/supynes/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SwingAddRequest struct {
	UserID       uint
	Name         string
	SeatCount    *int
	MaxWeightKg  *int
	IsAccessible bool
	Address      string
	Lat          *float64
	Lng          *float64
	City         string
	CreatedAt    time.Time
	Tags         []string
}

func AddSwing(c *gin.Context, db *gorm.DB) {
	var req SwingAddRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %s", err)})
		return
	}

	swing := models.Swing{
		UserID:       req.UserID,
		Name:         req.Name,
		Address:      req.Address,
		Lat:          req.Lat,
		Lng:          req.Lng,
		City:         req.City,
		SeatCount:    req.SeatCount,
		MaxWeightKg:  req.MaxWeightKg,
		IsAccessible: req.IsAccessible,
		CreatedAt:    time.Now(),
	}

	result := db.Create(&swing)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create swing"})
		return
	}

	if len(req.Tags) > 0 {
		var failedTags []string

		for _, tag := range req.Tags {

			tagErr := addTag(tag, db)
			if tagErr != nil {
				failedTags = append(failedTags, tag)
				continue
			}

			tagID, err := findTag(tag, db)
			if err != nil || tagID == nil {
				failedTags = append(failedTags, tag)
				continue
			}

			swingTag := models.SwingTag{
				SwingID:   swing.ID,
				TagID:     *tagID,
				CreatedAt: time.Now(),
			}

			if err := db.Create(&swingTag).Error; err != nil {
				failedTags = append(failedTags, tag)
			}
		}

		// Return success with optional warning about failed tags
		response := gin.H{
			"message": "Swing created successfully",
			"id":      swing.ID,
		}

		if len(failedTags) > 0 {
			response["warning"] = fmt.Sprintf("Some tags couldn't be associated: %v", failedTags)
		}

		c.JSON(http.StatusCreated, response)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Swing created successfully",
		"id":      swing.ID,
	})
}

func GetAllSwings(c *gin.Context, db *gorm.DB) {
	var swings []models.Swing
	result := db.Find(&swings)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch swings"})
		return
	}
	c.JSON(http.StatusOK, swings)
}
