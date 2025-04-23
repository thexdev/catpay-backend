package port

type LoginRequest struct {
	Email    string
	Password string
}

type LoginResponse struct {
	Ok bool
}

type RegisterRequest struct {
	Email           string
	Password        string
	PasswordConfirm string
}

type RegisterResponse struct {
	Ok bool
}

type Auth interface {
	Login(req LoginRequest) (LoginResponse, error)
	Register(req RegisterRequest) (RegisterResponse, error)
}
