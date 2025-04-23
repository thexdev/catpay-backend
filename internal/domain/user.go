package domain

import "time"

type User struct {
	ID           int16
	UUID         string
	Email        string
	PasswordHash string
	Role         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
