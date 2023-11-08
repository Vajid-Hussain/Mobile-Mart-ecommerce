package routes

import (
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/api/handler"
	"github.com/gin-gonic/gin"
)


func SellerRoutes(engin *gin.RouterGroup,
	seller *handler.SellerHandler){

	engin.POST("/signup", seller.SellerSignup)
}