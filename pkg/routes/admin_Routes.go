package routes

import (
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/api/handler"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/api/middlewire"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(engin *gin.RouterGroup, admin *handler.AdminHandler, seller *handler.SellerHandler, user *handler.UserHandler, category *handler.CategoryHandler, coupon *handler.CouponHandler) {

	engin.POST("/login", admin.AdminLogin)

	engin.Use(middlewire.AdminAuthorization)
	{
		engin.GET("/", admin.AdminDashBord)

		usermanagement := engin.Group("/users")
		{
			usermanagement.GET("/getuser", user.GetUser)
			usermanagement.PATCH("/block", user.BlockUser)
			usermanagement.PATCH("/unblock", user.UnblockUser)
		}

		sellermanagement := engin.Group("/sellers")
		{
			sellermanagement.GET("/getsellers", seller.GetSellers)
			sellermanagement.PATCH("/block", seller.BlockSeller)
			sellermanagement.PATCH("/unblock", seller.UnblockSeller)
			sellermanagement.GET("/pending", seller.GetPendingSellers)
			sellermanagement.GET("/singleview", seller.FetchSingleSeller)
			sellermanagement.PATCH("/verify", seller.VerifySeller)
		}

		categorymanagement := engin.Group("/category")
		{
			categorymanagement.POST("/", category.NewCategory)
			categorymanagement.GET("/", category.FetchAllCatogry)
			categorymanagement.PATCH("/", category.UpdateCategory)
			categorymanagement.DELETE("/", category.DeleteCategory)

		}

		brandmanagement := engin.Group("/brand")
		{
			brandmanagement.POST("/", category.CreateBrand)
			brandmanagement.GET("/", category.FetchAllBrand)
			brandmanagement.PATCH("/", category.UpdateBrand)
			brandmanagement.DELETE("/", category.DeleteBrand)
		}

		couponmanagment := engin.Group("/coupon")
		{
			couponmanagment.POST("/", coupon.CreateCoupon)
		}
	}
}
