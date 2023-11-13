package routes

import (
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/api/handler"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/api/middlewire"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(engin *gin.RouterGroup, admin *handler.AdminHandler, seller *handler.SellerHandler, user *handler.UserHandler, category *handler.CategoryHandler) {

	engin.POST("/login", admin.AdminLogin)

	engin.Use(middlewire.AdminAuthorization)
	{
		usermanagement := engin.Group("/users")
		{
			usermanagement.GET("/getuser", user.GetUser)
			usermanagement.GET("/block", user.BlockUser)
			usermanagement.GET("/unblock", user.UnblockUser)
		}

		sellermanagement := engin.Group("/sellers")
		{
			sellermanagement.GET("/getsellers", seller.GetSellers)
			sellermanagement.GET("/block", seller.BlockSeller)
			sellermanagement.GET("/unblock", seller.UnblockSeller)
			sellermanagement.GET("/pending", seller.GetPendingSellers)
			sellermanagement.GET("/singleview", seller.FetchSingleSeller)
		}

		categorymanagement := engin.Group("/category")
		{
			categorymanagement.POST("/add", category.NewCategory)
			categorymanagement.GET("/all", category.FetchAllCatogry)

		}
	}
}
