package handlers

import (
	"github.com/Rokas1212/supynes/internal/models"
	"gorm.io/gorm"
)

const (
	RoleUser  = 1
	RoleAdmin = 2
)

func IsAdmin(userID uint, db *gorm.DB) (bool, error) {
	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		return false, err
	}
	return user.Role == RoleAdmin, nil
}
