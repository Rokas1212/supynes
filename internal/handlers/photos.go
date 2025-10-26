package handlers

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"

	"github.com/Rokas1212/supynes/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func uploadPhoto(fileHeader *multipart.FileHeader) (string, error) {
	// Open the uploaded file
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Create destination file
	dstPath := fmt.Sprintf("/app/media/%d_%s", time.Now().UnixNano(), fileHeader.Filename)
	out, err := os.Create(dstPath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	// Copy the uploaded file to the destination
	if _, err := io.Copy(out, file); err != nil {
		return "", err
	}

	return dstPath, nil
}

func GetPhotosBySwingID(c *gin.Context, db *gorm.DB) {
	swingID := c.Param("id")

	var photos []string
	if err := db.Model(&models.Photo{}).Where("swing_id = ?", swingID).Pluck("url", &photos).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve photos"})
		return
	}

	c.JSON(200, gin.H{"photos": photos})
}
