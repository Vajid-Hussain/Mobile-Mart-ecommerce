package domain

type Category struct {
	ID       int `gorm:"unique key"`
	Category string
}
