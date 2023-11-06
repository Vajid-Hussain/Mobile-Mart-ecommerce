package usecase

import (
	"fmt"
	"regexp"
	"strings"

	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
	interfaces "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/repository/interface"
	serviceInterface "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/service/interface"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/utils/helper"
	"gopkg.in/validator.v2"
)

type userUseCase struct {
	repo interfaces.IUserRepo
	otp  serviceInterface.Ijwt
}

func NewUserUseCase(userRepository interfaces.IUserRepo, otp serviceInterface.Ijwt) interfaceUseCase.IuserUseCase {
	return &userUseCase{repo: userRepository, otp: otp}
}

//useCases

func (u *userUseCase) UserSignup(userData *requestmodel.UserDetails) responsemodel.SignupData {

	var resSignUpFailed responsemodel.SignupData

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

	if err := validator.Validate(userData); err != nil {
		fmt.Println(err)

		for key := range err.(validator.ErrorMap) {
			switch key {
			case "Name":
				resSignUpFailed.Name = "Field is empty"
			case "Phone":
				resSignUpFailed.Phone = "must contain 10 numbers"
			case "Password":
				resSignUpFailed.Password = "password need more than 4 digit "
			}
		}

		isValid := ValidateEmailStructure(userData.Email)
		resSignUpFailed.Email = isValid
		if userData.ConfirmPassword != userData.Password {
			resSignUpFailed.ConfirmPassword = "ConfirmPassword is not correct , cross check"
		}
		return resSignUpFailed
	}

	isValid := ValidateEmailStructure(userData.Email)
	if isValid != "" || userData.ConfirmPassword != userData.Password {
		resSignUpFailed.Email = isValid
		if userData.ConfirmPassword != userData.Password {
			resSignUpFailed.ConfirmPassword = "ConfirmPassword is not correct , cross check"
		}
		return resSignUpFailed
	}

	if isExist := u.repo.IsUserExist(userData); isExist >= 1 {
		resSignUpFailed.IsUserExist = "User Exist ,change mail"
		return resSignUpFailed
	} else {
		userData.Id = helper.GenerateUUID()
		u.otp.TwilioSetup()
		_,err:=u.otp.SendOtp(userData.Phone)
		if err!=nil{
			resSignUpFailed.OTP="error of otp creation"
			return resSignUpFailed
		}
		u.repo.CreateUser(userData)
	}

	return resSignUpFailed
}

