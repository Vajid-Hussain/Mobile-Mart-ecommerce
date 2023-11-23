package domain

import "time"

type Order struct {
	ID            uint `gorm:"primary key"`
	UserID        uint
	User          Users `gorm:"foreignkey:UserID;association_foreignkey:ID"`
	AddressID     uint
	Location      Address `gorm:"foreignkey:AddressID;association_foreignkey:ID"`
	InventoryID   uint
	Product       Inventories `gorm:"foreignkey:InventoryID;association_foreignkey:ID"`
	SellerID      uint        `gorm:"not null"`
	Seller        Seller      `gorm:"forgienKey:SellerID;association_foreignkey:ID"`
	PaymentMethod string
	Quantity      uint
	Price         uint
	OrderDate     time.Time
	DeliveryDate  time.Time
	PaymentStatus status `gorm:"default:pending"`
	OrderStatus   string `gorm:"default:processing"`
}
