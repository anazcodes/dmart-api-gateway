package payload

import "github.com/go-playground/validator/v10"

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required"`
}

type CreateProductRequest struct {
	CategoryID  string   `json:"category_id" validate:"required"`
	Name        string   `json:"name" validate:"required,min=3"`
	Description string   `json:"description" validate:"required,min=10"`
	Image       []string `json:"image"`
	Price       int64    `json:"price" validate:"required,gt=0"`
}

// Validates the struct using validate tags, if the validation fails it will return false else true
func ValidateStruct(s any) error {
	validate := validator.New()
	return validate.Struct(s)
}
