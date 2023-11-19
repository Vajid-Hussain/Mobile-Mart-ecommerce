package models

type Address struct {
	ID          uint
	Userid      uint
	FirstName   string `json:"firstName" validate:"required"`
	LastName    string `json:"lastName"`
	Street      string `json:"street" validate:"required,alpha"`
	City        string `json:"city" validate:"required,alpha"`
	State       string `json:"state" validate:"required,alpha"`
	Pincode     string `json:"pincode" validate:"min=6"`
	LandMark    string `json:"landmark" validate:"required"`
	PhoneNumber string `json:"phoneNumber" validate:"required,len=10,number"`
}

type EditAddress struct {
	ID          uint   `json:"id" validate:"required"`
	Userid      uint   `json:"userid" validate:"required"`
	FirstName   string `json:"firstName" validate:"required"`
	LastName    string `json:"lastName" validate:"required"`
	Street      string `json:"street" validate:"required,alpha"`
	City        string `json:"city" validate:"required,alpha"`
	State       string `json:"state" validate:"required,alpha"`
	Pincode     string `json:"pincode" validate:"required,min=6"`
	LandMark    string `json:"landmark" validate:"required"`
	PhoneNumber string `json:"phoneNumber" validate:"required,len=10,number"`
}
