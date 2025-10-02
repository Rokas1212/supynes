package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Rokas1212/supynes/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  struct {
		UserID      uint   `json:"userID"`
		Email       string `json:"email"`
		DisplayName string `json:"displayName"`
	} `json:"user"`
}

type DeleteRequest struct {
	UserID uint `json:"userID"`
}

type RegisterRequest struct {
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	DisplayName string `json:"displayName"`
}

func Register(c *gin.Context, db *gorm.DB) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user := models.User{
		Email:       req.Email,
		DisplayName: req.DisplayName,
		Password:    req.Password, // TODO: ENCRYPT
		CreatedAt:   time.Now(),
		IsAdmin:     false,
		IsModerator: false,
	}

	result := db.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":     fmt.Sprintf("User with id=%d, succesfully created", user.ID),
		"displayName": user.DisplayName,
	})
}

func Login(c *gin.Context, db *gorm.DB) {
	var reqUser LoginRequest
	if err := c.ShouldBindJSON(&reqUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var dbUser models.User
	result := db.First(&dbUser, "email = ?", reqUser.Email)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if dbUser.Password != reqUser.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong email or password!"})
		return
	}
	// TODO: Implement JWT

	resp := LoginResponse{
		Token: "dummy-wummy",
		User: struct {
			UserID      uint   `json:"userID"`
			Email       string `json:"email"`
			DisplayName string `json:"displayName"`
		}{
			UserID:      dbUser.ID,
			Email:       dbUser.Email,
			DisplayName: dbUser.DisplayName,
		},
	}
	c.JSON(http.StatusOK, resp)
}

func DeleteUser(c *gin.Context, db *gorm.DB) {
	var reqDelUser DeleteRequest
	if err := c.ShouldBindJSON(&reqDelUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// TODO: ADD Validation for who is deleting the user
	var dbUser models.User
	result := db.First(&dbUser, "id = ?", reqDelUser.UserID)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid input when deleting user"})
		return
	}

	result = db.Delete(&dbUser)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"message":     fmt.Sprintf("User with id=%d, succesfully delete", dbUser.ID),
		"displayName": dbUser.DisplayName,
	})
}
