package utils

import (
	"toko_online_gin/dtos"
	"toko_online_gin/models"

	"github.com/gosimple/slug"
)

/*USER*/
// format user response register
func FormatUserRegister(user *models.User) dtos.UserRegisterFormatter {
	formatter := dtos.UserRegisterFormatter{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
		Role:     user.Role,
	}

	return formatter
}

func FormatUserLogin(user models.User, token string) dtos.UserLoginFormatter {
	formatter := dtos.UserLoginFormatter{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
		Role:     user.Role,
		Token:    token,
	}

	return formatter
}

/*CATEGORY*/

// format ketika insert category
type CategoryFormatter struct {
	ID          uint   `json:"id"`
	UserID      uint   `json:"user_id"`
	Name        string `json:"name" gorm:"uniqueIndex;not null"`
	Description string `json:"description"`
	Slug        string `json:"slug" gorm:"uniqueIndex;not null"`
	ImageURL    string `json:"image_url"`
}

func FormatCategory(category models.Category) CategoryFormatter {
	categoryFormatter := CategoryFormatter{
		ID:          category.ID,
		UserID:      category.UserID,
		Name:        category.Name,
		Description: category.Description,
		Slug:        category.Slug,
		ImageURL:    category.ImageURL,
	}

	return categoryFormatter
}

type CategoryDetailFormatter struct {
	ID          uint                  `json:"id"`
	UserID      uint                  `json:"user_id"`
	Name        string                `json:"category_name"`
	Description string                `json:"description"`
	Slug        string                `json:"slug"`
	ImageURL    string                `json:"image_url"`
	User        CategoryUserFormatter `json:"user"`
}

type CategoryUserFormatter struct {
	FullName string `json:"full_name"`
	Role     string `json:"role"`
}

func FormateCategoryDetail(category models.Category) CategoryDetailFormatter {
	categoryDetailFormatter := CategoryDetailFormatter{
		ID:          category.ID,
		UserID:      category.UserID,
		Name:        category.Name,
		Description: category.Description,
		Slug:        category.Slug,
		ImageURL:    category.ImageURL,
	}
	user := category.User

	categoryUserFormatter := CategoryUserFormatter{
		FullName: user.FullName,
		Role:     user.Role,
	}

	categoryDetailFormatter.User = categoryUserFormatter

	return categoryDetailFormatter
}

func FormatCategories(categories []models.Category) []CategoryDetailFormatter {
	categoriesFormatter := []CategoryDetailFormatter{}

	for _, category := range categories {
		categoryFormatter := FormateCategoryDetail(category)
		categoriesFormatter = append(categoriesFormatter, categoryFormatter)
	}

	return categoriesFormatter
}

/* PRODUCT */
type ProductFormatter struct {
	ID          uint    `json:"id"`
	UserID      uint    `json:"user_id"`
	Name        string  `json:"product_name"`
	Slug        string  `json:"slug"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	CategoryID  uint    `json:"category_id"`
	ImageURL    string  `json:"image_url"`
}

func FormatProduct(product models.Product) ProductFormatter {
	productFormatter := ProductFormatter{
		ID:          product.ID,
		UserID:      product.UserID,
		CategoryID:  product.CategoryID,
		Name:        product.Name,
		Slug:        slug.Make(product.Name),
		Description: product.Description,
		Price:       product.Price,
		Quantity:    product.Quantity,
		ImageURL:    product.ImageURL,
	}

	return productFormatter
}

type ProductDetailFormatter struct {
	ID          uint                     `json:"id"`
	Name        string                   `json:"product_name"`
	Description string                   `json:"description"`
	Price       float64                  `json:"price"`
	Quantity    int                      `json:"quantity"`
	ImageURL    string                   `json:"image_url"`
	UserID      uint                     `json:"user_id"`
	User        ProductUserFormatter     `json:"user"`
	CategoryID  uint                     `json:"category_id"`
	Category    ProductCategoryFormatter `json:"category"`
}

type ProductCategoryFormatter struct {
	ID          uint   `json:"category_id"`
	Name        string `json:"product_name"`
	Description string `json:"description"`
}
type ProductUserFormatter struct {
	ID       uint   `json:"user_id"`
	FullName string `json:"full_name"`
	Role     string `json:"role"`
}

func FormatProductDetail(product models.Product) ProductDetailFormatter {
	productDetailFormatter := ProductDetailFormatter{
		ID:          product.ID,
		UserID:      product.UserID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Quantity:    product.Quantity,
		CategoryID:  product.CategoryID,
		ImageURL:    product.ImageURL,
	}
	user := product.User
	productUserFormatter := ProductUserFormatter{
		ID:       user.ID,
		FullName: user.FullName,
		Role:     user.Role,
	}
	productDetailFormatter.User = productUserFormatter

	category := product.Category
	productCategoryFormatter := ProductCategoryFormatter{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	productDetailFormatter.Category = productCategoryFormatter

	return productDetailFormatter
}

func FormatProducts(products []models.Product) []ProductDetailFormatter {
	productsFormatter := []ProductDetailFormatter{}

	for _, product := range products {
		productFormatter := FormatProductDetail(product)
		productsFormatter = append(productsFormatter, productFormatter)
	}

	return productsFormatter
}
