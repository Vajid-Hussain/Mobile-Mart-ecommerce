package handler

import (
	"fmt"
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
		finalReslt := response.Responses(http.StatusBadRequest, resCustomError.NotGetSellerIDinContexr, nil, nil)
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

func (u *InventotyHandler) BlockInventory(c *gin.Context) {
	sellerid, exist := c.MustGet("SellerID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, resCustomError.NotGetSellerIDinContexr, nil, nil)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	productID := c.Param("productid")
	err := u.userCase.BlockInventory(sellerid, productID)
	if err != nil {
		finalReslt := response.Responses(http.StatusNotFound, "", "", err.Error())
		c.JSON(http.StatusNotFound, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Succesfully product blocked", "", nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *InventotyHandler) UNBlockInventory(c *gin.Context) {
	sellerid, exist := c.MustGet("SellerID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, resCustomError.NotGetSellerIDinContexr, nil, nil)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	productID := c.Param("productid")
	err := u.userCase.UNBlockInventory(sellerid, productID)
	if err != nil {
		finalReslt := response.Responses(http.StatusNotFound, "", "", err.Error())
		c.JSON(http.StatusNotFound, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Succesfully product unblocked", "", nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *InventotyHandler) DeleteInventory(c *gin.Context) {
	sellerid, exist := c.MustGet("SellerID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, resCustomError.NotGetSellerIDinContexr, nil, nil)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	productID := c.Param("inventoryid")
	err := u.userCase.DeleteInventory(sellerid, productID)
	if err != nil {
		finalReslt := response.Responses(http.StatusNotFound, "", "", err.Error())
		c.JSON(http.StatusNotFound, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Succesfully product deleted", "", nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *InventotyHandler) GetInventory(c *gin.Context) {
	page := c.Query("page")
	limit := c.DefaultQuery("limit", "1")

	inverntories, err := u.userCase.GetAllInventory(page, limit)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		finalReslt := response.Responses(http.StatusOK, "", inverntories, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *InventotyHandler) GetAInventory(c *gin.Context) {
	id := c.Param("inventoryid")

	inverntory, err := u.userCase.GetAInventory(id)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", inverntory, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *InventotyHandler) GetSellerInventory(c *gin.Context) {
	page := c.Query("page")
	limit := c.DefaultQuery("limit", "1")
	sellerID := c.MustGet("SellerID").(string)

	inverntories, err := u.userCase.GetSellerInventory(page, limit, sellerID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", inverntories, nil)
		c.JSON(http.StatusOK, finalReslt)
	}

}

func (u *InventotyHandler) EditInventory(c *gin.Context) {

	inventoryID := c.Query("invenoryid")
	var edittedInventory requestmodel.EditInventory

	if err := c.BindJSON(&edittedInventory); err != nil {
		fmt.Println(err)
		finalReslt := response.Responses(http.StatusBadRequest, "something wrong", nil, nil)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	updatedInverntory, err := u.userCase.EditInventory(&edittedInventory, inventoryID)

	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", updatedInverntory, nil)
		c.JSON(http.StatusOK, finalReslt)
	}

}
