package requestmodel

type UserDetails struct {
	Name            string `json:"name"           validate:"nonzero"`
	Email           string `json:"email"`
	Phone           string `json:"phone"          validate:"len=10, numeric"`
	Password        string `json:"password"       validate:"min=4"`
	ConfirmPassword string `json:"confirmpassword"`
}
