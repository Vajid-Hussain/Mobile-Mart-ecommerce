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
