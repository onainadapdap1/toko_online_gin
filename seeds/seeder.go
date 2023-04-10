package seeds

import (
	"toko_online_gin/driver"
	"toko_online_gin/helpers"
	"toko_online_gin/models"

	"github.com/jinzhu/gorm"
)

func seedAdmin(db *gorm.DB) {
	adminRoleUsers := 0
	tx := db.Begin()
	tx.Model(&models.User{}).Where("role = ?", "admin").Count(&adminRoleUsers)
	if adminRoleUsers == 0 {
		password, _ := helpers.HassPass("password")
		user := models.User{
			FullName: "AdminFN", 
			Email: "admin@gmail.com",
			Password: password,
			Role: "admin",
		}

		tx.Set("gorm:association_autoupdate", false).Debug().Create(&user)

		if tx.Error != nil {
			tx.Rollback()
			print(db.Error)
		}
	}
	tx.Commit()
}

func Seed() {
	db := driver.ConnectDB()
	seedAdmin(db)
}