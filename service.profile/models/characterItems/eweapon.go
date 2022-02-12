package profileCharacter

import (
	"avirtan/work_craft/config"
	profile "avirtan/work_craft/service.profile/models"
	"time"
)

type Weapon struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"unique; not null"`
	Path      string `gorm:"not null"`
	CreatedAt time.Time
	Character []profile.Character `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (weapon *Weapon) Create() error {
	result := config.DB.Save(weapon)
	return result.Error
}

func GetAllWeapon() (*[]Weapon, error) {
	weapon := &[]Weapon{}
	result := config.DB.Find(&weapon)
	return weapon, result.Error
}

func (weapon *Weapon) Delete() error {
	result := config.DB.Delete(weapon, weapon.ID)
	return result.Error
}

func GetWeaponById(id *uint) (*Weapon, error) {
	if id == nil {
		return nil, nil
	}
	weapon := &Weapon{}
	result := config.DB.First(weapon, id)
	return weapon, result.Error
}
