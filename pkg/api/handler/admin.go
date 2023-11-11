package handler

import (
	"fmt"
	"net/http"

	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel/response"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	AdminUseCase interfaceUseCase.IAdminUseCase
}

func NewAdminHandler(useCase interfaceUseCase.IAdminUseCase) *AdminHandler {
	return &AdminHandler{AdminUseCase: useCase}
}

func (u *AdminHandler) AdminLogin(c *gin.Context) {
	var loginCredential requestmodel.AdminLoginData

	err := c.BindJSON(&loginCredential)
	if err != nil {
		finalReslt := response.Responses(http.StatusUnauthorized, "json is wrong can't bind", nil, "Double chuck json")
		c.JSON(http.StatusUnauthorized, finalReslt)
		return
	}

	result, err := u.AdminUseCase.AdminLogin(&loginCredential)
	fmt.Println(result, "=========")
	if err != nil {
		finalReslt := response.Responses(http.StatusUnauthorized, "", result, err)
		c.JSON(http.StatusUnauthorized, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "succesfully login", result, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}
