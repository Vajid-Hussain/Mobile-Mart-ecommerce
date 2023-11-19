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
