package payload

type CreateCategoryRequest struct {
	Name string `json:"name"`
}

type CreateProductRequest struct {
	CategoryID  string   `json:"category_id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Image       []string `json:"image"`
	Price       int64    `json:"price"`
}
