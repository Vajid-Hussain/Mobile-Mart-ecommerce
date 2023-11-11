package responsemodel

type AdminLoginRes struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Result   string `json:"result,omitempty"`
	Token    string `json:"token,omitempty"`
	// RefreshToken string `json:"refreshtoken,omitempty"`
}
