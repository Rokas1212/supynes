package models

import "time"

type SwingTag struct {
	SwingID   uint `gorm:"primaryKey;autoincrement:false"`
	TagID     uint `gorm:"primaryKey;autoincrement:false"`
	CreatedAt time.Time

	Swing Swing `gorm:"constraint:OnDelete:CASCADE"`
	Tag   Tag   `gorm:"constraint:OnDelete:CASCADE"`
}

func (SwingTag) TableName() string { return "swing_tags" }
