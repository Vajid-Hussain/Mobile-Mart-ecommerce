package server

import (
	"fmt"
	"log"

	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/api/handler"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHttp struct {
	engin *gin.Engine
}

func NewServerHttp(user *handler.UserHandler,
	seller *handler.SellerHandler,
	admin *handler.AdminHandler,
	category *handler.CategoryHandler,
	inventory *handler.InventotyHandler,
	cart *handler.CartHandler,
	order *handler.OrderHandler,
	payment *handler.PaymentHandler,
	coupon *handler.CouponHandler,
) *ServerHttp {
	engin := gin.New()
	engin.Use(gin.Logger())

	// load htmlpages
	engin.LoadHTMLGlob("./template/*.html")

	// use ginSwagger middleware to serve the API docs
	engin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.UserRoutes(engin.Group("/"), user, inventory, cart, order, payment)
	routes.SellerRoutes(engin.Group("/seller"), seller, inventory, order)
	routes.AdminRoutes(engin.Group("/admin"), admin, seller, user, category, coupon)

	return &ServerHttp{engin: engin}
}

func (server *ServerHttp) Start() {
	err := server.engin.Run(":7000")
	if err != nil {
		log.Fatal("gin engin couldn't start")
	}
	fmt.Println("gin engin start")
}
