package models

import "time"

type Playground struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null;"`
	Name      string `gorm:"not null"`
	Address   string
	Lat       *float64
	Lng       *float64
	City      string
	CreatedAt time.Time

	User User `gorm:"constraint:OnDelete:CASCADE"`
}
