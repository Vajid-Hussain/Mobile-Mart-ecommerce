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
	UpdatePassword(string, string) error

	AllUsers(int, int) (*[]responsemodel.UserDetails, error)
	UserCount(chan int)
	BlockUser(string) error
	UnblockUser(string) error

	CreateAddress(*requestmodel.Address) (*requestmodel.Address, error)
	GetAddress(string, int, int) (*[]requestmodel.Address, error)
	UpdateAddress(*requestmodel.EditAddress) (*requestmodel.EditAddress, error)
	GetAAddress(string) (*requestmodel.Address, error)
	DeleteAddress(string, string) error

	GetProfile(string) (*requestmodel.UserDetails, error)
	UpdateProfile(*requestmodel.UserDetails) (*requestmodel.UserDetails, error)
}
