package interfaceUseCase

import (
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
)

type IuserUseCase interface {
	UserSignup(*requestmodel.UserDetails) (*responsemodel.SignupData, error)
	VerifyOtp(requestmodel.OtpVerification, string) (responsemodel.OtpValidation, error)
	SendOtp(*requestmodel.SendOtp) (*string, error)
	UserLogin(requestmodel.UserLogin) (responsemodel.UserLogin, error)
	ForgotPassword(*requestmodel.ForgotPassword, string) error

	GetAllUsers(string, string) (*[]responsemodel.UserDetails, *int, error)
	BlcokUser(string) error
	UnblockUser(string) error

	AddAddress(*requestmodel.Address) (*requestmodel.Address, error)
	GetAddress(string, string, string) (*[]requestmodel.Address, error)
	EditAddress(*requestmodel.EditAddress) (*requestmodel.EditAddress, error)
	DeleteAddress(string, string) error

	GetProfile(string) (*requestmodel.UserDetails, error)
	UpdateProfile(*requestmodel.UserEditProfile) (*requestmodel.UserDetails, error)
}
