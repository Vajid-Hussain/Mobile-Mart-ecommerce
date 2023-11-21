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

type SellerDetails struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Email       string `json:"email,omitempty"`
	GST_NO      string `json:"gstno,omitempty"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status,omitempty"`
}

type SellerProfile struct {
	ID              string `json:"id,omitempty"              validate:"required"`
	Name            string `json:"name,omitempty"            validate:"required"`
	Email           string `json:"email,omitempty"           validate:"email"`
	Password        string `json:"-"        validate:"min=4"`
	ConfirmPassword string `json:"-"                         validate:"eqfield=Password"`
	GST_NO          string `json:"gstno,omitempty"           validate:"len=15"`
	Description     string `json:"description,omitempty"     validate:"required"`
}
