package usecase

import (
	"errors"
	"fmt"

	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/config"
	models "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/model"
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
	interfaces "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/repository/interface"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/service"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/utils/helper"
	"github.com/go-playground/validator/v10"
)

type userUseCase struct {
	repo  interfaces.IUserRepo
	token config.Token
}

func NewUserUseCase(userRepository interfaces.IUserRepo, token *config.Token) interfaceUseCase.IuserUseCase {
	return &userUseCase{repo: userRepository,
		token: *token,
	}
}

//useCases

func (u *userUseCase) UserSignup(userData *requestmodel.UserDetails) (responsemodel.SignupData, error) {

	var resSignup responsemodel.SignupData

	if isExist := u.repo.IsUserExist(userData.Phone); isExist >= 1 {
		resSignup.IsUserExist = "User Exist ,change phone number"
		return resSignup, errors.New("user is exist try again , with another phone number")
	} else {
		service.TwilioSetup()
		_, err := service.SendOtp(userData.Phone)
		if err != nil {
			resSignup.OTP = "error of otp creation"
			return resSignup, errors.New("error at attempt for creating new OTP")
		}

		HashedPassword := helper.HashPassword(userData.Password)
		userData.Password = HashedPassword

		u.repo.CreateUser(userData)

		token, err := service.TemperveryTokenForOtpVerification(u.token.TemperveryKey, userData.Phone)
		if err != nil {
			resSignup.Token = "error of create a token"
			return resSignup, errors.New("creating token is error")
		}
		resSignup.Token = token
	}

	return resSignup, nil
}

func (u *userUseCase) VerifyOtp(otpConstrain requestmodel.OtpVerification, token string) (responsemodel.OtpValidation, error) {

	var otpResponse responsemodel.OtpValidation

	phone, err := service.FetchPhoneFromToken(token, u.token.TemperveryKey)
	if err != nil {
		otpResponse.Token = "invalid token"
		return otpResponse, errors.New("error ar token extraction , cause by invalid token")
	}

	service.TwilioSetup()

	if err := service.VerifyOtp(phone, otpConstrain.Otp); err != nil {
		otpResponse.Otp = "otp verification failed"
		return otpResponse, errors.New("otp is not match with your number, try egain")
	}

	if err := u.repo.ChangeUserStatusActive(phone); err != nil {
		otpResponse.Phone = "no user exist with phone number , verify is phone number is it correct "
		return otpResponse, errors.New("no user exist ")
	}

	userID, err := u.repo.FetchUserID(phone)
	if err != nil {
		otpResponse.Result = "error at db attempt to featch userID"
		return otpResponse, errors.New("error cause by fetching user id")
	}

	accessToken, err := service.GenerateAcessToken(u.token.UserSecurityKey, userID)
	if err != nil {
		otpResponse.Result = "creating token not done succesfully"
		return otpResponse, errors.New("token creation cause error")
	}

	refreshToken, err := service.GenerateRefreshToken(u.token.UserSecurityKey)
	if err != nil {
		otpResponse.Result = "creating token not done succesfully"
		return otpResponse, errors.New("token creation cause error")
	}

	otpResponse.AccessToken = accessToken
	otpResponse.RefreshToken = refreshToken
	otpResponse.Result = "success"

	return otpResponse, nil
}

func (u *userUseCase) UserLogin(loginCredential requestmodel.UserLogin) (responsemodel.UserLogin, error) {
	var resUserLogin responsemodel.UserLogin

	password, err := u.repo.FetchPasswordUsingPhone(loginCredential.Phone)
	if err != nil {
		resUserLogin.Error = err.Error()
		return resUserLogin, err
	}

	err = helper.CompairPassword(password, loginCredential.Password)
	if err != nil {
		resUserLogin.Error = err.Error()
		return resUserLogin, err
	}

	userID, err := u.repo.FetchUserID(loginCredential.Phone)
	if err != nil {
		resUserLogin.Error = err.Error()
		return resUserLogin, err
	}

	accessToken, err := service.GenerateAcessToken(u.token.UserSecurityKey, userID)
	if err != nil {
		resUserLogin.Error = err.Error()
		return resUserLogin, err
	}

	refreshToken, err := service.GenerateRefreshToken(u.token.UserSecurityKey)
	if err != nil {
		resUserLogin.Error = err.Error()
	}

	resUserLogin.AccessToken = accessToken
	resUserLogin.RefreshToken = refreshToken

	return resUserLogin, nil
}

