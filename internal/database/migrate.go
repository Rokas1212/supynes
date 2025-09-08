package database

import (
	"github.com/Rokas1212/supynes/internal/models"
)

func AutoMigrate() error {
	if err := DB.AutoMigrate(
		&models.User{},
		&models.Playground{},
		&models.Swing{},
		&models.Review{},
		&models.Photo{},
		&models.Tag{},
		&models.Material{},
		&models.SwingTag{},
		&models.SwingMaterial{},
		&models.Favorite{},
	); err != nil {
		return err
	}

	// Ensure join-tables are used for many2many with extra fields
	if err := DB.SetupJoinTable(&models.Swing{}, "Tags", &models.SwingTag{}); err != nil {
		return err
	}
	if err := DB.SetupJoinTable(&models.Swing{}, "Materials", &models.SwingMaterial{}); err != nil {
		return err
	}

	return nil
}
