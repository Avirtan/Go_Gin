package profileCharacter

import (
	"avirtan/work_craft/config"
	profile "avirtan/work_craft/service.profile/models"
	"time"
)

type Eye struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"unique; not null"`
	Path      string `gorm:"not null"`
	CreatedAt time.Time
	Character []profile.Character `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (eye *Eye) Create() error {
	result := config.DB.Save(eye)
	return result.Error
}

func GetAllEyes() (*[]Eye, error) {
	eyes := &[]Eye{}
	result := config.DB.Find(&eyes)
	return eyes, result.Error
}

func (eye *Eye) Delete() error {
	result := config.DB.Delete(eye, eye.ID)
	return result.Error
}

func GetEyesById(id uint) (*Eye, error) {
	eye := &Eye{}
	result := config.DB.First(eye, id)
	return eye, result.Error
}
