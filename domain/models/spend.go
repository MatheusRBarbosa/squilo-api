package models

import "time"

type Spend struct {
	ID        int
	UserId    int `gorm:"column:userId"`
	Value     float32
	CreatedAt time.Time `gorm:"column:createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt"`
}
