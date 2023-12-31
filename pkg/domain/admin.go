package domain

type Admin struct {
	Name     string
	Email    string `gorm:"primary key"`
	Password string
}
