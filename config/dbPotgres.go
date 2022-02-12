package config

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabasePostgres() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	dsn := "host=127.0.0.1 user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASS") + " dbname=" + os.Getenv("DB") + " port=5432 sslmode=disable"
	//dsn := "postgres://" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@127.0.0.1:5432/" + os.Getenv("DB")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	//opt, err := pg.ParseURL(dsn)

	if err != nil {
		panic(err)
	}

	DB = db
}
