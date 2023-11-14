package domain

type Seller struct {
	ID          uint `gorm:"primary key"`
	Name        string
	Email       string `gorm:"unique"`
	Password    string
	GST_NO      string `gorm:"not null"`
	Description string
	Status      status_user `gorm:"default:pending"`
}
