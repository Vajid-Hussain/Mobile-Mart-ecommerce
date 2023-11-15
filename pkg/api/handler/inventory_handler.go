package handler

import (
	"net/http"

	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	resCustomError "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel/custom_error"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel/response"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type InventotyHandler struct {
	userCase interfaceUseCase.IInventoryUseCase
}

func NewInventoryHandler(usercase interfaceUseCase.IInventoryUseCase) *InventotyHandler {
	return &InventotyHandler{userCase: usercase}
}

func (u *InventotyHandler) AddInventory(c *gin.Context) {
	var inventoryDetails requestmodel.InventoryReq
	sellerid, exist := c.MustGet("SellerID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "not got seller id ", nil, nil)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	inventoryDetails.SellerID = sellerid

	if err := c.BindJSON(&inventoryDetails); err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, resCustomError.BindingConflict, nil, nil)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	result, product, err := u.userCase.AddInventory(&inventoryDetails)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "refine request", result, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "succesfully acomplish", product, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}
