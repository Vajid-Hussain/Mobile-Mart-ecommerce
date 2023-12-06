package domain

import "time"

type Category struct {
	ID     uint   `gorm:"unique key"`
	Name   string `gorm:"unique"`
	Status status `gorm:"default:active"`
}

type Brand struct {
	ID     uint   `gorm:"unique key"`
	Name   string `gorm:"unique"`
	Status status `gorm:"default:active"`
}

type CategoryOffer struct {
	ID               uint
	Title            string
	CategoryID       uint
	Categoryid       Category `gorm:"foreignkey:CategoryID;association_foreignkey:ID"`
	SellerID         uint
	Sellerid         Seller `gorm:"foreignkey:SellerID;association_foreignkey:ID"`
	CategoryDiscount uint
	StartDate        time.Time
	EndDate          time.Time
	Status           status ` gorm:"default:active"`
}
