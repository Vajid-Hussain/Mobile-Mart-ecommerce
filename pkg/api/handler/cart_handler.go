package handler

import (
	"net/http"
	"strings"

	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	resCustomError "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel/custom_error"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel/response"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	useCase interfaceUseCase.ICartUseCase
}

func NewCartHandler(carUseCase interfaceUseCase.ICartUseCase) *CartHandler {
	return &CartHandler{useCase: carUseCase}
}

//	@Summary		Create User Cart
//	@Description	Create a user's cart.
//	@Tags			UserCart
//	@Accept			json
//	@Produce		json
//	@Security		BearerTokenAuth
//	@Security		Refreshtoken
//	@Param			cart	body		requestmodel.Cart	true	"Cart details for creating"
//	@Success		201		{object}	response.Response	"User cart created successfully"
//	@Failure		400		{object}	response.Response	"Bad request"
//	@Router			/cart [post]
func (u *CartHandler) CreateCart(c *gin.Context) {

	var cart requestmodel.Cart

	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	cart.UserID = userID

	if err := c.ShouldBind(&cart); err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, resCustomError.BindingConflict, nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	result, err := u.useCase.CreateCart(&cart)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Succesfully added", result, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

//	@Summary		Delete Item from User Cart
//	@Description	Delete a product from the user's cart.
//	@Tags			UserCart
//	@Accept			json
//	@Produce		json
//	@Security		BearerTokenAuth
//	@Security		Refreshtoken
//	@Param			productID	query		string				true	"Product ID to delete from the cart"
//	@Success		200			{object}	response.Response	"Product deleted from the cart successfully"
//	@Failure		400			{object}	response.Response	"Bad request"
//	@Router			/cart [delete]
func (u *CartHandler) DeleteInventoryFromCart(c *gin.Context) {

	inventoryID := c.Query("productID")
	id := strings.TrimSpace(inventoryID)

	if len(id) == 0 {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.IDParamsEmpty)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	err := u.useCase.DeleteInventoryFromCart(id, userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Succesfully Deleted", "", nil)
		c.JSON(http.StatusOK, finalReslt)
	}

}

//	@Summary		Increment Product Count in User Cart
//	@Description	Increase the count of a product in the user's cart.
//	@Tags			UserCart
//	@Accept			json
//	@Produce		json
//	@Security		BearerTokenAuth
//	@Security		Refreshtoken
//	@Param			inventoryid	query		string				true	"Inventory ID of the product to increment in the cart"
//	@Success		200			{object}	response.Response	"Product count incremented in the cart successfully"
//	@Failure		400			{object}	response.Response	"Bad request"
//	@Router			/cart/increment [patch]
func (u *CartHandler) IncrementQuantityCart(c *gin.Context) {

	inventoryID := c.Query("productID")
	id := strings.TrimSpace(inventoryID)

	if len(id) == 0 {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.IDParamsEmpty)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	result, err := u.useCase.QuantityIncriment(id, userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", result, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

//	@Summary		Decrement Product Count in User Cart
//	@Description	Decrease the count of a product in the user's cart.
//	@Tags			UserCart
//	@Accept			json
//	@Produce		json
//	@Security		BearerTokenAuth
//	@Security		Refreshtoken
//	@Param			productID	path		string				true	"Product ID to decrement in the cart"
//	@Success		200			{object}	response.Response	"Product count decremented in the cart successfully"
//	@Failure		400			{object}	response.Response	"Bad request"
//	@Router			/cart/decrement/{productID} [patch]
func (u *CartHandler) DecrementQuantityCart(c *gin.Context) {

	id := c.Param("productID")

	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	result, err := u.useCase.QuantityDecrease(id, userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", result, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

//	@Summary		Get User Cart
//	@Description	Retrieve all items in the user's cart.
//	@Tags			UserCart
//	@Accept			json
//	@Produce		json
//	@Security		BearerTokenAuth
//	@Security		Refreshtoken
//	@Success		200	{object}	response.Response	"Successfully retrieved user cart items"
//	@Failure		400	{object}	response.Response	"Bad request"
//	@Router			/cart [get]
func (u *CartHandler) ShowCart(c *gin.Context) {

	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	result, err := u.useCase.ShowCart(userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", result, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}
