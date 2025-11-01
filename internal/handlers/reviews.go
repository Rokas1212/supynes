package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Rokas1212/supynes/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AddReviewRequest struct {
	UserID    uint       `json:"userID" binding:"required"`
	SwingID   uint       `json:"swingID" binding:"required"`
	Rating    int        `json:"rating" binding:"required,min=1,max=5"`
	Title     string     `json:"title" binding:"required"`
	Body      string     `json:"body" binding:"required"`
	VisitedOn *time.Time `json:"visitedOn"`
}

func GetSwingReviews(c *gin.Context, db *gorm.DB) {
	swingID := c.Param("id")

	var reviews []models.Review
	if err := db.Where("swing_id = ?", swingID).Find(&reviews).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve reviews"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reviews": reviews})
}

func RemoveReview(c *gin.Context, db *gorm.DB) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	reviewID := c.Param("id")

	var review models.Review
	if err := db.First(&review, reviewID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
		return
	}

	isAdmin, err := IsAdmin(userID.(uint), db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify user role"})
		return
	}

	if review.UserID != userID && !isAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "Cannot delete another user's review"})
		return
	}

	if err := db.Delete(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete review"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Review deleted successfully"})
}

func AddReview(c *gin.Context, db *gorm.DB) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var req AddReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid request %v", err)})
		return
	}

	if userID != req.UserID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Cannot add review for another user"})
		return
	}

	var existingReview models.Review
	result := db.Where("user_id = ? AND swing_id = ?", req.UserID, req.SwingID).First(&existingReview)
	if result.Error == nil {
		updateReview := models.Review{
			Rating:    req.Rating,
			Title:     req.Title,
			Body:      req.Body,
			VisitedOn: req.VisitedOn,
		}
		if err := db.Model(&existingReview).Updates(updateReview).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update review"})
			return
		}
		c.JSON(http.StatusOK, existingReview)
		return
	}

	review := models.Review{
		UserID:    req.UserID,
		SwingID:   req.SwingID,
		Rating:    req.Rating,
		Title:     req.Title,
		Body:      req.Body,
		VisitedOn: req.VisitedOn,
		CreatedAt: time.Now(),
	}

	if err := db.Create(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add review"})
		return
	}

	c.JSON(http.StatusCreated, review)
}
