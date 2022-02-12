package profile

import (
	"avirtan/work_craft/config"

	"github.com/gofrs/uuid"
)

type Character struct {
	ID       uint `gorm:"primaryKey"`
	EyeID    uint
	HairID   uint
	MouthID  uint
	NoseID   uint
	OtherID  uint
	BottomID uint
	ShoeID   uint
	TopID    uint
	WeaponID *uint
	UserID   uuid.UUID
}

func (char *Character) Create() error {
	result := config.DB.Create(char)
	return result.Error
}

func GetAllCharacter() ([]Character, error) {
	characters := []Character{}
	result := config.DB.Find(&characters)
	return characters, result.Error
}
