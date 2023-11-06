package di

import (
	server "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/api"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/api/handler"
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

	otpServices := service.NewOtpService(config.Otp)

	userRepository := repository.NewUserRepository(DB)
	userUseCase := usecase.NewUserUseCase(userRepository, otpServices)
	userHandler := handler.NewUserHandler(userUseCase)

	serverHttp := server.NewServerHttp(userHandler)

	return serverHttp, nil
}
