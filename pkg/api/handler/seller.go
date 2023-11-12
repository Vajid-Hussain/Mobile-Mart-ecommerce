package handler

import (
	"fmt"
	"net/http"

	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel/response"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type SellerHandler struct {
	usecase interfaceUseCase.ISellerUseCase
}

func NewSellerHandler(venderUuseCase interfaceUseCase.ISellerUseCase) *SellerHandler {
	return &SellerHandler{usecase: venderUuseCase}
}

func (u *SellerHandler) SellerSignup(c *gin.Context) {
	var sellerDetails requestmodel.SellerSignup

	if err := c.BindJSON(&sellerDetails); err != nil {
		c.JSON(http.StatusBadRequest, "can't bind json with struct")
	}

	result, err := u.usecase.SellerSignup(&sellerDetails)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", result, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	} else {
		finalReslt := response.Responses(http.StatusOK, "succesfully signup", result, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *SellerHandler) SellerLogin(c *gin.Context) {
	var loginData requestmodel.SellerLogin
	fmt.Println("seller login handler ***************s")
	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, "can't bind json with struct")
	}

	result, err := u.usecase.SellerLogin(&loginData)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", result, err.Error())

		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "succesfully login", result, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}
