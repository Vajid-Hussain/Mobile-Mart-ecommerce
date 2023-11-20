package routes

import (
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/api/handler"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/api/middlewire"
	"github.com/gin-gonic/gin"
)

func UserRoutes(engin *gin.RouterGroup, user *handler.UserHandler, inventory *handler.InventotyHandler) {

	engin.GET("/", inventory.GetInventory)
	engin.POST("/signup", user.UserSignup)
	engin.POST("/verifyOTP", user.VerifyOTP)
	engin.POST("/sendotp", user.SendOtp)
	engin.POST("/login", user.UserLogin)
	engin.POST("/forgotpassword", user.ForgotPassword)

	engin.Use(middlewire.UserAuthorization)
	{
		engin.GET("/:inventoryid", inventory.GetAInventory)

		engin.POST("/address", user.NewAddress)
		engin.GET("/address", user.GetAddress)
		engin.PATCH("/address", user.EditAddress)
		engin.DELETE("/address", user.DeleteAddress)

		engin.GET("/profile", user.GetProfile)
		engin.PATCH("/profile", user.EditProfile)

	}
}
