package handlers

import (
	"fmt"

	"github.com/Rokas1212/supynes/internal/models"
	"gorm.io/gorm"
)

func IsAdmin(userID uint, db *gorm.DB) (bool, error) {
	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		return false, err
	}
	fmt.Println("User role:", user.Role)
	fmt.Println("Is admin:", user.Role == models.RoleAdmin)
	return user.Role == models.RoleAdmin, nil
}
