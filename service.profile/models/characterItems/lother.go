package profileCharacter

import (
	"avirtan/work_craft/config"
	profile "avirtan/work_craft/service.profile/models"
	"time"
)

type Other struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"unique; not null"`
	Path      string `gorm:"not null"`
	CreatedAt time.Time
	Character []profile.Character `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (other *Other) Create() error {
	result := config.DB.Save(other)
	return result.Error
}

func GetAllOther() (*[]Other, error) {
	others := &[]Other{}
	result := config.DB.Find(&others)
	return others, result.Error
}

func (other *Other) Delete() error {
	result := config.DB.Delete(other, other.ID)
	return result.Error
}

func GetOtherById(id uint) (*Other, error) {
	other := &Other{}
	result := config.DB.First(other, id)
	return other, result.Error
}
