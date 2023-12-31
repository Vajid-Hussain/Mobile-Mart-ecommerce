package domain

type Cart struct {
	UserID      uint
	User        Users `gorm:"foreignkey:UserID;association_foreignkey:ID"`
	InventoryID uint
	Product     Inventories `gorm:"foreignkey:InventoryID;association_foreignkey:ID"`
	Quantity    uint        `gorm:"default:1"`
	// Price       uint
	Status status `gorm:"default:active"`
}
