package service

import (
	"golang.org/x/crypto/bcrypt"
)

type BcryptPasswordHasher struct {
}

func NewBcryptPasswordHasher() *BcryptPasswordHasher {
	return &BcryptPasswordHasher{}
}

func (ph *BcryptPasswordHasher) Make(plain string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plain), 8)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func (ph *BcryptPasswordHasher) Verify(plain string, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err == nil
}
