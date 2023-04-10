package repository

import (
	"fmt"
	"toko_online_gin/models"

	"github.com/jinzhu/gorm"
)

type CategoryRepoInterface interface {
	CreateCategory(category models.Category) (models.Category, error)
	UpdateCategory(category models.Category) (models.Category, error)
	FindBySlug(slug string) (models.Category, error)
	FindAllCategory() ([]models.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepoInterface {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) CreateCategory(category models.Category) (models.Category, error) {
	tx := r.db.Begin()

	if err := tx.Debug().Create(&category).Error; err != nil {
		tx.Rollback()
		return category, fmt.Errorf("[CreateCategory.Insert] Error when query save data with : %w", err)
	}

	tx.Commit()

	return category, nil
}

func (r *categoryRepository) UpdateCategory(category models.Category) (models.Category, error) {
	tx := r.db.Begin()
	err := tx.Debug().Save(&category).Error
	if err != nil {
		tx.Rollback()
		return category, err
	}
	tx.Commit()
	return category, nil
}

func (r *categoryRepository) FindBySlug(slug string) (models.Category, error) {
	var category models.Category
	tx := r.db.Begin()
	err := tx.Debug().Preload("User").Where("slug = ?", slug).Find(&category).Error
	// err := r.db.Where("slug = ?", slug).Find(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *categoryRepository) FindAllCategory() ([]models.Category, error) {
	var categories []models.Category
	tx := r.db.Begin()
	err := tx.Debug().Preload("User").Find(&categories).Error
	if err != nil {
		return categories, err
	}

	return categories, nil
}
