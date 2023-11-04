package requestmodel


type UserDetails struct{
	Name 			string	`json:"name"`	
	Email			string	`json:"email"`
	Phone			string	`json:"phone"`
	Password		string	`json:"password"`
	ConfirmPassword string	`json:"confirmpassword" `

}