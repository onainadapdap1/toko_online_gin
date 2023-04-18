package driver

import (
	"fmt"
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/jinzhu/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() *gorm.DB {
	if DB != nil {
		return DB
	}

	dbConfig := Config.DB
	if dbConfig.Adapter == "mysqlonai" {
		DB, err = gorm.Open("mysql", "root:my-secret-pw-23@tcp(127.0.0.1:3306)/db_toko_online?charset=utf8mb4&parseTime=True&loc=Local")
		log.Println("Connected to Database Development")
	} else if dbConfig.Adapter == "postgresonai" {
		DB, err = gorm.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbConfig.UserDB, dbConfig.Password, dbConfig.Host, dbConfig.Name))
		log.Println("Connected to Database Local")
	}

	if err != nil {
		log.Println("[Driver.ConnectDB] error when connect to database")
		log.Fatal(err)
	} else {
		log.Println("SUCCESS CONNECT TO DATABASE")
	}

	return DB
}

func GetDB() *gorm.DB {
	return DB
}
