package handlers

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"
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
	MaterialIDs  []uint
	Photos       []*multipart.FileHeader `form:"photos[]"`
}

func AddSwing(c *gin.Context, db *gorm.DB) {
	// Parse multipart form
	if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid multipart form"})
		return
	}

	// Get form values
	userID, _ := strconv.ParseUint(c.PostForm("UserID"), 10, 64)
	name := c.PostForm("Name")
	address := c.PostForm("Address")
	lat, _ := strconv.ParseFloat(c.PostForm("Lat"), 64)
	lng, _ := strconv.ParseFloat(c.PostForm("Lng"), 64)
	city := c.PostForm("City")
	seatCount, _ := strconv.Atoi(c.PostForm("SeatCount"))
	maxWeightKg, _ := strconv.Atoi(c.PostForm("MaxWeightKg"))
	isAccessible := c.PostForm("IsAccessible") == "true"

	// Create swing
	swing := models.Swing{
		UserID:       uint(userID),
		Name:         name,
		Address:      address,
		Lat:          &lat,
		Lng:          &lng,
		City:         city,
		SeatCount:    &seatCount,
		MaxWeightKg:  &maxWeightKg,
		IsAccessible: isAccessible,
		CreatedAt:    time.Now(),
	}

	result := db.Create(&swing)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create swing"})
		return
	}

	// Handle photo uploads
	form, err := c.MultipartForm()
	if err == nil && form.File != nil {
		files := form.File["photos[]"]
		for _, fileHeader := range files {
			uploadedPath, err := uploadPhoto(fileHeader)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to upload photo: %s", err)})
				return
			}
			swingPhoto := models.Photo{
				URL:       uploadedPath,
				CreatedAt: time.Now(),
				UserID:    uint(userID),
				SwingID:   &swing.ID,
				Caption:   "Photo for swing " + name,
			}
			db.Create(&swingPhoto)
		}
	}

	// Handle tags
	tags := c.PostFormArray("Tags[]")
	var failedTags []string
	for _, tag := range tags {
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

	// Handle materials
	materialIDs := c.PostFormArray("MaterialIDs[]")
	var failedMaterialIDs []uint
	for _, midStr := range materialIDs {
		mid, err := strconv.ParseUint(midStr, 10, 64)
		if err != nil {
			continue
		}
		var material models.Material
		if err := db.First(&material, mid).Error; err != nil {
			failedMaterialIDs = append(failedMaterialIDs, uint(mid))
			continue
		}
		swingMaterial := models.SwingMaterial{
			SwingID:    swing.ID,
			MaterialID: uint(mid),
			CreatedAt:  time.Now(),
		}
		if err := db.Create(&swingMaterial).Error; err != nil {
			failedMaterialIDs = append(failedMaterialIDs, uint(mid))
		}
	}

	// Build response
	response := gin.H{
		"message": "Swing created successfully",
		"id":      swing.ID,
	}
	if len(failedTags) > 0 {
		response["warning_tags"] = fmt.Sprintf("Some tags couldn't be associated: %v", failedTags)
	}
	if len(failedMaterialIDs) > 0 {
		response["warning_materials"] = fmt.Sprintf("Some materials couldn't be associated: %v", failedMaterialIDs)
	}

	c.JSON(http.StatusCreated, response)
}

func GetAllSwings(c *gin.Context, db *gorm.DB) {
	var swings []models.Swing
	result := db.Preload("Tags").Find(&swings)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch swings"})
		return
	}
	c.JSON(http.StatusOK, swings)
}

func GetSwingByID(c *gin.Context, db *gorm.DB) {
	swingID := c.Param("id")
	var swing models.Swing
	result := db.Preload("Tags").Preload("Materials").Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, display_name")
	}).First(&swing, swingID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Swing not found"})
		return
	}
	c.JSON(http.StatusOK, swing)
}

func DeleteSwing(c *gin.Context, db *gorm.DB) {
	swingID := c.Param("id")
	userId, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var swing models.Swing
	result := db.First(&swing, swingID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Swing not found"})
		return
	}

	var user models.User
	result = db.First(&user, userId)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}

	if swing.UserID != userId && user.Role != 2 {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to delete this swing"})
		return
	}

	result = db.Delete(&models.Swing{}, swingID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete swing"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Swing deleted successfully"})
}
