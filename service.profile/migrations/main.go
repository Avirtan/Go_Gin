package main

import (
	"avirtan/work_craft/config"
	"flag"
	"log"

	"github.com/joho/godotenv"
)

// метод для запуск всех миграций созданных при помощи утилты migrations
func main() {
	action := flag.String("a", "", "a string")
	flag.Parse()
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	config.ConnectDatabasePostgres()
	//fmt.Println(*action)
	initProfile_migrate(*action)
}

func initProfile_migrate(action string) {
	m := &initProfile{}
	if action == "up" {
		m.up()
	} else {
		m.down()
		//e := os.Remove("initProfilemigration2022-01-04_17:13:46.go")
		// if e != nil {
		// 	log.Fatal(e)
		// }
	}
}
