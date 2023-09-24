package payload

import "github.com/go-playground/validator/v10"

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required" binding:"required"`
}

type CreateProductRequest struct {
	CategoryID  string `json:"category_id" validate:"required" binding:"required"`
	Name        string `json:"name" validate:"required,min=3" binding:"required"`
	Description string `json:"description" validate:"required,min=10" binding:"required"`
	Price       int64  `json:"price" validate:"required,gt=0" binding:"required"`
}

// Validates the struct using validate tags, if the validation fails it will return false else true
func ValidateStruct(s any) error {
	validate := validator.New()
	return validate.Struct(s)
}
