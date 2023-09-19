// payload is a request response struct package
package payload

type AdminLoginRequest struct {
	Username string
	Password string
}

type AdminLoginResponse struct {
	Status int64
	Msg    string
	Error  string
}

type CreateAccountRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Phone     int64  `json:"phone"`
	Password  string `json:"password"`
	Password2 string `json:"password2"`
}

type CreateAccountResponse struct {
	Status int64
	Msg    string
	Error  string
}

type UserLoginRequest struct {
	LoginInput string // user can login using email, phone or using username
	Password   string
}

type UserLoginResponse struct {
	Status int64
	Msg    string
	Error  string
}
