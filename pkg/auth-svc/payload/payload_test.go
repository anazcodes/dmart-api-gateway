package payload

import (
	"log"
	"testing"
)

func TestCreateAccountRequest(t *testing.T) {
	s := CreateAccountRequest{
		Username:        "anas",
		Email:           "anas@gmail.com",
		Phone:           1234567890,
		Password:        "12345678",
		ConfirmPassword: "12345678",
	}
	err := ValidateStruct(s)
	if err != nil {
		log.Fatalln(err)
	}
}
