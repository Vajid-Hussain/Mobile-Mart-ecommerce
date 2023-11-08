package domain


type Seller struct {
	ID              string `gorm:"primary key"`
	Name            string
	Email           string `gorm:"unique"`
	Password        string
	ConfirmPassword string
	GST_NO          string	`gorm:"not null"`
	Description     string
	Status          status_user `gorm:"default:pending"`
}
