package dtos

import "time"

type VaultDto struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Total       float32   `json:"total"`
	TypeId      int       `json:"typeId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	TotalLabel  string    `json:"totalLabel"`

	Type IdName `json:"type"`
}
