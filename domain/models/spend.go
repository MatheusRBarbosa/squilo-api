package models

import "time"

type Spend struct {
	ID       int
	UserId   int `gorm:"column:userId"`
	Value    float32
	CreateAt time.Time `gorm:"column:createAt"`
}
