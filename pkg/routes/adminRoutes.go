package routes

import (
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/api/handler"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(engin *gin.RouterGroup,
	admin *handler.AdminHandler) {

	engin.POST("/login", admin.AdminLogin)
}
