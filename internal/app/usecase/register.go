package usecase

import (
	"catpay/internal/app/port"
	"catpay/internal/app/service"
	"fmt"
)

type RegisterRequest struct {
	Email    string
	Password string
}

type RegisterUseCase struct {
	userRepo       port.UserRepository
	passwordHasher service.PasswordHasher

	req RegisterRequest
}

func NewRegisterUseCase(useRepo port.UserRepository, passwordHasher service.PasswordHasher) *RegisterUseCase {
	return &RegisterUseCase{
		userRepo:       useRepo,
		passwordHasher: passwordHasher,
	}
}

func (uc *RegisterUseCase) SetCredential(req RegisterRequest) *RegisterUseCase {
	uc.req = req
	return uc
}

func (uc *RegisterUseCase) Execute() (bool, error) {
	if err := uc.userRepo.Exist(uc.req.Email); err != nil {
		fmt.Println(err, "exist")
		return false, err
	}

	hashedPassword, err := uc.passwordHasher.Make(uc.req.Password)
	if err != nil {
		return false, err
	}

	err = uc.userRepo.Create(uc.req.Email, hashedPassword, "user")
	fmt.Println(err, "usecaes")
	if err != nil {
		return false, err
	}

	return true, nil
}
