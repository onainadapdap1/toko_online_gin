package repository

import (
	"fmt"
	"toko_online_gin/models"

	"github.com/jinzhu/gorm"
)

type ProductRepoInterface interface {
	CreateProduct(product models.Product) (models.Product, error)
	GetCategoryByID(id uint) (models.Category, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepoInterface {
	return &productRepository{db: db}
}

func (r *productRepository) CreateProduct(product models.Product) (models.Product, error) {
	tx := r.db.Begin()

	if err := tx.Debug().Create(&product).Error; err != nil {
		tx.Rollback()
		return product, fmt.Errorf("[CreateProduct.Insert] Error when query save data with : %w", err)
	}

	tx.Commit()

	return product, nil
}

func (r *productRepository)	GetCategoryByID(id uint) (models.Category, error) {
	var category models.Category
	tx := r.db.Begin()

	// SELECT c.*, u.*
	// FROM categories c
	// JOIN users u ON u.id = c.user_id
	// WHERE c.id = ?;
	err := tx.Debug().Preload("User").Where("id = ?", id).Find(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}
