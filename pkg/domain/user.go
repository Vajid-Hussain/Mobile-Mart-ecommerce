package domain

type status_user string

const (
	Active  status_user = "active"
	Block   status_user = "block"
	Delete  status_user = "delete"
	Pending status_user = "pending"
)

type Users struct {
	ID       uint `gorm:"unique;not null; primary key"`
	Name     string
	Email    string
	Phone    string
	Password string
	Status   status_user `gorm:"default:pending"`
}
