package domain

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
	CategoryID       uint
	Categoryid       Category `gorm:"foreignkey:CategoryID;association_foreignkey:ID"`
	SellerID         uint
	Sellerid         Seller `gorm:"foreignkey:SellerID;association_foreignkey:ID"`
	CategoryDiscount uint   `gorm:"default:0"`
}
