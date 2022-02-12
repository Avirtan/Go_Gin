package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// небольшая утилита для создания файлов миграции
// использование ./migration -d=название_миграции
func main() {
	description := flag.String("d", "", "a string")
	flag.Parse()
	if *description == "" {
		panic("need args -d")
	}
	createMigration(*description)

}

func createMigration(description string) {
	currentTime := time.Now()
	fileName := description + "migration" + currentTime.Format("2006-01-02_15:04:05") + ".go"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	text := `
	package main

	import (
		"avirtan/work_craft/config"
	)

	type %d struct{}

	func (m *%d)up() {
		config.ConnectDatabasePostgres()
		//config.DB.Debug().Migrator()
	}

	func (m *%d)down(){
		config.ConnectDatabasePostgres()
		//config.DB.Debug().Migrator()
	}
	`
	text = strings.ReplaceAll(text, "%d", description)
	file.WriteString(text)
	defer file.Close()
	//
	f, err := os.OpenFile("main.go", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	text = `
		func %d_migrate(action string){
			m:=&%d{}
			if action == "up"{
			m.up()
			}else {
				m.down()
				//e := os.Remove("%f")
				// if e != nil {
				// 	log.Fatal(e)
				// }
			}
		}
	`
	text = strings.ReplaceAll(text, "%d", description)
	text = strings.ReplaceAll(text, "%f", fileName)
	if _, err = f.WriteString(text); err != nil {
		panic(err)
	}
}
