package interfaces

import (
	models "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/model"
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

	CreateAddress(*models.Address) (*models.Address, error)
	GetAddress(string, int, int) (*[]models.Address, error)
	UpdateAddress(*models.EditAddress) (*models.EditAddress, error)
	GetAAddress(string) (*models.Address, error)
	DeleteAddress(string, string) error

	GetProfile(string) (*models.UserDetails, error)
	UpdateProfile(*models.UserDetails) (*models.UserDetails, error)
}
