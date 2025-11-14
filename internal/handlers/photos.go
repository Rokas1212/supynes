package handlers

import (
	"fmt"
	"io/fs"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/Rokas1212/supynes/internal/models"
	"github.com/gin-gonic/gin"
	uuid "github.com/google/uuid"
	"gorm.io/gorm"
)

func uploadPhoto(c *gin.Context, fileHeader *multipart.FileHeader) (string, error) {
	id := uuid.New().String()
	filename := fmt.Sprintf("%s.png", id)
	dstPath := filepath.Join("/app/media", filename)

	if err := c.SaveUploadedFile(fileHeader, dstPath, fs.FileMode(0o640)); err != nil {
		return "", err
	}

	return "/media/" + filename, nil
}

func GetPhotosBySwingID(c *gin.Context, db *gorm.DB) {
	swingID := c.Param("id")

	var photos []string
	if err := db.Model(&models.Photo{}).Where("swing_id = ?", swingID).Pluck("url", &photos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve photos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"photos": photos})
}
