package service

import (
	"toko_online_gin/models"
	"toko_online_gin/repository"
)

type ProductServiceInterface interface {
	CreateProduct(input models.CreateProductInput) (models.Product, error)
	GetCategoryByID(inputID uint) (models.Category, error)
}

type productService struct {
	repo repository.ProductRepoInterface
}

func NewProductService(repo repository.ProductRepoInterface) ProductServiceInterface {
	return &productService{repo: repo}
}

func (s *productService) CreateProduct(input models.CreateProductInput) (models.Product, error) {
	product := models.Product{
		CategoryID: input.CategoryID,
		Name: input.Name,
		Description: input.Description,
		Price: input.Price,
		Quantity: input.Quantity,
		ImageURL: input.ImageURL,
	}

	product, err := s.repo.CreateProduct(product)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *productService) GetCategoryByID(inputID uint) (models.Category, error) {
	category, err := s.repo.GetCategoryByID(inputID)
	if err != nil {
		return category, err
	}

	return category, nil

}