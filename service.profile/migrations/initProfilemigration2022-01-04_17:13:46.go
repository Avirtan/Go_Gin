package main

import (
	"avirtan/work_craft/config"
	profile "avirtan/work_craft/service.profile/models"
	profileCharacter "avirtan/work_craft/service.profile/models/characterItems"
)

type initProfile struct{}

func (m *initProfile) up() {
	config.ConnectDatabasePostgres()
	config.DB.Debug().AutoMigrate(&profile.Role{})
	config.DB.Debug().AutoMigrate(&profile.Achivement{})
	config.DB.Debug().AutoMigrate(&profile.Statistic{})
	config.DB.Debug().AutoMigrate(&profile.Icon{})
	config.DB.Debug().AutoMigrate(&profileCharacter.Eye{})
	config.DB.Debug().AutoMigrate(&profileCharacter.Hair{})
	config.DB.Debug().AutoMigrate(&profileCharacter.Bottom{})
	config.DB.Debug().AutoMigrate(&profileCharacter.Mouth{})
	config.DB.Debug().AutoMigrate(&profileCharacter.Nose{})
	config.DB.Debug().AutoMigrate(&profileCharacter.Other{})
	config.DB.Debug().AutoMigrate(&profileCharacter.Shoe{})
	config.DB.Debug().AutoMigrate(&profileCharacter.Top{})
	config.DB.Debug().AutoMigrate(&profileCharacter.Weapon{})
	config.DB.Debug().AutoMigrate(&profile.Character{})
	config.DB.Debug().AutoMigrate(&profile.User{})
}

func (m *initProfile) down() {
	config.ConnectDatabasePostgres()
	config.DB.Debug().Migrator().DropTable(&profile.Role{})
	config.DB.Debug().Migrator().DropTable(&profile.Achivement{})
	config.DB.Debug().Migrator().DropTable(&profile.Statistic{})
	config.DB.Debug().Migrator().DropTable(&profile.Icon{})
	config.DB.Debug().Migrator().DropTable(&profileCharacter.Eye{})
	config.DB.Debug().Migrator().DropTable(&profileCharacter.Hair{})
	config.DB.Debug().Migrator().DropTable(&profileCharacter.Bottom{})
	config.DB.Debug().Migrator().DropTable(&profileCharacter.Mouth{})
	config.DB.Debug().Migrator().DropTable(&profileCharacter.Nose{})
	config.DB.Debug().Migrator().DropTable(&profileCharacter.Other{})
	config.DB.Debug().Migrator().DropTable(&profileCharacter.Shoe{})
	config.DB.Debug().Migrator().DropTable(&profileCharacter.Top{})
	config.DB.Debug().Migrator().DropTable(&profileCharacter.Weapon{})
	config.DB.Debug().Migrator().DropTable(&profile.Character{})
	config.DB.Debug().Migrator().DropTable(&profile.User{})
}
