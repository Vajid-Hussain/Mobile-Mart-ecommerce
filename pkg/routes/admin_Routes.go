package routes

import (
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/api/handler"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/api/middlewire"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(engin *gin.RouterGroup, admin *handler.AdminHandler) {

	engin.POST("/login", admin.AdminLogin)

	engin.Use(middlewire.AdminAuthorization)
	{
		usermanagement := engin.Group("/users")
		{
			usermanagement.GET("/getuser", admin.GetUser)
			usermanagement.GET("/block", admin.BlockUser)
			usermanagement.GET("/unblock", admin.UnblockUser)
		}

		sellermanagement := engin.Group("/sellers")
		{
			sellermanagement.GET("/getsellers", admin.GetSellers)
			sellermanagement.GET("/block", admin.BlockSeller)
			sellermanagement.GET("/unblock", admin.UnblockSeller)
			sellermanagement.GET("/pending", admin.GetPendingSellers)
			sellermanagement.GET("/singleview", admin.FetchSingleSeller)
		}
	}
}
