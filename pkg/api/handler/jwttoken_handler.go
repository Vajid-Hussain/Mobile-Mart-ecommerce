package handler

import (
	"errors"
	"net/http"

	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/config"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel/response"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/service"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type TokenRequirement struct {
	JwtTokenUseCase interfaceUseCase.IJwtTokenUseCase
	securityKeys    config.Token
}

func NewJwtTokenHandler(jwtUseCase interfaceUseCase.IJwtTokenUseCase, keys config.Token) *TokenRequirement {
	return &TokenRequirement{JwtTokenUseCase: jwtUseCase, securityKeys: keys}
}

func (u *TokenRequirement) NewUserAcessToken(c *gin.Context) {
	token := c.Query("accesstoken")
	var newToken string

	id, err := service.VerifyAcessToken(token, u.securityKeys.UserSecurityKey)
	if id == "" {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, errors.New("can't fetch userid from token. There is no possibility for creating a new access token"))
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}
	if err != nil {
		_, err := u.JwtTokenUseCase.GetStatusOfUser(id)
		if err != nil {
			finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
			c.JSON(http.StatusBadRequest, finalReslt)
			return
		}

		newToken, err = service.GenerateAcessToken(u.securityKeys.UserSecurityKey, id)
		if err != nil {
			finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
			c.JSON(http.StatusBadRequest, finalReslt)
			return
		}
	} else {
		finalReslt := response.Responses(http.StatusOK, "access token is uptodate", token, nil)
		c.JSON(http.StatusOK, finalReslt)
		return
	}
	finalReslt := response.Responses(http.StatusOK, "", newToken, nil)
	c.JSON(http.StatusOK, finalReslt)
}

func (u *TokenRequirement) NewSellerAcessToken(c *gin.Context) {
	token := c.Query("accesstoken")
	var newToken string

	id, err := service.VerifyAcessToken(token, u.securityKeys.SellerSecurityKey)
	if id == "" {
		finalReslt := response.Responses(http.StatusBadRequest, err.Error(), nil, errors.New("can't fetch userid from token. There is no possibility for creating a new access token").Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}
	if err != nil {
		_, err := u.JwtTokenUseCase.GetDataForCreteAccessToken(id)
		if err != nil {
			finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
			c.JSON(http.StatusBadRequest, finalReslt)
			return
		}

		newToken, err = service.GenerateAcessToken(u.securityKeys.SellerSecurityKey, id)
		if err != nil {
			finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
			c.JSON(http.StatusBadRequest, finalReslt)
			return
		}
	} else {
		finalReslt := response.Responses(http.StatusOK, "access token is uptodate", token, nil)
		c.JSON(http.StatusOK, finalReslt)
		return
	}
	finalReslt := response.Responses(http.StatusOK, "", newToken, nil)
	c.JSON(http.StatusOK, finalReslt)
}
