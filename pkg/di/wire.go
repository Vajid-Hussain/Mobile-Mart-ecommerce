package di

import (
	server "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/api"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/api/handler"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/api/middlewire"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/config"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/db"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/repository"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/service"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase"
)

func InitializeAPI(config config.Config) (*server.ServerHttp, error) {
	DB, err := db.ConnectDatabase(config.DB)
	if err != nil {
		return nil, err
	}

	service.OtpService(config.Otp)

	jwtRepository := repository.NewJwtTokenRepository(DB)
	jwtUseCase := usecase.NewJwtTokenUseCase(jwtRepository)
	middlewire.NewJwtTokenMiddleWire(jwtUseCase, config.Token)

	userRepository := repository.NewUserRepository(DB)
	userUseCase := usecase.NewUserUseCase(userRepository, &config.Token)
	userHandler := handler.NewUserHandler(userUseCase)

	sellerRepository := repository.NewSellerRepository(DB)
	sellerUseCase := usecase.NewSellerUseCase(sellerRepository, &config.Token)
	sellerHandler := handler.NewSellerHandler(sellerUseCase)

	adminRepository := repository.NewAdminRepository(DB)
	adminUseCase := usecase.NewAdminUseCase(adminRepository, &config.Token)
	adminHandler := handler.NewAdminHandler(adminUseCase)

	categoryRepository := repository.NewCategoryRepository(DB)
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryUseCase)

	inventoryRepository := repository.NewInventoryRepository(DB)
	inventoryUseCase := usecase.NewInventoryUseCase(inventoryRepository, &config.S3aws)
	inventoryHandler := handler.NewInventoryHandler(inventoryUseCase)

	cartRepository := repository.NewCartRepository(DB)
	cartUseCase := usecase.NewCartUseCase(cartRepository)
	cartHanlder := handler.NewCartHandler(cartUseCase)

	orderRepository := repository.NewOrderRepository(DB)
	orderUseCase := usecase.NewOrderUseCase(orderRepository, cartRepository)
	orderHandler := handler.NewOrderHandler(orderUseCase)

	serverHttp := server.NewServerHttp(userHandler,
		sellerHandler,
		adminHandler,
		categoryHandler,
		inventoryHandler,
		cartHanlder,
		orderHandler,
	)

	return serverHttp, nil
}
