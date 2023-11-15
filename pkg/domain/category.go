package domain

type Category struct {
	ID     uint   `gorm:"unique key"`
	Name   string `gorm:"unique"`
	Status status `gorm:"default:active"`
}

type Brand struct {
	ID     uint   `gorm:"unique key"`
	Name   string `gorm:"unique"`
	Status status `gorm:"default:active"`
}
