package handler

import (
	"fmt"
	"net/http"

	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	resCustomError "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel/custom_error"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel/response"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/utils/helper"
	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	useCase interfaceUseCase.IPaymentUseCase
}

func NewPaymentHandler(useCase interfaceUseCase.IPaymentUseCase) *PaymentHandler {
	return &PaymentHandler{useCase: useCase}
}

//	@Summary		Get Razorpay Payment Page
//	@Description	Retrieve the Razorpay payment page for the specified user.
//	@Tags			PaymentIntegration
//	@Accept			html
//	@Produce		html
//	@Param			userID	query		int					true	"User ID for which the payment page is requested"
//	@Success		200		{string}	html				"HTML page for Razorpay payment"
//	@Failure		400		{object}	response.Response	"Bad request. Please provide a valid user ID."
//	@Router			/razorpay [get]
func (u *PaymentHandler) OnlinePayment(c *gin.Context) {
	userID := c.Query("userID")
	orderID := c.Query("orderID")
	fmt.Println("**", userID, orderID)
	orderDetails, err := u.useCase.OnlinePayment(userID, orderID)
	if err != nil {
		c.HTML(http.StatusBadRequest, "razopay.html", gin.H{"badRequest": "Refine your request"})
	} else {
		c.HTML(http.StatusOK, "razopay.html", orderDetails)
	}
}

//	@Summary		Verify Online Payment
//	@Description	Verify an online payment using the provided details.
//	@Tags			PaymentIntegration
//	@Accept			json
//	@Produce		json
//	@Security		BearerTokenAuth
//	@Security		Refreshtoken
//	@Param			verificationDetails	body		requestmodel.OnlinePaymentVerification	true	"Details for online payment verification"
//	@Success		200					{object}	response.Response						"Payment verification successful"
//	@Failure		400					{object}	response.Response						"Bad request. Please provide valid verification details."
//	@Router			/payment/verify [post]
func (u *PaymentHandler) VerifyOnlinePayment(c *gin.Context) {
	var onlinePaymentDetails requestmodel.OnlinePaymentVerification

	if err := c.BindJSON(&onlinePaymentDetails); err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, resCustomError.BindingConflict, nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	data, err := helper.Validation(onlinePaymentDetails)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", data, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	order, err := u.useCase.OnlinePaymentVerification(&onlinePaymentDetails)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", order, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

//	@Summary		Get User Wallet
//	@Description	Retrieve details of the user's wallet.
//	@Tags			User Wallet
//	@Accept			json
//	@Produce		json
//	@Security		BearerTokenAuth
//	@Security		Refreshtoken
//	@Success		200	{object}	response.Response	"User wallet details retrieved successfully"
//	@Failure		400	{object}	response.Response	"Bad request. Unable to retrieve user wallet details."
//	@Router			/wallet [get]
func (u *PaymentHandler) ViewWallet(c *gin.Context) {
	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	userWallet, err := u.useCase.GetUserWallet(userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", userWallet, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

//	@Summary		Get User Wallet Transactions
//	@Description	Retrieve the transactions history of the user's wallet.
//	@Tags			User Wallet
//	@Accept			json
//	@Produce		json
//	@Security		BearerTokenAuth
//	@Security		Refreshtoken
//	@Success		200	{object}	response.Response	"User wallet transactions retrieved successfully"
//	@Failure		400	{object}	response.Response	"Bad request. Unable to retrieve user wallet transactions."
//	@Router			/wallet/transaction [get]
func (u *PaymentHandler) GetWalletTransaction(c *gin.Context) {
	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	walletTransactions, err := u.useCase.GetWalletTransaction(userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", walletTransactions, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}
