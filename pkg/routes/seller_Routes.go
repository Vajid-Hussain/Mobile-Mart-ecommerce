package routes

import (
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/api/handler"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/api/middlewire"
	"github.com/gin-gonic/gin"
)

func SellerRoutes(engin *gin.RouterGroup, seller *handler.SellerHandler, inventory *handler.InventotyHandler, order *handler.OrderHandler, category *handler.CategoryHandler, jwt *handler.TokenRequirement) {

	engin.POST("/signup", seller.SellerSignup)
	engin.POST("/login", seller.SellerLogin)

	engin.GET("/accesstoken", jwt.NewSellerAcessToken)

	engin.Use(middlewire.SellerAuthorization)
	{
		engin.GET("/", seller.SellerDashbord)
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
			ordermanagenent.GET("/cancelled", order.GetSellerOrdersCancelled)
			ordermanagenent.PATCH("/", order.ConfirmDeliverd)
			ordermanagenent.PATCH("/:orderID/cancel", order.CancelOrder)
		}

		salesreportmanagement := engin.Group("/report")
		{
			// salesreportmanagement.GET("", order.SalesReportByYear)
			// salesreportmanagement.GET("/month", order.SalesReportByMonth)
			// salesreportmanagement.GET("/week", order.SalesReportByWeek)
			salesreportmanagement.GET("/day", order.SalesReport)
			salesreportmanagement.GET("/days", order.SalesReportCustomDays)
			salesreportmanagement.GET("xlsx", order.SalesReportXlSX)
		}

		categorymanagement := engin.Group("/categoryoffer")
		{
			categorymanagement.GET("/brand", category.FetchAllBrand)
			categorymanagement.GET("/category", category.FetchAllCatogry)
			categorymanagement.GET("/", category.GetAllCategoryOffer)
			categorymanagement.POST("/", category.CreateCategoryOffer)
			categorymanagement.PATCH("/", category.EditCategoryOffer)
			categorymanagement.PATCH("/block", category.BlockCategoryOffer)
			categorymanagement.PATCH("/unblock", category.UnBlockCategoryOffer)
			categorymanagement.DELETE("/delete", category.DeleteCategoryOffer)
		}
	}
}
