package domain

import "time"

type Wallet struct {
	ID        int16
	UserID    int16
	UUID      string
	Balance   float32
	CreatedAt time.Time
	UpdatedAt time.Time
}
