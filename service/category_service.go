package service

import (
	"toko_online_gin/models"
	"toko_online_gin/repository"
)

type CategoryServiceInterface interface {
	CreateCategory(input models.CreateCategoryInput) (models.Category, error)
	UpdateCategory(inputSlug models.GetCategoryDetailInput, inputData models.CreateCategoryInput) (models.Category, error)
	FindBySlug(inputSlug models.GetCategoryDetailInput) (models.Category, error) 
	FindAllCategory() ([]models.Category, error)
}

type categoryService struct {
	repo repository.CategoryRepoInterface
}

func NewCategoryService(repo repository.CategoryRepoInterface) CategoryServiceInterface {
	return &categoryService{repo: repo}
}

func (s *categoryService) CreateCategory(input models.CreateCategoryInput) (models.Category, error) {
	category := models.Category{
		UserID: input.User.ID,
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

func (s *categoryService) FindAllCategory() ([]models.Category, error) {
	categories, err := s.repo.FindAllCategory()
	if err != nil {
		return categories, err
	}

	return categories, nil
}