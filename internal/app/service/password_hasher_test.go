package service_test

import (
	"catpay/internal/app/service"
	"catpay/internal/infra/repository"
	"catpay/internal/infra/repository/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasswordHasher(t *testing.T) {
	t.Run("initialize without error", func(t *testing.T) {
		s := service.NewPasswordHasher()

		assert.IsType(t, &service.PasswordHasher{}, s)
	})
}

func TestPasswordHasher_Make(t *testing.T) {
	t.Run("returns hashed password", func(t *testing.T) {
		hasher := service.NewPasswordHasher()

		hashed, err := hasher.Make("new password")

		assert.NoError(t, err)
		assert.IsType(t, "", hashed)
	})
}

func TestPasswordHasher_Verify(t *testing.T) {
	t.Run("returns true when password match", func(t *testing.T) {
		var (
			repo    = repository.NewInMemoryRepository()
			service = service.NewPasswordHasher()
		)

		hashed, err := repo.GetHashedPasswordByEmail("user1@example.com")

		assert.NoError(t, err)
		assert.True(t, service.Verify("user1_password", hashed))
	})

	t.Run("returns false when password unmatch", func(t *testing.T) {
		var (
			repo    = repository.NewInMemoryRepository()
			service = service.NewPasswordHasher()
		)

		hashed, err := repo.GetHashedPasswordByEmail("user2@example.com")

		var errUserNotFound *entity.ErrUserNotFound
		assert.ErrorAs(t, err, &errUserNotFound)

		assert.False(t, service.Verify("password", hashed))
	})
}
