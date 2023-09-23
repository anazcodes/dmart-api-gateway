package payload

import (
	"log"
	"testing"
)

func TestCreateProductRequest(t *testing.T) {
	s := CreateProductRequest{
		CategoryID:  "jdksfjl",
		Name:        "product",
		Description: "description",
		Price:       1,
	}
	err := ValidateStruct(s)
	if err != nil {
		log.Fatalln(err)
	}
}
