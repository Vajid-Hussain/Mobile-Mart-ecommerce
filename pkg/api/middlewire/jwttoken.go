package middlewire

import (
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/service"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

var JwtTokenUseCase interfaceUseCase.IJwtTokenUseCase

func NewJwtTokenMiddleWire(jwtUseCase interfaceUseCase.IJwtTokenUseCase) {
	JwtTokenUseCase = jwtUseCase
}

func TokenVerify(c *gin.Context){
	accessToken:=c.Request.Header.Get("tokenname")
	refreshToken:=c.Request.Header.Get("refreshtoke")

	service.
}