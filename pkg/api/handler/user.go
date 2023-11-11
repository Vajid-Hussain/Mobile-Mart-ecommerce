package handler

import (
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
		resSignup, err := u.userUseCase.UserSignup(&userSignupData)
		if err != nil {
			response := response.Responses(http.StatusUnauthorized, err.Error(), resSignup, nil)
			c.JSON(http.StatusUnauthorized, response)
		} else {
			response := response.Responses(http.StatusOK, "", resSignup, nil)
			c.JSON(http.StatusOK, response)
		}
	}
}

func (u *UserHandler) VerifyOTP(c *gin.Context) {

	var otpData requestmodel.OtpVerification
	token := c.Request.Header.Get("Authorization")

	if err := c.BindJSON(&otpData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result, err := u.userUseCase.VerifyOtp(otpData, token)
	if err != nil {
		response := response.Responses(http.StatusUnauthorized, err.Error(), result, nil)
		c.JSON(http.StatusUnauthorized, response)
	} else {
		response := response.Responses(http.StatusOK, "Succesfully verified", result, nil)
		c.JSON(http.StatusOK, response)
	}
}

func (u *UserHandler) UserLogin(c *gin.Context) {
	var loginCredential requestmodel.UserLogin
	if err := c.BindJSON(&loginCredential); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result, err := u.userUseCase.UserLogin(loginCredential)
	if err != nil {
		response := response.Responses(http.StatusUnauthorized, "", result, nil)
		c.JSON(http.StatusUnauthorized, response)
	} else {
		response := response.Responses(http.StatusOK, "Succesfully login", result, nil)
		c.JSON(http.StatusOK, response)
	}
}
