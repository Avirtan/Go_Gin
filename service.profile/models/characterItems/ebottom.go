package profileCharacter

import (
	"avirtan/work_craft/config"
	profile "avirtan/work_craft/service.profile/models"
	"time"
)

type Bottom struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"unique; not null"`
	Path      string `gorm:"not null"`
	CreatedAt time.Time
	Character []profile.Character `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (bottom *Bottom) Create() error {
	result := config.DB.Save(bottom)
	return result.Error
}

func GetAllBottom() (*[]Bottom, error) {
	bottom := &[]Bottom{}
	result := config.DB.Find(&bottom)
	return bottom, result.Error
}

func (bottom *Bottom) Delete() error {
	result := config.DB.Delete(bottom, bottom.ID)
	return result.Error
}

func GetBottomById(id uint) (*Bottom, error) {
	bottom := &Bottom{}
	result := config.DB.First(bottom, id)
	return bottom, result.Error
}
