package handlers

import (
	"net/http"

	"github.com/Rokas1212/supynes/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllMaterials(c *gin.Context, db *gorm.DB) {
	var materials []models.Material
	result := db.Find(&materials)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch materials"})
		return
	}
	c.JSON(http.StatusOK, materials)
}
