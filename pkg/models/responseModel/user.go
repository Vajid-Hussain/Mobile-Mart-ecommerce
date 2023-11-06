package responsemodel

type SignupData struct{
	Name            string `json:"name,omitempty"`
	Email           string `json:"email,omitempty"`
	Phone           string `json:"phone,omitempty"`
	Password        string `json:"password,omitempty"`
	OTP 			string `json:"otp,omitempty"`
	ConfirmPassword	string `json:"confirmPassword,omitempty"`
	IsUserExist		string `json:"isUserExist,omitempty"`
}