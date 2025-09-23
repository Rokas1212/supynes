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
		Email       string `json:"email"`
		DisplayName string `json:"displayName"`
	} `json:"user"`
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

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// TODO: ADD VALIDATION (anything is allowed in for now)

	resp := LoginResponse{
		Token: "dummy-wummy",
		User: struct {
			Email       string `json:"email"`
			DisplayName string `json:"displayName"`
		}{
			Email:       req.Email,
			DisplayName: "User",
		},
	}
	c.JSON(http.StatusOK, resp)
}
