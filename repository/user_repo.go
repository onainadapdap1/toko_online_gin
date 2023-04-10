package repository

import (
	"fmt"
	"toko_online_gin/models"

	"github.com/jinzhu/gorm"
)

type UserRepoInterface interface {
	RegisterUser(user *models.User) (*models.User, error)
	FindByEmail(email string) (models.User, error)
	FindByID(ID uint) (models.User, error)
	SaveAvatar(user models.User) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepoInterface {
	return &userRepository{db: db}
}

// register user
func (r *userRepository) RegisterUser(user *models.User) (*models.User, error) {
	tx := r.db.Begin()

	if err := tx.Debug().Create(&user).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("[RegisterAdminRepoImpl.Insert] Error when query save data with : %w", err)
	}
	tx.Commit()

	return user, nil
}

func (r *userRepository) FindByEmail(email string) (models.User, error) {
	var user models.User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindByID(ID uint) (models.User, error) {
	var user models.User

	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) SaveAvatar(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
