package profileCharacter

import (
	"avirtan/work_craft/config"
	profile "avirtan/work_craft/service.profile/models"
	"time"
)

type Top struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"unique; not null"`
	Path      string `gorm:"not null"`
	CreatedAt time.Time
	Character []profile.Character `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (top *Top) Create() error {
	result := config.DB.Save(top)
	return result.Error
}

func GetAllTop() (*[]Top, error) {
	top := &[]Top{}
	result := config.DB.Find(&top)
	return top, result.Error
}

func (top *Top) Delete() error {
	result := config.DB.Delete(top, top.ID)
	return result.Error
}

func GetTopById(id uint) (*Top, error) {
	top := &Top{}
	result := config.DB.First(top, id)
	return top, result.Error
}
