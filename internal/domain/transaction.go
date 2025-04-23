package domain

import "time"

type Transaction struct {
	ID        int16
	UUID      string
	Amount    float32
	Status    string
	CreatdAt  time.Time
	UpdatedAt time.Time
}
