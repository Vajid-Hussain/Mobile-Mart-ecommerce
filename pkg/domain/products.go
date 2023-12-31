package domain

type Inventories struct {
	ID                 uint `gorm:"primary key"`
	Productname        string
	Description        string
	BrandID            uint     `gorm:"not null"`
	Brand              Brand    `gorm:"forgienkey:BrandID;association_foreignkey:ID"`
	CategoryID         uint     `gorm:"not null"`
	Category           Category `gorm:"forgienKey:CategoryID;association_foreignkey:ID"`
	SellerID           string   `gorm:"not null"`
	Seller             Seller   `gorm:"forgienKey:SellerID;association_foreignkey:ID"`
	Mrp                int
	Discount           uint
	Saleprice          int
	Units              int64
	Os                 string
	CellularTechnology string
	Ram                int
	Screensize         float64
	Batterycapacity    int
	Processor          string
	ImageURL           string
	Status             status `gorm:"default:active"`
}
