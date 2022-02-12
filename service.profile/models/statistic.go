package profile

import (
	"avirtan/work_craft/config"

	"github.com/gofrs/uuid"
)

type Statistic struct {
	ID          uint         `gorm:"primary_key"`
	Earned      uint32       `gorm:"default:0"`
	Spent       uint32       `gorm:"default:0"`
	DoneQuest   uint16       `gorm:"default:0"`
	SentQuest   uint16       `gorm:"default:0"`
	Grade       uint8        `gorm:"default:0"`
	UserID      uuid.UUID    `gorm:"type:uuid"`
	Achivements []Achivement `gorm:"many2many:statistic_achivement"`
}

func (state *Statistic) Create() error {
	result := config.DB.Create(state)
	return result.Error
}
