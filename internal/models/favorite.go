package models

import "time"

type Favorite struct {
	UserID    uint `gorm:"primaryKey;autoIncrement:false"`
	SwingID   uint `gorm:"primaryKey;autoIncrement:false"`
	CreatedAt time.Time

	User  User  `gorm:"constraint:OnDelete:CASCADE"`
	Swing Swing `gorm:"constraint:OnDelete:CASCADE"`
}
