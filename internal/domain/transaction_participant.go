package domain

import "time"

type TransactionParticipant struct {
	ID              int16
	TransactionUUID string
	WalletID        int16
	Direction       int16
	CreatedAt       time.Time
}
