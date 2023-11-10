package requestmodel

type SellerSignup struct {
	ID              string
	Name            string `json:"name"            validate:"required"`
	Email           string `json:"email"`
	Password        string `json:"password"        validate:"min=4"`
	ConfirmPassword string `json:"confirmpassword"`
	GST_NO          string `json:"gstno"           validate:"len=15"`
	Description     string `json:"description"     validate:"required"`
}

type SellerLogin struct {
	Email    string `json:"email"    validate:"email"`
	Password string `json:"password" validate:"min=4"`
}
