package entity

import (
	"time"
)

type ErrUserNotFound struct {
}

func (err *ErrUserNotFound) Error() string {
	return "user with the given email does not exists"
}

type ErrUserAlreadyExist struct {
}

func (err *ErrUserAlreadyExist) Error() string {
	return "user with the given email is already exists"
}

type User struct {
	ID           int16
	UUID         string
	Email        string
	PasswordHash string
	Role         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