func (r *userUseCase) GetAllUsers(page string, limit string) (*[]responsemodel.UserDetails, *int, error) {

	ch := make(chan int)

	go r.repo.UserCount(ch)
	count := <-ch

	offSet, limits, err := helper.Pagination(page, limit)
	if err != nil {
		return nil, &count, err
	}

	userDetails, err := r.repo.AllUsers(offSet, limits)
	if err != nil {
		return nil, nil, err
	}

	return userDetails, &count, nil
}

func (r *userUseCase) BlcokUser(id string) error {
	err := r.repo.BlockUser(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *userUseCase) UnblockUser(id string) error {
	err := r.repo.UnblockUser(id)
	if err != nil {
		return err
	}
	return nil
}

// ------------------------------------------user Address------------------------------------\\

func (r *userUseCase) AddAddress(address *models.Address) (*models.Address, error) {

	add, err := r.repo.CreateAddress(address)
	if err != nil {
		return nil, err
	}
	return add, nil
}

func (r *userUseCase) GetAddress(userID string, page string, limit string) (*[]models.Address, error) {

	offset, limits, err := helper.Pagination(page, limit)
	if err != nil {
		return nil, err
	}

	address, err := r.repo.GetAddress(userID, offset, limits)
	if err != nil {
		return nil, err
	}
	return address, nil
}

func (r *userUseCase) EditAddress(address *models.EditAddress) (*models.EditAddress, error) {

	add, err := r.repo.GetAAddress(address.ID)
	if err != nil {
		return nil, err
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(address)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				fieldName := e.Field()
				switch fieldName {
				case "ID":
					address.ID = add.ID
				case "Userid":
					address.Userid = add.Userid
				case "FirstName":
					address.FirstName = add.FirstName
				case "LastName":
					address.LastName = add.LastName
				case "Street":
					address.Street = add.Street
				case "City":
					address.City = add.City
				case "State":
					address.State = add.State
				case "Pincode":
					address.Pincode = add.Pincode
				case "LandMark":
					address.LandMark = add.LandMark
				case "PhoneNumber":
					address.PhoneNumber = add.PhoneNumber
				}
			}
		}

	}

	editedAddress, err := r.repo.UpdateAddress(address)
	if err != nil {
		return nil, err
	}
	return editedAddress, nil
}

func (r *userUseCase) DeleteAddress(addressID string, userID string) error {
	err := r.repo.DeleteAddress(addressID, userID)
	if err != nil {
		return err
	}
	return nil
}

// ------------------------------------------user Profile------------------------------------\\

func (r *userUseCase) GetProfile(userID string) (*models.UserDetails, error) {
	userDetails, err := r.repo.GetProfile(userID)
	if err != nil {
		return nil, err
	}
	return userDetails, nil
}

func (r *userUseCase) UpdateProfile(editedProfile *models.UserEditProfile) (*models.UserDetails, error) {

	if editedProfile.Password != editedProfile.ConfirmPassword {
		return nil, errors.New("password and confirmpassword is not match")
	}

	if editedProfile.Password != "" {
		editedProfile.Password = helper.HashPassword(editedProfile.Password)
	}

	userProfile, err := r.repo.GetProfile(editedProfile.Id)
	if err != nil {
		return nil, err
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(editedProfile)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				fieldName := e.Field()
				switch fieldName {
				case "Id":
					editedProfile.Id = userProfile.Id
				case "Name":
					editedProfile.Name = userProfile.Name
				case "Email":
					editedProfile.Email = userProfile.Email
				case "Password":
					editedProfile.Password = userProfile.Password
				}
			}
		}

	}

	userProfile, err = r.repo.UpdateProfile((*models.UserDetails)(editedProfile))
	if err != nil {
		return nil, err
	}
	fmt.Println("-----", editedProfile)

	return userProfile, nil

}
