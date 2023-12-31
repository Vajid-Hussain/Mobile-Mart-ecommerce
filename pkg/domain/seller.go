package domain

type Seller struct {
	ID           uint `gorm:"primary key"`
	Name         string
	Email        string `gorm:"unique"`
	Password     string
	GST_NO       string `gorm:"not null"`
	Description  string
	SellerCredit uint   `gorm:"default:0"`
	Status       status `gorm:"default:pending"`
}
