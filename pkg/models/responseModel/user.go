package responsemodel

type SignupData struct{
	Name            string `json:"name,omitempty"`
	Email           string `json:"email,omitempty"`
	Phone           string `json:"phone,omitempty"`
	Password        string `json:"password,omitempty"`
	IsUserExist		string `json:"isUserExist,omitempty"`
}