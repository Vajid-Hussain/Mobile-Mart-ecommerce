package handler

import (
	"fmt"
	"net/http"
	"strings"

	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	resCustomError "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel/custom_error"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel/response"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/utils/helper"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase interfaceUseCase.IuserUseCase
}

func NewUserHandler(userUseCase interfaceUseCase.IuserUseCase) *UserHandler {
	return &UserHandler{userUseCase: userUseCase}
}

// @Summary		User Signup
// @Description	using this handler User can SIgnup
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			user	body		requestmodel.UserDetails{}	true	"User Signup details"
// @Success		200		{object}	responsemodel.SignupData{}
// @Failure		400		{object}	response.Response{}
// @Router			/signup/ [post]
func (u *UserHandler) UserSignup(c *gin.Context) {

	var userSignupData requestmodel.UserDetails

	if err := c.BindJSON(&userSignupData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := helper.Validation(userSignupData)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", data, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	resSignup, err := u.userUseCase.UserSignup(&userSignupData)
	if err != nil {
		response := response.Responses(http.StatusBadRequest, "", resSignup, err.Error())
		c.JSON(http.StatusBadRequest, response)
	} else {
		response := response.Responses(http.StatusOK, "", resSignup, nil)
		c.JSON(http.StatusOK, response)
	}
}

// @Summary		User Otp verification
// @Description	using this handler User can send otp
// @Tags			User
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Param			user	body		requestmodel.OtpVerification{}	true	"User otp details"
// @Success		200		{object}	response.Response{}
// @Failure		400		{object}	response.Response{}
// @Router			/verifyOTP/ [post]
func (u *UserHandler) VerifyOTP(c *gin.Context) {

	var otpData requestmodel.OtpVerification

	token := c.Request.Header.Get("Authorization")

	if err := c.BindJSON(&otpData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	data, err := helper.Validation(otpData)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", data, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	result, err := u.userUseCase.VerifyOtp(otpData, token)
	if err != nil {
		response := response.Responses(http.StatusBadRequest, "", result, err.Error())
		c.JSON(http.StatusBadRequest, response)
	} else {
		response := response.Responses(http.StatusOK, "Succesfully verified", result, nil)
		c.JSON(http.StatusOK, response)
	}
}

// @Summary		Send OTP
// @Description	Send OTP (One-Time Password) for verification.
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			otp	body		requestmodel.SendOtp	true	"OTP details for sending"
// @Success		200	{object}	response.Response		"OTP sent successfully"
// @Failure		400	{object}	response.Response		"Bad request"
// @Router			/sendotp [post]
func (u *UserHandler) SendOtp(c *gin.Context) {

	var sendOtp requestmodel.SendOtp

	if err := c.BindJSON(&sendOtp); err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, resCustomError.BindingConflict, nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	data, err := helper.Validation(sendOtp)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", data, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	tempToken, err := u.userUseCase.SendOtp(&sendOtp)
	if err != nil {
		response := response.Responses(http.StatusBadRequest, "", "", err.Error())
		c.JSON(http.StatusBadRequest, response)
	} else {
		response := response.Responses(http.StatusOK, "Succesfully otp send", tempToken, nil)
		c.JSON(http.StatusOK, response)
	}
}

// @Summary		User Login
// @Description	using this handler User can Login
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			user	body		requestmodel.UserLogin{}	true	"User Login details"
// @Success		200		{object}	response.Response{}
// @Failure		400		{object}	response.Response{}
// @Router			/login/ [post]
func (u *UserHandler) UserLogin(c *gin.Context) {
	var loginCredential requestmodel.UserLogin
	if err := c.BindJSON(&loginCredential); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	data, err := helper.Validation(loginCredential)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", data, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	result, err := u.userUseCase.UserLogin(loginCredential)
	if err != nil {
		response := response.Responses(http.StatusBadRequest, "", result, nil)
		c.JSON(http.StatusBadRequest, response)
	} else {
		response := response.Responses(http.StatusOK, "Succesfully login", result, nil)
		c.JSON(http.StatusOK, response)
	}
}

// @Summary		All User
// @Description	using this handler admin can view user
// @Tags			Admins
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Param			page	query		int						false	"Page number"				default(1)
// @Param			limit	query		int						false	"Number of items per page"	default(5)
// @Success		200	{object}	response.Response{}
// @Failure		400	{object}	response.Response{}
// @Router			/admin/users/getuser [get]
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

// @Summary		Block User
// @Description	using this handler admin can block user
// @Tags			Admins
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Param			id	query		string	true	"User ID in the URL path"
// @Success		200	{object}	response.Response{}
// @Failure		400	{object}	response.Response{}
// @Router			/admin/users/block [patch]
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

// @Summary		Unblock User
// @Description	using this handler admin Unblock user
// @Tags			Admins
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Param			id	query		string	true	"User ID in the URL path"
// @Success		200	{object}	response.Response{}
// @Failure		400	{object}	response.Response{}
// @Router			/admin/users/unblock [patch]
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

// ------------------------------------------user Address------------------------------------\\

// @Summary		Add Address
// @Description	Add a new address.
// @Tags			User Addresses
// @Accept			json
// @Produce		json
// @security		BearerTokenAuth
// @security		Refreshtoken
// @Param			address	body		requestmodel.Address{}	true	"Address object to be added"
// @Success		201		{object}	response.Response{}		"Successfully added the address"
// @Failure		400		{object}	response.Response{}		"Bad request"
// @Router			/address [post]
func (u *UserHandler) NewAddress(c *gin.Context) {

	var Address requestmodel.Address

	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	Address.Userid = userID

	if err := c.ShouldBind(&Address); err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, resCustomError.BindingConflict, nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	data, err := helper.Validation(Address)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", data, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	userAddress, err := u.userUseCase.AddAddress(&Address)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", userAddress, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// @Summary		Get Addresses
// @Description	Retrieve a list of addresses.
// @Tags			User Addresses
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Security		Refreshtoken
// @Param			page	query		int						false	"Page number"				default(1)
// @Param			limit	query		int						false	"Number of items per page"	default(5)
// @Success		200		{object}	[]requestmodel.Address	"Successfully retrieved addresses"
// @Failure		400		{object}	response.Response		"Bad request"
// @Router			/address [get]
func (u *UserHandler) GetAddress(c *gin.Context) {

	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "1")

	userAddress, err := u.userUseCase.GetAddress(userID, page, limit)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", userAddress, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// @Summary		Update Address
// @Description	Update an existing address.
// @Tags			User Addresses
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Security		Refreshtoken
// @Param			address	body		requestmodel.EditAddress	true	"Updated address information"
// @Success		200		{object}	response.Response			"Successfully updated the address"
// @Failure		400		{object}	response.Response			"Bad request"
// @Router			/address [patch]
func (u *UserHandler) EditAddress(c *gin.Context) {

	var Address requestmodel.EditAddress

	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	Address.Userid = userID

	if err := c.ShouldBind(&Address); err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, resCustomError.BindingConflict, nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	userAddress, err := u.userUseCase.EditAddress(&Address)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Succesfully Edited", userAddress, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// @Summary		Delete Address
// @Description	Delete an address by ID.
// @Tags			User Addresses
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Security		Refreshtoken
// @Param			id	query		string				true	"Address ID in the query parameter"
// @Success		200	{object}	response.Response	"Successfully deleted the address"
// @Failure		400	{object}	response.Response	"Bad request"
// @Router			/address [delete]
func (u *UserHandler) DeleteAddress(c *gin.Context) {

	addressID := c.Query("id")
	id := strings.TrimSpace(addressID)

	if len(id) == 0 {
		finalReslt := response.Responses(http.StatusBadRequest, "", "", resCustomError.IDParamsEmpty)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	err := u.userUseCase.DeleteAddress(id, userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Succesfully Deleted", "", nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// ------------------------------------------user Profile------------------------------------\\

// @Summary		Get User
// @Description	Retrieve the user's profile.
// @Tags			User Profile
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Security		Refreshtoken
// @Success		200	{object}	requestmodel.UserDetails{}	"Successfully retrieved the user's profile"
// @Failure		400	{object}	response.Response			"Bad request"
// @Router			/profile [get]
func (u *UserHandler) GetProfile(c *gin.Context) {

	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	UserProfile, err := u.userUseCase.GetProfile(userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", UserProfile, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// @Summary		Update User Profile
// @Description	Update the user's profile.
// @Tags			User Profile
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Security		Refreshtoken
// @Param			profile	body		requestmodel.UserEditProfile	true	"User profile details for updating"
// @Success		200		{object}	response.Response				"Successfully updated the user's profile"
// @Failure		400		{object}	response.Response				"Bad request"
// @Router			/profile [patch]
func (u *UserHandler) EditProfile(c *gin.Context) {

	var profile requestmodel.UserEditProfile

	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	profile.Id = userID

	if err := c.ShouldBind(&profile); err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, resCustomError.BindingConflict, nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	userProfile, err := u.userUseCase.UpdateProfile(&profile)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Succesfully Edited", userProfile, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// ------------------------------------------User Forgot Password------------------------------------\\

// @Summary		Forgot Password
// @Description	Initiate the process for resetting the password.
// @Tags			User
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Param			forgotPassword	body		requestmodel.ForgotPassword	true	"Details for initiating password reset"
// @Success		200				{object}	response.Response			"Password reset initiated successfully"
// @Failure		400				{object}	response.Response			"Bad request"
// @Router			/forgotpassword [post]
func (u *UserHandler) ForgotPassword(c *gin.Context) {
	var ForgotPassword requestmodel.ForgotPassword

	if err := c.BindJSON(&ForgotPassword); err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, resCustomError.BindingConflict, nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	token := c.Request.Header.Get("Authorization")

	data, err := helper.Validation(ForgotPassword)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", data, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	err = u.userUseCase.ForgotPassword(&ForgotPassword, token)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Succesfully Edited", "", nil)
		c.JSON(http.StatusOK, finalReslt)
	}

}
