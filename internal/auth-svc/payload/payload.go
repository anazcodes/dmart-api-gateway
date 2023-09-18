// payload is a request response struct package
package payload

type AdminLoginRequest struct {
	username string
	password string
}

type AdminLoginResponse struct {
	status int64
	msg    string
	error  string
}

type CreateAccountRequest struct {
	username string
	email    string
	phone    int64
	password string
}

type CreateAccountResponse struct {
	status int64
	msg    string
	error  string
}

type UserLoginRequest struct {
	loginInput string // user can login using email, phone or using username
	password   string
}

type UserLoginResponse struct {
	status int64
	msg    string
	error  string
}
