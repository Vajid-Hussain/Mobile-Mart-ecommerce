package routes

import (
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/api/handler"
	"github.com/gin-gonic/gin"
)

func UserRoutes(engin *gin.RouterGroup, user *handler.UserHandler) {

	engin.POST("/signup", user.UserSignup)
	engin.POST("/verifyOTP", user.VerifyOTP)
	engin.POST("/login", user.UserLogin)

}
