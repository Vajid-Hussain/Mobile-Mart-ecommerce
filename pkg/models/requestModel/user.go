package requestmodel

type UserDetails struct {
	Id              uint
	Name            string `json:"name"           validate:"required"`
	Email           string `json:"email"          validate:"email"`
	Phone           string `json:"phone"          validate:"len=10"`
	Password        string `json:"password"       validate:"min=4"`
	ConfirmPassword string `json:"confirmpassword" validate:"eqfield=Password"`
}

type OtpVerification struct {
	Otp string `json:"otp"   validate:"len=6"`
}

type UserLogin struct {
	Phone    string `json:"phone"    validate:"len=10,number"`
	Password string `json:"password" validate:"required,min=4"`
}

type Address struct {
	UserID      uint
	ID          uint
	FirstName   string `json:"firstName" validate:"required"`
	LastName    string `json:"lastName"`
	Street      string `json:"street" validate:"required,alpha"`
	City        string `json:"city" validate:"required,alpha"`
	State       string `json:"state" validate:"required,alpha"`
	Pincode     uint   `json:"pincode" validate:"min=6"`
	LandMark    string `json:"landmark" validate:"required"`
	PhoneNumber string `json:"phoneNumber" validate:"required,len=10,number"`
}
