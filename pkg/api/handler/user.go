package handler

import (
	"fmt"
	"net/http"
	"strings"

	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	resCustomError "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel/custom_error"
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

// @Summary		User Signup
// @Description	using this handler User can SIgnup
// @Tags			User
// @Accept			json
// @Produce			json
// @Param			user	body		requestmodel.UserDetails{}	true	"User Signup details"
// @Success		200	{object}	responsemodel.SignupData{}
// @Failure		400	{object}	response.Response{}
// @Router			/user/signup/ [post]
func (u *UserHandler) UserSignup(c *gin.Context) {
	var userSignupData requestmodel.UserDetails
	if err := c.BindJSON(&userSignupData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		resSignup, err := u.userUseCase.UserSignup(&userSignupData)
		if err != nil {
			response := response.Responses(http.StatusBadRequest, err.Error(), resSignup, nil)
			c.JSON(http.StatusUnauthorized, response)
		} else {
			response := response.Responses(http.StatusOK, "", resSignup, nil)
			c.JSON(http.StatusOK, response)
		}
	}
}

// @Summary		User Otp verification
// @Description	using this handler User can send otp
// @Tags			User
// @Accept			json
// @Produce			json
// @Param			user	body		requestmodel.OtpVerification{}	true	"User otp details"
// @Success		200	{object}	response.Response{}
// @Failure		400	{object}	response.Response{}
// @Router			/user/verifyOTP/ [post]
func (u *UserHandler) VerifyOTP(c *gin.Context) {

	var otpData requestmodel.OtpVerification
	token := c.Request.Header.Get("Authorization")

	if err := c.BindJSON(&otpData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result, err := u.userUseCase.VerifyOtp(otpData, token)
	if err != nil {
		response := response.Responses(http.StatusBadRequest, err.Error(), result, nil)
		c.JSON(http.StatusUnauthorized, response)
	} else {
		response := response.Responses(http.StatusOK, "Succesfully verified", result, nil)
		c.JSON(http.StatusOK, response)
	}
}

// @Summary		User Login
// @Description	using this handler User can Login
// @Tags			User
// @Accept			json
// @Produce			json
// @Param			user	body		requestmodel.UserLogin{}	true	"User Login details"
// @Success		200	{object}	response.Response{}
// @Failure		400	{object}	response.Response{}
// @Router			/user/login/ [post]
func (u *UserHandler) UserLogin(c *gin.Context) {
	var loginCredential requestmodel.UserLogin
	if err := c.BindJSON(&loginCredential); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result, err := u.userUseCase.UserLogin(loginCredential)
	if err != nil {
		response := response.Responses(http.StatusBadRequest, "", result, nil)
		c.JSON(http.StatusUnauthorized, response)
	} else {
		response := response.Responses(http.StatusOK, "Succesfully login", result, nil)
		c.JSON(http.StatusOK, response)
	}
}

// @Summary			All User
// @Description	 	using this handler admin can view user
// @Tags			Admins
// @Accept			json
// @Produce			json
// @Security        BearerTokenAuth
// @Param			id		path	string	true	"User ID in the URL path"
// @Success			200	{object}	response.Response{}
// @Failure			400	{object}	response.Response{}
// @Router			/admin/user/getuser [get]
func (u *UserHandler) GetUser(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "1")

	users, count, err := u.userUseCase.GetAllUsers(page, limit)
	if err != nil {
		// message := fmt.Sprintf("total users  %d", *count)
		// finalReslt := response.Responses(http.StatusNotFound, message, "", err.Error())
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		message := fmt.Sprintf("total users  %d", *count)
		finalReslt := response.Responses(http.StatusOK, message, users, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// @Summary		    Block User
// @Description	using this handler admin can block user
// @Tags			Admins
// @Accept			json
// @Produce			json
// @Security        BearerTokenAuth
// @Param			id		path	string	true	"User ID in the URL path"
// @Success		200	{object}	response.Response{}
// @Failure		400	{object}	response.Response{}
// @Router			/admin/user/block [patch]
func (u *UserHandler) BlockUser(c *gin.Context) {
	userID := c.Query("id")
	id := strings.TrimSpace(userID)

	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": resCustomError.IDParamsEmpty})
		return
	}

	err := u.userUseCase.BlcokUser(userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusNotFound, "", "", err.Error())
		c.JSON(http.StatusNotFound, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Succesfully block", "", nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// @Summary		    Unblock User
// @Description	using this handler admin Unblock user
// @Tags			Admins
// @Accept			json
// @Produce			json
// @Security        BearerTokenAuth
// @Param			id		path	string	true	"User ID in the URL path"
// @Success		200	{object}	response.Response{}
// @Failure		400	{object}	response.Response{}
// @Router			/admin/user/unblock [patch]
func (u *UserHandler) UnblockUser(c *gin.Context) {
	userID := c.Query("id")
	id := strings.TrimSpace(userID)

	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id parameter is empty"})
		return
	}

	err := u.userUseCase.UnblockUser(userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusNotFound, "", "", err.Error())
		c.JSON(http.StatusNotFound, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Succesfully unblock", "", nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}
