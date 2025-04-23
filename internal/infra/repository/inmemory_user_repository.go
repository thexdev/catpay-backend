package repository

import (
	"catpay/internal/infra/repository/entity"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type InMemoryUserRepository struct {
	// just fake users for testing...
	users []entity.User
}

func NewInMemoryRepository() *InMemoryUserRepository {
	// Seed some fake users
	fakeUsers := []entity.User{
		{
			ID:           1,
			UUID:         uuid.NewString(),
			Email:        "user1@example.com",
			PasswordHash: hashPassword("user1_password"),
			Role:         "user",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}

	return &InMemoryUserRepository{
		users: fakeUsers,
	}
}

func (r *InMemoryUserRepository) Create(email, password, role string) error {
	lastUser := r.users[len(r.users)-1]

	r.users = append(r.users, entity.User{
		ID:           lastUser.ID + 1,
		UUID:         uuid.NewString(),
		Email:        email,
		PasswordHash: hashPassword(password),
		Role:         role,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	})

	return nil
}

func (repo *InMemoryUserRepository) Exist(email string) error {
	return nil
}

func (repo *InMemoryUserRepository) GetHashedPasswordByEmail(
	email string,
) (string, error) {
	for _, user := range repo.users {
		if user.Email == email {
			return user.PasswordHash, nil
		}
	}

	return "", &entity.ErrUserNotFound{}
}

func hashPassword(plain string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(plain), 8)
	return string(hash)
}
