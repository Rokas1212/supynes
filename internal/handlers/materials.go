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

type AddMaterialRequest struct {
	Name string `json:"name"`
}

func AddNewMaterial(c *gin.Context, db *gorm.DB) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	isAdmin, err := IsAdmin(userID.(uint), db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify user role"})
		return
	}

	if !isAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "Cannot delete another user's review"})
		return
	}

	var req AddMaterialRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var existingMaterial models.Material
	result := db.Where("name = ?", req.Name).First(&existingMaterial)
	if result.Error == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Material already exists"})
		return
	}

	newMaterial := models.Material{Name: req.Name}
	if err := db.Create(&newMaterial).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add material"})
		return
	}

	c.JSON(http.StatusCreated, newMaterial)
}

func UpdateMaterial(c *gin.Context, db *gorm.DB) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	isAdmin, err := IsAdmin(userID.(uint), db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify user role"})
		return
	}

	if !isAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "Cannot update material without admin rights"})
		return
	}

	materialID := c.Param("id")

	var req AddMaterialRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var material models.Material
	if err := db.First(&material, materialID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Material not found"})
		return
	}

	material.Name = req.Name
	if err := db.Save(&material).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update material"})
		return
	}

	c.JSON(http.StatusOK, material)
}

func RemoveMaterial(c *gin.Context, db *gorm.DB) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	isAdmin, err := IsAdmin(userID.(uint), db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify user role"})
		return
	}

	if !isAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "Cannot delete another user's review"})
		return
	}

	materialID := c.Param("id")

	var material models.Material
	if err := db.First(&material, materialID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Material not found"})
		return
	}

	if err := db.Delete(&material).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete material"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Material deleted successfully"})
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
