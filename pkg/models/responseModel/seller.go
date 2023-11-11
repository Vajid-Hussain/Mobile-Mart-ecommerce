package responsemodel

type SellerSignupRes struct {
	Name            string `json:"name,omitempty"`
	Email           string `json:"email,omitempty"`
	Password        string `json:"password,omitempty"`
	ConfirmPassword string `json:"confirmPassword,omitempty"`
	GST_NO          string `json:"gstno,omitempty"`
	Description     string `json:"description,omitempty"`
	SellerExist     string `json:"sellerExist,omitempty"`
	Result          string `json:"result,omitempty"`
}

type SellerLoginRes struct {
	Email        string `json:"email,omitempty"`
	Password     string `json:"password,omitempty"`
	Result       string `json:"result,omitempty"`
	AccessToken  string `json:"accesstoken,omitempty"`
	RefreshToken string `json:"refreshtoken,omitempty"`
}
