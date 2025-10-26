package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Rokas1212/supynes/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AddFavoriteRequest struct {
	SwingID uint `json:"swingID"`
}

type GetFavoritesRequest struct {
	UserID uint `json:"userID"`
}

type GetFavoritesResponse struct {
	Swings []uint `json:"swings"`
}

func AddFavorite(c *gin.Context, db *gorm.DB) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	uid, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid userID type"})
		return
	}

	var req AddFavoriteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	favorite := models.Favorite{
		UserID:    uid,
		SwingID:   req.SwingID,
		CreatedAt: time.Now(),
	}

	var existingFavorite models.Favorite
	result := db.Where("user_id = ? AND swing_id = ?", uid, req.SwingID).First(&existingFavorite)
	if result.Error == nil {
		db.Delete(&existingFavorite)
		c.JSON(http.StatusOK, gin.H{"msg": "Favorite removed"})
		return
	}

	resultCreate := db.Create(&favorite)
	if resultCreate.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add favorite"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"msg": "Successully added swing to favorites!"})
}

func GetFavorites(c *gin.Context, db *gorm.DB) {
	userIDStr := c.Query("userID")
	if userIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing userID"})
		return
	}
	var userID uint
	_, err := fmt.Sscanf(userIDStr, "%d", &userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid userID"})
		return
	}

	var favorites []models.Favorite
	result := db.Where("user_id = ?", userID).Find(&favorites)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get favorites"})
		return
	}

	swingIDs := make([]uint, len(favorites))
	for i, fav := range favorites {
		swingIDs[i] = fav.SwingID
	}

	c.JSON(http.StatusOK, GetFavoritesResponse{Swings: swingIDs})
}
