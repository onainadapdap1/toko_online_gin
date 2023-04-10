package main

import (
	"os"
	"toko_online_gin/driver"
	"toko_online_gin/models"
	"toko_online_gin/seeds"
	"toko_online_gin/server"

	"github.com/jinzhu/gorm"
)

func drop(db *gorm.DB) {
	db.DropTableIfExists(
		&models.User{},
		&models.Category{},
	)
}

func create(database *gorm.DB) {
	drop(database)
	models.InitTable(database)
}

func main() {
	database := driver.ConnectDB()
	defer database.Close()

	args := os.Args
	if len(args) > 1 {
		first := args[1]
		// second := ""
		// if len(args) > 2 {
		// 	second = args[2]
		// }

		if first == "create" {
			create(database)
		} else if first == "seed" {
			seeds.Seed()
			os.Exit(0)
		} else if first == "migrate" {
			models.InitTable(database)
		}

		if first != "" {
			os.Exit(0)
		}
	}

	// Create the "static" folder if it doesn't already exist
	// if err := os.MkdirAll("static/images", 0755); err != nil {
	// 	log.Fatal("Failed to create static folder: ", err)
	// }

	server.StartServer()
}
