package handler

import (
	"fmt"
	"net/http"

	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel/response"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase interfaceUseCase.IuserUseCase
}

func NewUserHandler(userUseCase interfaceUseCase.IuserUseCase) *UserHandler {
	return &UserHandler{userUseCase: userUseCase}
}

//handlers

func (u *UserHandler) UserSignup(c *gin.Context) {
	var userSignupData requestmodel.UserDetails
	if err := c.BindJSON(&userSignupData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		resSignup := u.userUseCase.UserSignup(&userSignupData)
		c.JSON(http.StatusOK, resSignup)
	}
}

func (u *UserHandler) VerifyOTP(c *gin.Context) {
	var otpData requestmodel.OtpVerification
	if err:=c.BindJSON(&otpData); err!=nil{
		fmt.Println(err)
	}
	
	result, err := u.userUseCase.VerifyOtp(otpData)
	if err == "" {
		response := response.Responses(http.StatusUnauthorized, "Invalid credentials", result, nil)
		c.JSON(http.StatusUnauthorized, response)
	}else{
		response := response.Responses(http.StatusUnauthorized, "Succesfully verified", result, nil)
		c.JSON(http.StatusUnauthorized, response)
	}
}
