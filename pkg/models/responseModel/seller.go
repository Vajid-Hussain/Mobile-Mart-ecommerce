package responsemodel

type SellerSignupRes struct{
	Name            string `json:"name,omitempty"`
	Email           string `json:"email,omitempty"`
	Password        string `json:"password,omitempty"`
	ConfirmPassword string `json:"confirmPassword,omitempty"`
	GST_NO			string `json:"gstno,omitempty"`
	IsVenderExist   string `json:"isUserExist,omitempty"`
}