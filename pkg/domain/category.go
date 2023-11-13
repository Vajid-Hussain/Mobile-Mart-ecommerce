package domain

type Category struct {
	ID   uint   `gorm:"unique key"`
	Name string `gorm:"unique"`
}
