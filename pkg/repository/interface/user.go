package interfaces

import requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"

type IUserRepo interface {
	CreateUser(*requestmodel.UserDetails)
	IsUserExist(string) int
	ChangeUserStatusActive(string) error
	FetchUserID(string) (string, error)
	FetchPasswordUsingPhone(string) (string, error)
}
