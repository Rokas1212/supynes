package models

import "time"

type Swing struct {
	ID           uint   `gorm:"primaryKey"`
	UserID       uint   `gorm:"not null;"`
	Name         string `gorm:"not null"`
	Description  string
	SeatCount    *int
	MaxWeightKg  *int
	IsAccessible bool `gorm:"not null;default:false"`
	Address      string
	Lat          *float64
	Lng          *float64
	City         string
	CreatedAt    time.Time

	Tags      []Tag      `gorm:"many2many:swing_tags;"`
	Materials []Material `gorm:"many2many:swing_materials;"`
	User      User       `gorm:"constraint:OnDelete:CASCADE"`
}
