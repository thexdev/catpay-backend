package request

import "github.com/go-playground/validator/v10"

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func NewLoginRequest() *LoginRequest {
	return &LoginRequest{}
}

func (req *LoginRequest) Validate() (bool, error) {
	v := validator.New()

	if err := v.Struct(req); err != nil {
		return false, err
	}

	return true, nil
}

func (req *LoginRequest) Errors(err error) []any {
	errs := []any{}

	for _, err := range err.(validator.ValidationErrors) {
		errs = append(errs, map[string]string{
			"field":   err.Field(),
			"message": err.Tag(),
		})
	}

	return errs
}
