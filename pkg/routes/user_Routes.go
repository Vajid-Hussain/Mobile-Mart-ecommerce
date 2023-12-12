package routes

import (
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/api/handler"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/api/middlewire"
	"github.com/gin-gonic/gin"
)

func UserRoutes(engin *gin.RouterGroup, user *handler.UserHandler, inventory *handler.InventotyHandler, cart *handler.CartHandler, order *handler.OrderHandler, payment *handler.PaymentHandler, jwt *handler.TokenRequirement) {

	engin.GET("/", inventory.GetInventory)
	engin.GET("/:productid", inventory.GetAInventory)
	engin.GET("/filter", inventory.FilterProduct)

	engin.GET("/razopay", payment.OnlinePayment)

	engin.POST("/signup", user.UserSignup)
	engin.POST("/verifyOTP", user.VerifyOTP)
	engin.POST("/sendotp", user.SendOtp)
	engin.POST("/login", user.UserLogin)
	engin.POST("/forgotpassword", user.ForgotPassword)

	engin.GET("/accesstoken", jwt.NewUserAcessToken)

	engin.Use(middlewire.UserAuthorization)
	{
		addressmanagement := engin.Group("/address")
		{
			addressmanagement.POST("/", user.NewAddress)
			addressmanagement.GET("/", user.GetAddress)
			addressmanagement.PATCH("/", user.EditAddress)
			addressmanagement.DELETE("/", user.DeleteAddress)
		}

		profilemanagement := engin.Group("/profile")
		{
			profilemanagement.GET("/", user.GetProfile)
			profilemanagement.PATCH("/", user.EditProfile)
		}

		cartmanagement := engin.Group("/cart")
		{
			cartmanagement.POST("/", cart.CreateCart)
			cartmanagement.DELETE("/", cart.DeleteInventoryFromCart)
			cartmanagement.PATCH("/increment", cart.IncrementQuantityCart)
			cartmanagement.PATCH("/decrement/:productID", cart.DecrementQuantityCart)
			cartmanagement.GET("/", cart.ShowCart)

		}

		ordermanagement := engin.Group("/order")
		{
			ordermanagement.POST("", order.NewOrder)
			ordermanagement.GET("", order.ShowAbstractOrders)
			ordermanagement.GET(":orderID", order.SingleOrderDetails)
			ordermanagement.PATCH("", order.CancelUserOrder)
			ordermanagement.PATCH("/return", order.ReturnUserOrder)
			ordermanagement.GET("/invoice", order.GetInvoice)
		}

		paymentmanagement := engin.Group("/payment")
		{
			paymentmanagement.POST("/verify", payment.VerifyOnlinePayment)
		}

		walletmenagement := engin.Group("/wallet")
		{
			walletmenagement.GET("/", payment.ViewWallet)
			walletmenagement.GET("/transaction", payment.GetWalletTransaction)
		}
	}
}
