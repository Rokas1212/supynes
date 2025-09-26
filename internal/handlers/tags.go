package handlers

import (
	"github.com/Rokas1212/supynes/internal/models"
	"gorm.io/gorm"
)

type AddTagRequest struct {
	Name string `json:"name" binding:"required"`
}

func findTag(tag string, db *gorm.DB) (*uint, error) {
	var dbTag models.Tag
	result := db.First(&dbTag, "name = ?", tag)

	err := result.Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &dbTag.ID, nil
}

func addTag(tag string, db *gorm.DB) error {
	// TODO: Do smth with descriptions
	id, err := findTag(tag, db)

	if err != nil {
		return err
	}

	if id != nil {
		return nil
	}

	tagModel := models.Tag{Name: tag, Description: ""}

	result := db.Create(&tagModel)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
