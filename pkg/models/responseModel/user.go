package responsemodel

type SignupData struct {
	ID            string `json:"userID"`
	Name          string `json:"name,omitempty"`
	Email         string `json:"email,omitempty"`
	Phone         string `json:"phone,omitempty"`
	OTP           string `json:"otp,omitempty"`
	Token         string `json:"token,omitempty"`
	IsUserExist   string `json:"isUserExist,omitempty"`
	ReferalCode   string `json:"referalCode"`
	WalletBelance uint   `json:"walletBelance"`
}

type OtpValidation struct {
	Phone        string `json:"phone,omitempty"`
	Otp          string `json:"otp,omitempty"`
	Result       string `json:"result,omitempty"`
	Token        string `json:"token,omitempty"`
	AccessToken  string `json:"accesstoken,omitempty"`
	RefreshToken string `json:"refreshtoken,omitempty"`
}

type UserLogin struct {
	Phone        string `json:"phone,omitempty"`
	Password     string `json:"password,omitempty"`
	AccessToken  string `json:"accesstoken,omitempty"`
	RefreshToken string `json:"refreshtoken,omitempty"`
	Error        string `json:"error,omitempty"`
}

type TokenVerificationMiddlewire struct {
	Error string `json:"error"`
}

type UserDetails struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Email  string `json:"email,omitempty"`
	Phone  string `json:"phone,omitempty"`
	Status string `json:"status,omitempty"`
}
