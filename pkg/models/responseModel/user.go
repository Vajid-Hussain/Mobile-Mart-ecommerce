package responsemodel


type SignupData struct {
	Name            string `json:"name,omitempty"`
	Email           string `json:"email,omitempty"`
	Phone           string `json:"phone,omitempty"`
	Password        string `json:"password,omitempty"`
	OTP             string `json:"otp,omitempty"`
	Token           string `json:"token,omitempty"`
	ConfirmPassword string `json:"confirmPassword,omitempty"`
	IsUserExist     string `json:"isUserExist,omitempty"`
}

type OtpValidation struct {
	Phone string `json:"phone,omitempty"`
	Otp   string `json:"otp,omitempty"`
	Result string`json:"result,omitempty"`
	Token string `json:"token,omitempty"`
}

type UserLogin struct{
	Phone 		string	`json:"phone,omitempty"`
	Password 	string	`json:"password,omitempty"`
	Token		string	`json:"token,omitempty"`
	Error		string	`json:"error,omitempty"`
}

type TokenVerificationMiddlewire struct{
	Error		string	`json:"error"`
}