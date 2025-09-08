package models

import "time"

type SwingMaterial struct {
	SwingID    uint `gorm:"primaryKey;autoincrement:false"`
	MaterialID uint `gorm:"primaryKey;autoincrement:false"`
	CreatedAt  time.Time

	Swing    Swing    `gorm:"constraint:OnDelete:CASCADE"`
	Material Material `gorm:"constraint:OnDelete:CASCADE"`
}

func (SwingMaterial) TableName() string { return "swing_materials" }
