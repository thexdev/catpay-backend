package auth_test

import (
	"bytes"
	"catpay/internal/infra/http/handler"
	"catpay/internal/infra/http/request"
	"catpay/test/mock/repository"
	"catpay/test/mock/service"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin(t *testing.T) {
	var (
		mockUserRepo         = repository.NewMockUserRepository()
		mockBcryptPassHasher = service.NewMockBcryptPasswordHasher()

		handler = handler.NewAuthHandler(mockUserRepo, mockBcryptPassHasher)

		app = fiber.New()
	)

	app.Post("/login", handler.Login)

	t.Run("returns http 422", func(t *testing.T) {
		t.Run("when sending empty JSON", func(t *testing.T) {
			req := httptest.NewRequest("POST", "/login", nil)

			res, err := app.Test(req)
			assert.NoError(t, err)

			assert.Equal(t, 422, res.StatusCode)
		})
	})

	t.Run("returns http 404", func(t *testing.T) {
		t.Run("when the given email doesn't exist", func(t *testing.T) {
			t.Cleanup(func() {
				mockUserRepo.Calls = nil
				mockUserRepo.ExpectedCalls = nil
			})

			// Mimic PostgreSQL error (no rows). In another word, no record was
			// found with the give email
			mockUserRepo.On(
				"GetHashedPasswordByEmail",
				mock.Anything,
			).Return("abcd1234", errors.New("no row"))

			credential := request.LoginRequest{
				Email:    faker.Email(),
				Password: "verystrong_password123",
			}

			payload, err := json.Marshal(credential)
			assert.NoError(t, err)

			req := httptest.NewRequest(
				"POST",
				"/login",
				bytes.NewBuffer(payload),
			)
			req.Header.Set("Content-Type", "application/json")

			res, err := app.Test(req)
			assert.NoError(t, err)

			assert.Equal(t, 404, res.StatusCode)

			mockUserRepo.AssertNumberOfCalls(t, "GetHashedPasswordByEmail", 1)
		})
	})

	t.Run("returns http 401", func(t *testing.T) {
		t.Run("when user credential doesn't match", func(t *testing.T) {
			t.Cleanup(func() {
				mockUserRepo.Calls = nil
				mockUserRepo.ExpectedCalls = nil

				mockBcryptPassHasher.Calls = nil
				mockBcryptPassHasher.ExpectedCalls = nil
			})

			userHashedPassword := "abcd1234"

			mockUserRepo.On(
				"GetHashedPasswordByEmail",
				mock.Anything,
			).Return(userHashedPassword, nil)

			// Mock the Verify() function to returns false. This means the
			// `plain password` doesn't match with the `hashed password`
			mockBcryptPassHasher.On("Verify", mock.Anything).Return(false)

			credential := request.LoginRequest{
				Email:    faker.Email(),
				Password: "wrong_password",
			}

			payload, err := json.Marshal(credential)
			assert.NoError(t, err)

			req := httptest.NewRequest(
				"POST",
				"/login",
				bytes.NewBuffer(payload),
			)
			req.Header.Set("Content-Type", "application/json")

			res, err := app.Test(req)
			assert.NoError(t, err)

			assert.Equal(t, 401, res.StatusCode)

			mockUserRepo.AssertNumberOfCalls(t, "GetHashedPasswordByEmail", 1)

			mockBcryptPassHasher.AssertNumberOfCalls(t, "Verify", 1)
		})
	})
}
