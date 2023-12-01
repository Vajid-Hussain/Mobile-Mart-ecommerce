package domain

import "time"

type Order struct {
	ID            uint `gorm:"primary key"`
	UserID        uint
	User          Users `gorm:"foreignkey:UserID;association_foreignkey:ID"`
	AddressID     uint
	Location      Address `gorm:"foreignkey:AddressID;association_foreignkey:ID"`
	PaymentMethod string
	TotalPrice    uint
	OrderDate     time.Time
	DeliveryDate  time.Time
	PaymentStatus status `gorm:"default:pending"`
	OrderStatus   string
	OrderID       string
}

type OrderProducts struct {
	OrderID     uint
	Orderid     Order `gorm:"foreignkey:OrderID;association_foreignkey:ID"`
	InventoryID uint
	Product     Inventories `gorm:"foreignkey:InventoryID;association_foreignkey:ID"`
	SellerID    uint        `gorm:"not null"`
	Seller      Seller      `gorm:"forgienKey:SellerID;association_foreignkey:ID"`
	Quantity    uint
	Price       uint
}
