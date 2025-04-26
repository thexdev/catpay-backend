package repository

import "github.com/stretchr/testify/mock"

type MockUserRepository struct {
	mock.Mock
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{}
}

func (mr *MockUserRepository) Create(email, password, role string) error {
	args := mr.Called()
	return args.Error(0)
}

func (mr *MockUserRepository) GetHashedPasswordByEmail(email string) (string, error) {
	args := mr.Called()
	return args.String(0), args.Error(1)
}

func (mr *MockUserRepository) Exist(email string) error {
	args := mr.Called()
	return args.Error(0)
}
