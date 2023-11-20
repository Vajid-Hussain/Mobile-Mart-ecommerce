package domain

type status string

const (
	Active  status = "active"
	Block   status = "block"
	Delete  status = "delete"
	Pending status = "pending"
)

type Users struct {
	ID       uint `gorm:"unique;not null; primary key"`
	Name     string
	Email    string
	Phone    string
	Password string
	Status   status `gorm:"default:pending"`
}

type Address struct {
	ID          uint `gorm:"unique;not null;primaryKey"`
	Userid      uint
	User        Users `gorm:"foreignkey:Userid;association_foreignkey:ID"`
	FirstName   string
	LastName    string
	Street      string
	City        string
	State       string
	Pincode     string
	LandMark    string
	PhoneNumber string
	Status      status `gorm:"default:active"`
}

// type Cart struct {
// 	ID
// 	Inventories
// 	Quantity uint
// 	Price    uint
// }
