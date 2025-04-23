package usecase_test

import (
	"catpay/internal/app/service"
	"catpay/internal/app/usecase"
	"catpay/internal/infra/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginUseCase(t *testing.T) {
	var (
		userRepo       = repository.NewInMemoryRepository()
		passwordHasher = service.NewPasswordHasher()

		wrongLoginReq = usecase.LoginRequest{
			Email:    "user1@example.com", // available user
			Password: "password123",       // incorrect password
		}
		validLoginReq = usecase.LoginRequest{
			Email:    "user1@example.com", // available user
			Password: "user1_password",    // correct password
		}
	)

	t.Run("intialize without error", func(t *testing.T) {
		loginUseCase := usecase.NewLoginUseCase(
			userRepo,
			*passwordHasher,
		).SetCredential(wrongLoginReq)

		assert.IsType(t, &usecase.LoginUseCase{}, loginUseCase)
	})

	t.Run("returns true when login success", func(t *testing.T) {
		loginUseCase := usecase.NewLoginUseCase(
			userRepo,
			*passwordHasher,
		).SetCredential(validLoginReq)

		ok, err := loginUseCase.Execute()

		assert.NoError(t, err)
		assert.True(t, ok)
	})

	t.Run("returns error when login failed", func(t *testing.T) {
		loginUseCase := usecase.NewLoginUseCase(
			userRepo,
			*passwordHasher,
		).SetCredential(wrongLoginReq)

		ok, err := loginUseCase.Execute()

		assert.NoError(t, err)
		assert.False(t, ok)
	})
}
