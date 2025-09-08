package models

import "time"

type Photo struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null;index"`
	SwingID   *uint  `gorm:"index"`
	ReviewID  *uint  `gorm:"index"`
	URL       string `gorm:"not null"`
	Caption   string
	CreatedAt time.Time

	User   User    `gorm:"constraint:OnDelete:SET NULL"`
	Swing  *Swing  `gorm:"constraint:OnDelete:CASCADE"`
	Review *Review `gorm:"constraint:OnDelete:CASCADE"`
}
