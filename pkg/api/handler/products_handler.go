package handler

import (
	"fmt"
	"net/http"
	"strconv"

	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	resCustomError "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel/custom_error"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel/response"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/utils/helper"
	"github.com/gin-gonic/gin"
)

type InventotyHandler struct {
	userCase interfaceUseCase.IInventoryUseCase
}

func NewInventoryHandler(usercase interfaceUseCase.IInventoryUseCase) *InventotyHandler {
	return &InventotyHandler{userCase: usercase}
}

//	@Summary	Add Product
//
// // @Description Add a new product from the seller.
// // @Tags Seller Products
// // @Accept json
// // @Produce json
// // @Security BearerTokenAuth
// // @Security Refreshtoken
// // @Param product formData AddProductRequest true "Product details for adding"
// // @Success 201 {object} response.Response "Successfully added the product"
// // @Failure 400 {object} response.Response "Bad request"
// // @Router /seller/products [post]
func (u *InventotyHandler) AddInventory(c *gin.Context) {

	var inventoryDetails requestmodel.InventoryReq

	sellerid, exist := c.MustGet("SellerID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, resCustomError.NotGetSellerIDinContexr, nil, nil)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	if err := c.ShouldBind(&inventoryDetails); err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, resCustomError.BindingConflict, nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}
	sellerID, _ := strconv.ParseUint(sellerid, 16, 16)
	inventoryDetails.SellerID = uint(sellerID)
	fmt.Println("****", sellerID)

	data, err := helper.Validation(inventoryDetails)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", data, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	product, err := u.userCase.AddInventory(&inventoryDetails)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "refine request", "", err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "succesfully acomplish", product, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// @Summary		Block Product
// @Description	Block a product from being displayed.
// @Tags			Seller Products
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Security		Refreshtoken
// @Param			id	path		string				true	"Product ID in the URL path"
// @Success		200	{object}	response.Response	"Successfully blocked the product"
// @Failure		400	{object}	response.Response	"Bad request"
// @Router			/seller/products/{id}/block [patch]
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

// @Summary		Unblock Product
// @Description	Unblock a product for display.
// @Tags			Seller Products
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Security		Refreshtoken
// @Param			id	path		string				true	"Product ID in the URL path"
// @Success		200	{object}	response.Response	"Successfully unblocked the product"
// @Failure		400	{object}	response.Response	"Bad request"
// @Router			/seller/products/{id}/unblock [patch]
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

// @Summary		Delete Product
// @Description	Delete a product by ID.
// @Tags			Seller Products
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Security		Refreshtoken
// @Param			id	path		string				true	"Product ID in the URL path"
// @Success		200	{object}	response.Response	"Successfully deleted the product"
// @Failure		400	{object}	response.Response	"Bad request"
// @Router			/seller/products/{id} [delete]
func (u *InventotyHandler) DeleteInventory(c *gin.Context) {
	sellerid, exist := c.MustGet("Sellerid").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, resCustomError.NotGetSellerIDinContexr, nil, nil)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	productID := c.Param("productid")
	err := u.userCase.DeleteInventory(sellerid, productID)
	if err != nil {
		finalReslt := response.Responses(http.StatusNotFound, "", "", err.Error())
		c.JSON(http.StatusNotFound, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Succesfully product deleted", "", nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// @Summary		Get Seller Products
// @Description	Retrieve a list of products.
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			page	query		int					false	"Page number"				default(1)
// @Param			limit	query		int					false	"Number of items per page"	default(5)
// @Success		200		{object}	[]response.Response	"Successfully retrieved products"
// @Failure		400		{object}	response.Response	"Bad request"
// @Router			/ [get]
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

// @Summary		Get Seller Product
// @Description	Retrieve details of a single seller product.
// @Tags			Seller Products
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Security		Refreshtoken
// @Param			id	path		string				true	"Product ID in the URL path"
// @Success		200	{object}	response.Response	"Successfully retrieved the seller product"
// @Failure		400	{object}	response.Response	"Bad request"
// @Router			/seller/products/{id} [get]
func (u *InventotyHandler) GetAInventory(c *gin.Context) {
	id := c.Param("productid")

	inverntory, err := u.userCase.GetAInventory(id)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", inverntory, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// @Summary		Get Seller Products
// @Description	Retrieve a list of seller products.
// @Tags			Seller Products
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Security		Refreshtoken
// @Param			page	query		int					false	"Page number"				default(1)
// @Param			limit	query		int					false	"Number of items per page"	default(5)
// @Success		200		{object}	[]response.Response	"Successfully retrieved seller products"
// @Failure		400		{object}	response.Response	"Bad request"
// @Router			/seller/products [get]
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

// @Summary		Edit Seller Product
// @Description	Edit details of a seller product.
// @Tags			Seller Products
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Security		Refreshtoken
// @Param			productid	query		string						true	"Product ID in the query parameter"
// @Param			product		body		requestmodel.EditInventory	true	"Updated product details"
// @Success		200			{object}	response.Response			"Successfully edited the seller product"
// @Failure		400			{object}	response.Response			"Bad request"
// @Router			/seller/products [patch]
func (u *InventotyHandler) EditInventory(c *gin.Context) {

	inventoryID := c.Query("productid")
	var edittedInventory requestmodel.EditInventory

	if err := c.BindJSON(&edittedInventory); err != nil {
		fmt.Println(err)
		finalReslt := response.Responses(http.StatusBadRequest, resCustomError.BindingConflict, nil, nil)
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

func (u *InventotyHandler) FilterProduct(c *gin.Context) {

	var criterial requestmodel.FilterCriterion

	if err := c.BindJSON(&criterial); err != nil {
		fmt.Println(err)
		finalReslt := response.Responses(http.StatusBadRequest, resCustomError.BindingConflict, nil, nil)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	filteredProduct, err := u.userCase.GetProductFilter(&criterial)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", filteredProduct, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}
