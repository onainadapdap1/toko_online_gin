package models

/* USER INPUT*/ 

// register user input field
type RegisterUserInput struct {
	FullName       string `json:"full_name" binding:"required"` 
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
}

// login user input field
type LoginUserInput struct {
	Email string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

// CheckEmailInput
type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}

/* END USER INPUT */ 

/* CATEGORY INPUT */ 

// create category input
type CreateCategoryInput struct {
	Name string	`gorm:"not null" form:"name" json:"name"`
	Description string `gorm:"not null" form:"description" json:"description"`
	ImageURL string `gorm:"not null" form:"image" json:"image"`
	User User
}

type GetCategoryDetailInput struct {
	Slug string `uri:"slug" binding:"required"`
}

/* END CATEGORY INPUT */ 

/* PRODUCT INPUT */ 
type CreateProductInput struct {
	Name string `gorm:"not null" form:"name" json:"name"`
	Description string `gorm:"not null" form:"description" json:"description"`
	Price float64 `gorm:"not null" form:"price" json:"price"`
	Quantity int `gorm:"not null" form:"quantity" json:"quantity"`
	ImageURL string `gorm:"not null" form:"image" json:"image"`
	CategoryID uint `gorm:"not null" form:"category_id" json:"category_id"`
	User User
	Category Category
}

type GetProductDetailInput struct {
	Slug string `uri:"slug" binding:"required"`
}
/* END PRODUCT INPUT */ 