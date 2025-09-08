package models

import "time"

type Review struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"not null;index;uniqueIndex:uniq_user_swing"`
	SwingID   uint `gorm:"not null;index;uniqueIndex:uniq_user_swing"`
	Rating    int  `gorm:"not null;check:rating >= 1 AND rating <= 5"`
	Title     string
	Body      string
	VisitedOn *time.Time
	CreatedAt time.Time

	User  User  `gorm:"constraint:OnDelete:CASCADE"`
	Swing Swing `gorm:"constraint:OnDelete:CASCADE"`
}
