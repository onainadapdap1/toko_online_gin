package service

import (
	"toko_online_gin/models"
	"toko_online_gin/repository"
)

type CategoryServiceInterface interface {
	CreateCategory(input models.CreateCategoryInput) (models.Category, error)
	UpdateCategory(inputSlug models.GetCategoryDetailInput, inputData models.CreateCategoryInput) (models.Category, error)
	FindBySlug(inputSlug models.GetCategoryDetailInput) (models.Category, error)
	FindByCategoryID(categoryID uint) (models.Category, error)
	FindAllCategory() ([]models.Category, error)
	DeleteCategory(category models.Category) error
	// FindAllProductByCategory(categoryID uint) (models.Category, error)
	// DeleteAllProductByCategory(products []models.Product) error
	// DeleteCategoryByID(categoryID uint) error
}

type categoryService struct {
	repo repository.CategoryRepoInterface
}

func NewCategoryService(repo repository.CategoryRepoInterface) CategoryServiceInterface {
	return &categoryService{repo: repo}
}

func (s *categoryService) CreateCategory(input models.CreateCategoryInput) (models.Category, error) {
	category := models.Category{
		UserID:      input.User.ID,
		Name:        input.Name,
		Description: input.Description,
		ImageURL:    input.ImageURL,
	}

	category, err := s.repo.CreateCategory(category)
	if err != nil {
		return category, err
	}

	return category, nil
}

func (s *categoryService) UpdateCategory(inputSlug models.GetCategoryDetailInput, inputData models.CreateCategoryInput) (models.Category, error) {
	category, err := s.repo.FindBySlug(inputSlug.Slug)
	if err != nil {
		return category, err
	}

	category.Name = inputData.Name
	category.Description = inputData.Description
	category.ImageURL = inputData.ImageURL

	updatedCategory, err := s.repo.UpdateCategory(category)
	if err != nil {
		return updatedCategory, err
	}

	return updatedCategory, nil
}

func (s *categoryService) FindBySlug(inputSlug models.GetCategoryDetailInput) (models.Category, error) {
	category, err := s.repo.FindBySlug(inputSlug.Slug)
	if err != nil {
		return category, err
	}

	return category, nil
}

func (s *categoryService) FindByCategoryID(categoryID uint) (models.Category, error) {
	category, err := s.repo.FindByCategoryID(categoryID)
	if err != nil {
		return category, err
	}

	return category, nil
}

func (s *categoryService) FindAllCategory() ([]models.Category, error) {
	categories, err := s.repo.FindAllCategory()
	if err != nil {
		return categories, err
	}

	return categories, nil
}

func (s *categoryService) DeleteCategory(category models.Category) error {
	// if err := s.repo.DeleteCategoryProducts(category); err != nil {
	// 	return err
	// }
	if err := s.repo.DeleteCategory(category); err != nil {
		return err
	}
	return nil
}

// func (s *categoryService) FindAllProductByCategory(slug string) ([]models.Product, error) {
// func (s *categoryService) FindAllProductByCategory(categoryID uint) (models.Category, error) {
// 	category, err := s.repo.FindAllProductByCategory(categoryID)
// 	if err != nil {
// 		return category, err
// 	}

// 	return category, nil
// }

// func (s *categoryService) DeleteAllProductByCategory(products []models.Product) error {
// 	if err := s.repo.DeleteAllProductByCategory(products); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (s *categoryService) DeleteAllProductByCategory(categoryID uint) error {
// 	category, err := s.repo.FindAllProductByCategory(categoryID)
// 	if err != nil {
// 		return err
// 	}

// 	err = s.repo.DeleteAllProductByCategory(category)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (s *categoryService) DeleteCategoryByID(categoryID uint) error {
// 	category, err := s.repo.FindByCategoryID(categoryID)
// 	if err != nil {
// 		return err
// 	}

// 	err = s.repo.DeleteCategoryByID(category)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
