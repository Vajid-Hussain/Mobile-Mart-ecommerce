package routes

import (
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/api/handler"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/api/middlewire"
	"github.com/gin-gonic/gin"
)

func SellerRoutes(engin *gin.RouterGroup, seller *handler.SellerHandler) {

	engin.POST("/signup", seller.SellerSignup)
	engin.POST("/login", seller.SellerLogin)

	engin.Use(middlewire.SellerAuthorization)
	// {

	// }

}
