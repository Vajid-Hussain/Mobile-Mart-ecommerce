package usecase

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/config"
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
	resCustomError "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel/custom_error"
	interfaces "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/repository/interface"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/service"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/utils/helper"
	"github.com/go-playground/validator/v10"
	validaters "gopkg.in/validator.v2"
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
		// userData.Id = helper.GenerateUUID()
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
	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(loginCredential)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "Phone":
					resUserLogin.Phone = "phone number should be 10"
				case "Password":
					resUserLogin.Password = "Password atleast 4 digit"
				}
			}
		}
		return resUserLogin, errors.New("login credential not obey")
	}

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

	pageNO, err := strconv.Atoi(page)
	if err != nil {
		return nil, nil, resCustomError.ErrConversionOFPage
	}

	limits, err := strconv.Atoi(limit)
	if err != nil {
		return nil, nil, resCustomError.ErrConversionOfLimit
	}

	if pageNO < 1 {
		return nil, nil, resCustomError.ErrPagination
	}

	if limits <= 0 {
		return nil, nil, resCustomError.ErrPageLimit
	}

	offSet := (pageNO * limits) - limits
	limits = pageNO * limits

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
