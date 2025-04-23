package request

import "github.com/go-playground/validator/v10"

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func NewRegisterRequest() *RegisterRequest {
	return &RegisterRequest{}
}

func (req *RegisterRequest) Validate() (bool, error) {
	v := validator.New()

	if err := v.Struct(req); err != nil {
		return false, err
	}

	return true, nil
}

func (req *RegisterRequest) Errors(err error) []any {
	errs := []any{}

	for _, err := range err.(validator.ValidationErrors) {
		errs = append(errs, map[string]string{
			"field":   err.Field(),
			"message": err.Tag(),
		})
	}

	return errs
}
