package routes

import (
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/api/handler"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/api/middlewire"
	"github.com/gin-gonic/gin"
)

func SellerRoutes(engin *gin.RouterGroup, seller *handler.SellerHandler, inventory *handler.InventotyHandler, order *handler.OrderHandler) {

	engin.POST("/signup", seller.SellerSignup)
	engin.POST("/login", seller.SellerLogin)

	engin.Use(middlewire.SellerAuthorization)
	{
		engin.GET("/profile", seller.GetSellerProfile)
		engin.PATCH("/profile", seller.EditSellerProfile)

		inventorymanagement := engin.Group("/products")
		{
			inventorymanagement.POST("/", inventory.AddInventory)
			inventorymanagement.GET("/", inventory.GetSellerInventory)
			inventorymanagement.GET("/:productid", inventory.GetAInventory)
			inventorymanagement.PATCH("/", inventory.EditInventory)
			inventorymanagement.DELETE("/:productid", inventory.DeleteInventory)
			inventorymanagement.PATCH("/:productid/block", inventory.BlockInventory)
			inventorymanagement.PATCH("/:productid/unblock", inventory.UNBlockInventory)
		}

		ordermanagenent := engin.Group("/order")
		{
			ordermanagenent.GET("", order.GetSellerOrders)
			ordermanagenent.GET("/processing", order.GetSellerOrdersProcessing)
			ordermanagenent.GET("/delivered", order.GetSellerOrdersDeliverd)
			ordermanagenent.PATCH("/", order.ConfirmDeliverd)
			ordermanagenent.PATCH("/:orderID/cancel", order.CancelOrder)
		}

		salesreportmanagement := engin.Group("/report")
		{
			salesreportmanagement.GET("", order.SalesReportByYear)
			salesreportmanagement.GET("/month", order.SalesReportByMonth)
			salesreportmanagement.GET("/week", order.SalesReportByWeek)
			salesreportmanagement.GET("/day", order.SalesReportByDay)
		}
	}
}
