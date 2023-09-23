// payload is a request response struct package
package payload

import "github.com/go-playground/validator/v10"

type AdminLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AdminLoginResponse struct {
	Status int64
	Msg    string
	Error  string
}

type CreateAccountRequest struct {
	Username        string `json:"username" validate:"required,min=3"`
	Email           string `json:"email" validate:"required,email"`
	Phone           int64  `json:"phone" validate:"required,min=10"`
	Password        string `json:"password" validate:"required,min=8,eqfield=ConfirmPassword"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=8,eqfield=Password"`
}

type CreateAccountResponse struct {
	Status int64
	Msg    string
	Error  string
}

type UserLoginRequest struct {
	LoginInput string `json:"login_input" validate:"required"` // user can login using email, phone or using username
	Password   string `json:"password" validate:"required,min=8"`
}

type UserLoginResponse struct {
	Status int64
	Msg    string
	Error  string
}

// Validates the struct using validate tags, if the validation fails it will return false else true
func ValidateStruct(s any) error {
	validate := validator.New()
	return validate.Struct(s)
}
