package requestmodel

type UserDetails struct {
	Id				string
	Name            string `json:"name"           validate:"nonzero"`
	Email           string `json:"email"`
	Phone           string `json:"phone"          validate:"len=10"`
	Password        string `json:"password"       validate:"min=4"`
	ConfirmPassword string `json:"confirmpassword"`
}

//validate:"len=10, numeric"
type OtpVerification struct{
	Otp		string	`json:"otp"   validate:"len=6"`
}