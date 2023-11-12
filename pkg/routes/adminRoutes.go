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
		}
	}

}
