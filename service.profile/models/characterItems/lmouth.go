package profileCharacter

import (
	"avirtan/work_craft/config"
	profile "avirtan/work_craft/service.profile/models"
	"time"
)

type Mouth struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"unique; not null"`
	Path      string `gorm:"not null"`
	CreatedAt time.Time
	Character []profile.Character `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (mouth *Mouth) Create() error {
	result := config.DB.Save(mouth)
	return result.Error
}

func GetAllMouth() (*[]Mouth, error) {
	mouths := &[]Mouth{}
	result := config.DB.Find(&mouths)
	return mouths, result.Error
}

func (mouth *Mouth) Delete() error {
	result := config.DB.Delete(mouth, mouth.ID)
	return result.Error
}

func GetMouthById(id uint) (*Mouth, error) {
	mouth := &Mouth{}
	result := config.DB.First(mouth, id)
	return mouth, result.Error
}
