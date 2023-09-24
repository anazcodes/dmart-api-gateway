package payload

import (
	"log"
	"testing"
)

func TestCreateProductRequest(t *testing.T) {
	s := CreateProductRequest{
		CategoryID:  "650d84646bd31094bb99bd51",
		Name:        "product",
		Description: "description",
		Price:       1,
	}
	err := ValidateStruct(s)
	if err != nil {
		log.Fatalln(err)
	}
}
