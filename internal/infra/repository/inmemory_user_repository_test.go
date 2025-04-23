package repository_test

import (
	"catpay/internal/infra/repository"
	"catpay/internal/infra/repository/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInMemoryUserRepository(t *testing.T) {
	t.Run("initialize without error", func(t *testing.T) {
		repo := repository.NewInMemoryRepository()

		assert.IsType(t, &repository.InMemoryUserRepository{}, repo)
	})
}

func TestInMemoryUserRepository_Create(t *testing.T) {
	t.Run("create new user", func(t *testing.T) {
		repo := repository.NewInMemoryRepository()

		err := repo.Create("another_user@gmail.com", "password", "user")

		assert.NoError(t, err)
	})
}

func TestInMemoryUserRepository_GetHashedPasswordByEmail(t *testing.T) {
	t.Run("returns ErrUserNotFound when the given user's email doesn't exists", func(t *testing.T) {
		var errUserNotFound *entity.ErrUserNotFound

		repo := repository.NewInMemoryRepository()

		hash, err := repo.GetHashedPasswordByEmail("notexist_user@mail.com")

		assert.ErrorAs(t, err, &errUserNotFound)
		assert.IsType(t, "", hash)
	})

	t.Run("returns user hashed password", func(t *testing.T) {
		repo := repository.NewInMemoryRepository()

		hash, err := repo.GetHashedPasswordByEmail("user1@example.com")

		assert.NoError(t, err)
		assert.IsType(t, "", hash)
	})
}
