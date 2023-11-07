package domain

type UserDetails struct{
	ID			string	`gorm:"unique;not null; primary key"`
	Name 		string
	Email		string	`gorm:"not null"`
	Phone		string
	Password	string
	Blocked		bool	`gorm:"default:false"`
	Delete		bool	`gorm:"default:false"`
}