package requestmodel

type SellerSignup struct{
	Name		    string `json:"name"            validate:"required"`
	Email           string `json:"email"`
	Password        string `json:"password"        validate:"min=4"`
	ConfirmPassword string `json:"confirmpassword" validate:"gte=4"`
	GST_NO			string `json:"gstno"           validate:"len=15"`
	Discription     string `json:"discription"     validate:"min=10"`
}