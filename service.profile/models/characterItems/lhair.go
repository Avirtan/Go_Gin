package profileCharacter

import (
	"avirtan/work_craft/config"
	profile "avirtan/work_craft/service.profile/models"
	"time"
)

type Hair struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"unique; not null"`
	Path      string `gorm:"not null"`
	CreatedAt time.Time
	Character []profile.Character `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (hair *Hair) Create() error {
	result := config.DB.Save(hair)
	return result.Error
}

func GetAllHair() ([]Hair, error) {
	hairs := []Hair{}
	result := config.DB.Find(&hairs)
	return hairs, result.Error
}

func (hair *Hair) Delete() error {
	result := config.DB.Delete(hair, hair.ID)
	return result.Error
}

func GetHairById(id uint) (*Hair, error) {
	hair := &Hair{}
	result := config.DB.First(hair, id)
	return hair, result.Error
}
