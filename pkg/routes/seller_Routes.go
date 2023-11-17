package routes

import (
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/api/handler"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/api/middlewire"
	"github.com/gin-gonic/gin"
)

func SellerRoutes(engin *gin.RouterGroup, seller *handler.SellerHandler, inventory *handler.InventotyHandler) {

	engin.POST("/signup", seller.SellerSignup)
	engin.POST("/login", seller.SellerLogin)

	engin.Use(middlewire.SellerAuthorization)
	{
		inventorymanagement := engin.Group("/inventory")
		{
			inventorymanagement.POST("/", inventory.AddInventory)
			inventorymanagement.GET("/", inventory.GetSellerInventory)
			inventorymanagement.GET("/:inventoryid", inventory.GetAInventory)
			inventorymanagement.PATCH("/", inventory.EditInventory)
			inventorymanagement.DELETE("/:inventoryid", inventory.DeleteInventory)
			inventorymanagement.PATCH("/:productid/block", inventory.BlockInventory)
			inventorymanagement.PATCH("/:productid/unblock", inventory.UNBlockInventory)
		}
	}
}
