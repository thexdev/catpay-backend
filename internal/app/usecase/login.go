package usecase

import (
	"catpay/internal/app/port"
	"catpay/internal/app/service"
)

type LoginRequest struct {
	Email    string
	Password string
}

type LoginUseCase struct {
	userRepo       port.UserRepository
	passwordHasher service.PasswordHasher

	req LoginRequest
}

func NewLoginUseCase(
	userRepo port.UserRepository,
	passwordHasher service.PasswordHasher,
) *LoginUseCase {

	return &LoginUseCase{
		userRepo:       userRepo,
		passwordHasher: passwordHasher,
	}
}

func (uc *LoginUseCase) SetCredential(req LoginRequest) *LoginUseCase {
	uc.req = req
	return uc
}

func (uc *LoginUseCase) Execute() (bool, error) {
	password, err := uc.userRepo.GetHashedPasswordByEmail(uc.req.Email)

	if err != nil {
		return false, err
	}

	ok := uc.passwordHasher.Verify(uc.req.Password, password)
	if !ok {
		return false, nil
	}

	return true, nil
}
