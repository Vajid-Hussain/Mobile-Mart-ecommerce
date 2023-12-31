package requestmodel

type AdminLoginData struct {
	Email    string `json:"email"    validate:"email"`
	Password string `json:"password" validate:"min=4"`
}
