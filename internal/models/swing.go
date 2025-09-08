package models

import "time"

type Swing struct {
	ID           uint   `gorm:"primaryKey"`
	PlaygroundID uint   `gorm:"not null;index"`
	Name         string `gorm:"not null"`
	SeatCount    *int
	MaxWeightKg  *int
	IsAccessible bool `gorm:"not null;default:false"`
	CreatedAt    time.Time

	Playground Playground `gorm:"constraint:OnDelete:CASCADE"`

	Tags      []Tag      `gorm:"many2many:swing_tags;"`
	Materials []Material `gorm:"many2many:swing_materials;"`
}
