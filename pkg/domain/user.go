package domain

type UserDetails struct{
	ID			uint	`gorm:"unique;not null"`
	Name 		string
	Email		string
	Phone		string
	Password	string
	Blocked		bool	`gorm:"default:false"`
}