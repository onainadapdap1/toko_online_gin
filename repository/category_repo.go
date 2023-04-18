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
	FindByCategoryID(categoryID uint) (models.Category, error)
	FindAllCategory() ([]models.Category, error)
	FindAllProductByCategory(categoryID uint) (models.Category, error) 
	// DeleteCategoryProducts(category models.Category) error
	DeleteCategory(category models.Category) error
	// DeleteAllProductByCategory(products []models.Product) error 
	// DeleteCategoryByID(category models.Category) error 
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

func (r *categoryRepository) FindByCategoryID(categoryID uint) (models.Category, error) {
	var category models.Category
	tx := r.db.Begin()
	if err := tx.Debug().Where("id = ?", categoryID).First(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}

func (r *categoryRepository) FindAllProductByCategory(categoryID uint) (models.Category, error) {
	var category models.Category
	tx := r.db.Begin()
	// 	if err := tx.Debug().Model(models.Category{}).Association("Products").Where("slug = ?", slug).Find(&products).Error; err != nil {
	// if err := tx.Debug().Where("category_id = ?", categoryID).Find(&products).Error; err != nil {
	
	if err := tx.Debug().Preload("Products").First(&category, categoryID).Error; err != nil {
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

// func (r *categoryRepository) DeleteCategoryProducts(category models.Category) error {
// 	tx := r.db.Begin()

// 	if err := tx.Debug().Model(category).Association("Products").Clear().Error; err != nil {
// 		return err
// 	}

// 	return nil
// }

func (r *categoryRepository) DeleteCategory(category models.Category) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("category_id = ?", category.ID).Delete(&models.Product{}).Error; err != nil {
			return err
		}

		if err := tx.Delete(category).Error; err != nil {
			return err
		}

		return nil
	})

	

	// if err := tx.Debug().Unscoped().Delete(category).Error; err != nil {
	// 	return err
	// }

	// return nil
}

// func (r *categoryRepository) DeleteAllProductByCategory(products []models.Product) error {
// 	tx := r.db.Begin()

// 	if err := tx.Debug().Unscoped().Delete(&products).Error; err != nil {
// 		tx.Rollback()
// 		return err
// 	}
// 	tx.Commit()
// 	return nil
// }

// func (r *categoryRepository) DeleteCategoryByID(category models.Category) error {
// 	tx := r.db.Begin()
// 	err := tx.Debug().Unscoped().Delete(&category).Error
// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}
// 	tx.Commit()
// 	return nil
// }