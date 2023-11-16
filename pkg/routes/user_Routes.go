package routes

import (
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/api/handler"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/api/middlewire"
	"github.com/gin-gonic/gin"
)

func UserRoutes(engin *gin.RouterGroup, user *handler.UserHandler, inventory *handler.InventotyHandler) {

	usermanagement := engin.Group("/user")
	{
		usermanagement.POST("/signup", user.UserSignup)
		usermanagement.POST("/verifyOTP", user.VerifyOTP)
		usermanagement.POST("/login", user.UserLogin)

		engin.Use(middlewire.UserAuthorization)
		{
			engin.GET("/", inventory.GetInventory)
			engin.GET("/:inventoryid", inventory.GetAInventory)
		}
	}

}
