package service

import (
	"github.com/stretchr/testify/mock"
)

type MockBcryptPasswordHasher struct {
	mock.Mock
}

func NewMockBcryptPasswordHasher() *MockBcryptPasswordHasher {
	return &MockBcryptPasswordHasher{}
}

func (ph *MockBcryptPasswordHasher) Make(plain string) (string, error) {
	args := ph.Called()
	return args.String(0), args.Error(1)
}

func (ph *MockBcryptPasswordHasher) Verify(plain string, hashed string) bool {
	args := ph.Called()
	return args.Bool(0)
}
