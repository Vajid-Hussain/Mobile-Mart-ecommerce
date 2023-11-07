package usecase

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/config"
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
	interfaces "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/repository/interface"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/service"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/utils/helper"
	"github.com/go-playground/validator/v10"
	validaters "gopkg.in/validator.v2"
)

type userUseCase struct {
	repo interfaces.IUserRepo
	token config.Token
}

func NewUserUseCase(userRepository interfaces.IUserRepo, token *config.Token) interfaceUseCase.IuserUseCase {
	return &userUseCase{repo: userRepository, 
		token: *token,
	}
}

//useCases

func (u *userUseCase) UserSignup(userData *requestmodel.UserDetails) (responsemodel.SignupData ,error){

	var resSignup responsemodel.SignupData

	ValidateEmailStructure := func(email string) string {
		pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

		match, _ := regexp.MatchString(pattern, email)

		domain := strings.Split(email, "@")
		if len(domain) == 2 && domain[1] == "gmail.com" && match {
			return ""
		} else {
			return "Email is wrong"
		}
	}

	if err := validaters.Validate(userData); err != nil {

		for key := range err.(validaters.ErrorMap) {
			switch key {
			case "Name":
				resSignup.Name = "Field is empty"
			case "Phone":
				resSignup.Phone = "must contain 10 numbers"
			case "Password":
				resSignup.Password = "password need more than 4 digit "
			}
		}

		isValid := ValidateEmailStructure(userData.Email)
		resSignup.Email = isValid
		if userData.ConfirmPassword != userData.Password {
			resSignup.ConfirmPassword = "ConfirmPassword is not correct , cross check"
		}
		return resSignup, errors.New("not satisfying user credentials")
	}

	isValid := ValidateEmailStructure(userData.Email)
	if isValid != "" || userData.ConfirmPassword != userData.Password {
		resSignup.Email = isValid
		if userData.ConfirmPassword != userData.Password {
			resSignup.ConfirmPassword = "ConfirmPassword is not match , cross check"
		}
		return resSignup, errors.New("not satisfying user credentials")
	}

	if isExist := u.repo.IsUserExist(userData.Phone); isExist >= 1 {
		resSignup.IsUserExist = "User Exist ,change phone number"
		return resSignup, errors.New("user is exist try again , with another phone number")
	} else {
		userData.Id = helper.GenerateUUID()
			service.TwilioSetup()
		_, err := service.SendOtp(userData.Phone)
		if err != nil {
			resSignup.OTP = "error of otp creation"
			return resSignup, errors.New("error at attempt for creating new OTP")
		}

		HashedPassword:=helper.HashPassword(userData.Password)
		userData.Password=HashedPassword
		u.repo.CreateUser(userData)
		token, err:=service.TemperveryTokenStorePhone(u.token.TemperveryKey, userData.Phone)
		if err!=nil{
			resSignup.Token="error of create a token" 
			return resSignup, errors.New("creating token is error")
		}
		resSignup.Token=token
	}

	return resSignup , nil 
}

func (u *userUseCase) VerifyOtp(otpConstrain requestmodel.OtpVerification, token string) (responsemodel.OtpValidation, error) {

	var otpResponse responsemodel.OtpValidation
	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(otpConstrain)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "otp":
					otpResponse.Otp = "otp should 6 numbers"
				}
			}
		}
		return otpResponse, nil
	}
	phone,err:=service.FetchPhoneFromToken(token, u.token.TemperveryKey)
	if err!=nil{
		otpResponse.Token="invalid token"
		return otpResponse, errors.New("error ar token extraction , cause by invalid token")
	}

	service.TwilioSetup()

	if err := service.VerifyOtp(phone, otpConstrain.Otp); err != nil {
		otpResponse.Otp = "otp verification failed"
		return otpResponse, errors.New("otp is not match with your number, try egain")
	}

	if err := u.repo.CheckUserByPhone(phone); err != nil {
		otpResponse.Phone = "no user exist with phone number , verify is phone number is it correct "
		return otpResponse, errors.New("no user exist ")
	}

	fmt.Println("00000")
	otpResponse.Result= "success"
	return otpResponse, nil
}
