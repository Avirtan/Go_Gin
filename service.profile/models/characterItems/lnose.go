package profileCharacter

import (
	"avirtan/work_craft/config"
	profile "avirtan/work_craft/service.profile/models"
	"time"
)

type Nose struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"unique; not null"`
	Path      string `gorm:"not null"`
	CreatedAt time.Time
	Character []profile.Character `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (nose *Nose) Create() error {
	result := config.DB.Save(nose)
	return result.Error
}

func GetAllNose() (*[]Nose, error) {
	noses := &[]Nose{}
	result := config.DB.Find(&noses)
	return noses, result.Error
}

func (nose *Nose) Delete() error {
	result := config.DB.Delete(nose, nose.ID)
	return result.Error
}

func GetNoseById(id uint) (*Nose, error) {
	nose := &Nose{}
	result := config.DB.First(nose, id)
	return nose, result.Error
}
