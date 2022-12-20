package dtos

import "time"

type TransactionDto struct {
	Id          int       `json:"id"`
	VaultId     int       `json:"vaultId"`
	Date        time.Time `json:"date"`
	Observation string    `json:"observation"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
