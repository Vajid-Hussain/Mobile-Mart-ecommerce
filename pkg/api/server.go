package server

import (
	"fmt"
	"log"

	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/api/handler"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/routes"
	"github.com/gin-gonic/gin"
)

type ServerHttp struct {
	engin *gin.Engine
}

func NewServerHttp(user *handler.UserHandler, seller *handler.SellerHandler, admin *handler.AdminHandler) *ServerHttp {
	engin := gin.New()
	engin.Use(gin.Logger())

	routes.UserRoutes(engin.Group("/user"), user)
	routes.SellerRoutes(engin.Group("/seller"), seller)
	routes.AdminRoutes(engin.Group("/admin"), admin)

	return &ServerHttp{engin: engin}
}

func (server *ServerHttp) Start() {
	err := server.engin.Run(":7000")
	if err != nil {
		log.Fatal("gin engin couldn't start")
	}
	fmt.Println("gin engin start")
}
