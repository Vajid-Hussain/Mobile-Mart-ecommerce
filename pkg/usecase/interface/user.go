package interfaceUseCase

import (
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
}
