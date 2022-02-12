package profileCharacter

import (
	"avirtan/work_craft/config"
	profile "avirtan/work_craft/service.profile/models"
	"time"
)

type Shoe struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"unique; not null"`
	Path      string `gorm:"not null"`
	CreatedAt time.Time
	Character []profile.Character `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (shoe *Shoe) Create() error {
	result := config.DB.Save(shoe)
	return result.Error
}

func GetAllShoes() (*[]Shoe, error) {
	shoes := &[]Shoe{}
	result := config.DB.Find(&shoes)
	return shoes, result.Error
}

func (shoe *Shoe) Delete() error {
	result := config.DB.Delete(shoe, shoe.ID)
	return result.Error
}

func GetShoeById(id uint) (*Shoe, error) {
	shoe := &Shoe{}
	result := config.DB.First(shoe, id)
	return shoe, result.Error
}
