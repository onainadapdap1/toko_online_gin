package models

import "github.com/jinzhu/gorm"

func InitTable(db *gorm.DB) {
	db.Debug().AutoMigrate(&User{})
	db.Debug().AutoMigrate(&Category{})
	db.Debug().AutoMigrate(&Product{})
}
