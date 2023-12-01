package domain

import "time"

type Coupons struct {
	ID              uint `gorm:"unique key"`
	Description     string
	Type            string
	Discount        uint
	MinimumRequired uint
	MaximumAllowed  uint
	StartDate       time.Time
	EndDate         time.Time
	Status          status `json:"status" gorm:"default:ACTIVE"`
}
