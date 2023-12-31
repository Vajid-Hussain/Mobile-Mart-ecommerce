package domain

import "time"

type Coupons struct {
	ID              uint `gorm:"unique key"`
	Name            string
	Type            string
	Discount        uint
	MinimumRequired uint
	MaximumAllowed  uint
	StartDate       time.Time
	EndDate         time.Time
	Status          status `gorm:"default:active"`
}
