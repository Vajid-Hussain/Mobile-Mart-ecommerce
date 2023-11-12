package responsemodel

type AdminLoginRes struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Result   string `json:"result,omitempty"`
	Token    string `json:"token,omitempty"`
	// RefreshToken string `json:"refreshtoken,omitempty"`
}

type UserDetails struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Email  string `json:"email,omitempty"`
	Phone  string `json:"phone,omitempty"`
	Status string `json:"status,omitempty"`
}
