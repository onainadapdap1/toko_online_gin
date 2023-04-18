package models

import (
	"log"
	"toko_online_gin/helpers"
	// "toko_online_gin/helpers"

	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

// generate new user table
type User struct {
	gorm.Model
	FullName       string 
	Email          string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password       string `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimun length of 6 characters"`
	Role           string `json:"role"`
	AvatarFileName string
	Category []Category
	Product []Product
}

// naming convention
func (u *User) TableName() string {
	return "tb_users"
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password, err = helpers.HassPass(u.Password)
	if err != nil {
		log.Println("error while hashing password")
		return
	}
	err = nil
	return
}
