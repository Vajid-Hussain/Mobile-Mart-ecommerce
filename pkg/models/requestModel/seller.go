package requestmodel

type SellerSignup struct {
	ID              uint   `swaggerignore:"true"`
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

type SellerEditProfile struct {
	ID              string `json:"id,omitempty" swaggerignore:"true"              validate:"required"`
	Name            string `json:"name,omitempty"            validate:"required"`
	Email           string `json:"email,omitempty"           validate:"required,email"`
	Password        string `json:"password,omitempty"        validate:"required,min=4"`
	ConfirmPassword string `json:"confirmpassword,omitempty" validate:"required,eqfield=Password"`
	// GST_NO          string `json:"gstno,omitempty"           validate:"required,len=15"`
	Description string `json:"description,omitempty"     validate:"required"`
}
