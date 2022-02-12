package profile

import (
	"time"
)

type Icon struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"not null"`
	Path         string `gorm:"not null"`
	CreatedAt    time.Time
	AchivementID *uint
	User         []*User
}
