package requestmodel

type SellerSignup struct {
	ID              uint
	Name            string `json:"name"            validate:"required"`
	Email           string `json:"email"           validate:"email"`
	Password        string `json:"password"        validate:"min=4"`
	ConfirmPassword string `json:"confirmpassword" validate:"eqfield=Password"`
	GST_NO          string `json:"gstno"           validate:"len=15"`
	Description     string `json:"description"     validate:"required"`
}

type SellerLogin struct {
	Email    string `json:"email"    validate:"email"`
	Password string `json:"password" validate:"min=4"`
}
