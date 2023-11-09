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
	usecase interfaceUseCase.IVenderUseCase
}

func NewSellerHandler(venderUuseCase interfaceUseCase.IVenderUseCase) *SellerHandler {
	return &SellerHandler{usecase: venderUuseCase}
}

func (u *SellerHandler) SellerSignup(c *gin.Context) {
	var sellerDetails requestmodel.SellerSignup
	if err := c.BindJSON(&sellerDetails); err != nil {
		fmt.Println(err)
	}

	result, err := u.usecase.SellerSignup(sellerDetails)
	if err != nil {
		finalReslt:=response.Responses(http.StatusUnauthorized, err.Error(), result, nil)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}else{
		finalReslt:=response.Responses(http.StatusOK, "succesfully signup", result, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}
                    



