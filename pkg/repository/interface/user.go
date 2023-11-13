package interfaces

import (
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
)

type IUserRepo interface {
	CreateUser(*requestmodel.UserDetails)
	IsUserExist(string) int
	ChangeUserStatusActive(string) error
	FetchUserID(string) (string, error)
	FetchPasswordUsingPhone(string) (string, error)
	AllUsers(int, int) (*[]responsemodel.UserDetails, error)
	UserCount(chan int)
	BlockUser(string) error
	UnblockUser(string) error
}
