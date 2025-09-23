package models

import "time"

type User struct {
	ID          uint   `gorm:"primaryKey"`
	Email       string `gorm:"uniqueIndex;not null"`
	Password    string `gorm:"not null"`
	DisplayName string `gorm:"not null"`
	IsAdmin     bool   `gorm:"not null;default:false"`
	IsModerator bool   `gorm:"not null;default:false"`
	CreatedAt   time.Time
}
