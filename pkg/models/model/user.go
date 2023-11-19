package models

type UserDetails struct {
	Id              string `json:"userid"`
	Name            string `json:"name"           validate:"required"`
	Email           string `json:"email"          validate:"email"`
	Phone           string `json:"phone,omitempty"          validate:"len=10"`
	Password        string `json:"-"       validate:"min=4"`
	ConfirmPassword string `json:"confirmpassword,omitempty" validate:"eqfield=Password"`
}

type UserEditProfile struct {
	Id              string `json:"userid,omitempty" validate:"required"`
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Phone           string `json:"phone,omitempty" validate:"required,len=10"`
	Password        string `json:"password,omitempty" validate:"required,min=4"`
	ConfirmPassword string `json:"confirmpassword,omitempty" validate:"required,eqfield=Password"`
}

type Address struct {
	ID          string
	Userid      string
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
	ID          string `json:"id" validate:"required"`
	Userid      string `json:"userid" validate:"required"`
	FirstName   string `json:"firstName" validate:"required"`
	LastName    string `json:"lastName" validate:"required"`
	Street      string `json:"street" validate:"required,alpha"`
	City        string `json:"city" validate:"required,alpha"`
	State       string `json:"state" validate:"required,alpha"`
	Pincode     string `json:"pincode" validate:"required,min=6"`
	LandMark    string `json:"landmark" validate:"required"`
	PhoneNumber string `json:"phoneNumber" validate:"required,len=10,number"`
}