package interfaceUseCase

import (
	models "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/model"
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
)

type IuserUseCase interface {
	UserSignup(*requestmodel.UserDetails) (responsemodel.SignupData, error)
	VerifyOtp(requestmodel.OtpVerification, string) (responsemodel.OtpValidation, error)
	UserLogin(requestmodel.UserLogin) (responsemodel.UserLogin, error)
	GetAllUsers(string, string) (*[]responsemodel.UserDetails, *int, error)
	BlcokUser(string) error
	UnblockUser(string) error

	AddAddress(*models.Address) (*models.Address, error)
	GetAddress(string, string, string) (*[]models.Address, error)
	EditAddress(*models.EditAddress) (*models.EditAddress, error)
	DeleteAddress(string, string) error
}
