package models

type SellerProfile struct {
	ID              string `json:"id,omitempty"              validate:"required"`
	Name            string `json:"name,omitempty"            validate:"required"`
	Email           string `json:"email,omitempty"           validate:"email"`
	Password        string `json:"-"        validate:"min=4"`
	ConfirmPassword string `json:"-"                         validate:"eqfield=Password"`
	GST_NO          string `json:"gstno,omitempty"           validate:"len=15"`
	Description     string `json:"description,omitempty"     validate:"required"`
}

type SellerEditProfile struct {
	ID              string `json:"id,omitempty"              validate:"required"`
	Name            string `json:"name,omitempty"            validate:"required"`
	Email           string `json:"email,omitempty"           validate:"required,email"`
	Password        string `json:"password,omitempty"        validate:"required,min=4"`
	ConfirmPassword string `json:"confirmpassword,omitempty" validate:"required,eqfield=Password"`
	GST_NO          string `json:"gstno,omitempty"           validate:"required,len=15"`
	Description     string `json:"description,omitempty"     validate:"required"`
}
