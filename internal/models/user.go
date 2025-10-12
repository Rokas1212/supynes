package models

import "time"

type User struct {
	ID          uint   `gorm:"primaryKey"`
	Email       string `gorm:"uniqueIndex;not null"`
	Password    string `gorm:"not null"`
	DisplayName string `gorm:"not null"`
	Role        int    `gorm:"not null;default:1"`
	CreatedAt   time.Time
}

const (
	RoleUser      = 1
	RoleModerator = 2
	RoleAdmin     = 3
)
