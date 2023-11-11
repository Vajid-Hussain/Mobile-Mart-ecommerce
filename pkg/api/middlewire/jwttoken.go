package middlewire

import (
	"fmt"
	"net/http"

	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/config"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/service"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type TokenRequirement struct {
	JwtTokenUseCase interfaceUseCase.IJwtTokenUseCase
	securityKeys    config.Token
}

var token TokenRequirement

func NewJwtTokenMiddleWire(jwtUseCase interfaceUseCase.IJwtTokenUseCase, keys config.Token) {
	token = TokenRequirement{JwtTokenUseCase: jwtUseCase, securityKeys: keys}
}

func TokenVerify(c *gin.Context) {
	accessToken := c.Request.Header.Get("accesstoken")
	refreshToken := c.Request.Header.Get("refreshtoken")

	id, status, err := service.VerifyAcessToken(accessToken, token.securityKeys.SellerSecurityKey)

	if err != nil {
		err := service.VerifyRefreshToken(refreshToken, token.securityKeys.SellerSecurityKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"err": err})
			fmt.Println(status, "--------1----", err)

		} else {
			status, err := token.JwtTokenUseCase.ValidateJwtToken(id)
			if err != nil {
				c.JSON(http.StatusUnauthorized, err)
				fmt.Println(status, "---------2---", err)

			} else {
				token, err := service.GenerateAcessToken(token.securityKeys.SellerSecurityKey, id, status)
				if err != nil {
					c.JSON(http.StatusUnauthorized, err)
					fmt.Println(status, "---------3---", err)

				} else {
					c.JSON(http.StatusOK, token)
					fmt.Println(status, "----------4--", token)

				}
			}
		}
	} else {
		c.JSON(http.StatusOK, "all perfect you access token is uptodate")
		fmt.Println("refresh id sonde")
	}
	fmt.Println(status, "chumma ------------")
}
