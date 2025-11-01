package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Rokas1212/supynes/internal/auth"
	"github.com/Rokas1212/supynes/internal/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	DisplayName string `json:"displayName"`
}

type ProfileResponse struct {
	UserID      uint           `json:"userID"`
	Email       string         `json:"email"`
	DisplayName string         `json:"displayName"`
	Role        int            `json:"role"`
	CreatedAt   time.Time      `json:"createdAt"`
	Swings      []models.Swing `json:"swings,omitempty"`
}

func GetProfile(c *gin.Context, db *gorm.DB) {
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

	var dbUser models.User
	result := db.First(&dbUser, "id = ?", uid)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid input when getting user profile"})
		return
	}

	var swings []models.Swing
	db.Where("user_id = ?", uid).Find(&swings)

	response := ProfileResponse{
		UserID:      dbUser.ID,
		Email:       dbUser.Email,
		DisplayName: dbUser.DisplayName,
		Role:        dbUser.Role,
		CreatedAt:   dbUser.CreatedAt,
		Swings:      swings,
	}

	c.JSON(http.StatusOK, response)
}

func Register(c *gin.Context, db *gorm.DB) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	encryptedPassword, err := hashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encrypt password"})
		return
	}

	user := models.User{
		Email:       req.Email,
		DisplayName: req.DisplayName,
		Password:    encryptedPassword,
		CreatedAt:   time.Now(),
		Role:        models.RoleUser,
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

	if !verifyHash(reqUser.Password, dbUser.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong email or password!"})
		return
	}

	token, err := auth.GenerateToken(dbUser.ID, dbUser.Email, dbUser.DisplayName, dbUser.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate jwt token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func DeleteUser(c *gin.Context, db *gorm.DB) {
	userIDFromToken, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var dbUser models.User
	result := db.First(&dbUser, "id = ?", userIDFromToken)
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

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func verifyHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
