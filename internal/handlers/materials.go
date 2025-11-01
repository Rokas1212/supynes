package handlers

import (
	"fmt"
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

func AddDefaultMaterials(db *gorm.DB) {
	defaultMaterials := []models.Material{
		{Name: "Wood"},
		{Name: "Metal"},
		{Name: "Plastic"},
		{Name: "Composite"},
		{Name: "Carbon Fiber"},
		{Name: "Fiberglass"},
		{Name: "Aluminum"},
		{Name: "Steel"},
		{Name: "Titanium"},
		{Name: "Rubber"},
		{Name: "Foam"},
		{Name: "Leather"},
		{Name: "Nylon"},
		{Name: "Kevlar"},
		{Name: "Cork"},
	}

	for _, material := range defaultMaterials {
		var existingMaterial models.Material
		result := db.Where("name = ?", material.Name).First(&existingMaterial)
		if result.Error == gorm.ErrRecordNotFound {
			db.Create(&material)
		}
	}
	fmt.Println("Default materials added or already exist.")
}
