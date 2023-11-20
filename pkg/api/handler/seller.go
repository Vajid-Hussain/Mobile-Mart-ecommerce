package handler

import (
	"fmt"
	"net/http"
	"strings"

	models "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/model"
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	resCustomError "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel/custom_error"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel/response"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/utils/helper"
	"github.com/gin-gonic/gin"
)

type SellerHandler struct {
	usecase interfaceUseCase.ISellerUseCase
}

func NewSellerHandler(venderUuseCase interfaceUseCase.ISellerUseCase) *SellerHandler {
	return &SellerHandler{usecase: venderUuseCase}
}

// @Summary		Seller Signup
// @Description	using this handler Seller can signup
// @Tags			Seller
// @Accept			json
// @Produce		json
// @Param			Seller	body		requestmodel.SellerSignup{}	true	"Seller signup details"
// @Success		200		{object}	response.Response{}
// @Failure		400		{object}	response.Response{}
// @Router			/seller/signup [post]
func (u *SellerHandler) SellerSignup(c *gin.Context) {
	var sellerDetails requestmodel.SellerSignup

	if err := c.BindJSON(&sellerDetails); err != nil {
		c.JSON(http.StatusBadRequest, "can't bind json with struct")
	}

	data, err := helper.Validation(sellerDetails)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", data, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
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

// @Summary		Seller Login
// @Description	using this handler Seller can Login
// @Tags			Seller
// @Accept			json
// @Produce		json
// @Param			Seller	body		requestmodel.SellerLogin{}	true	"Seller Login details"
// @Success		200		{object}	response.Response{}
// @Failure		400		{object}	response.Response{}
// @Router			/seller/login [post]
func (u *SellerHandler) SellerLogin(c *gin.Context) {
	var loginData requestmodel.SellerLogin
	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	data, err := helper.Validation(loginData)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", data, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
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

// @Summary		Get Sellers
// @Description	Using this handler, admin can get a list of sellers
// @Tags			Admins
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Param			page	query		int	true	"Page number for pagination (default 1)"
// @Param			limit	query		int	true	"Number of items to return per page (default 5)"
// @Success		200		{object}	response.Response{}
// @Failure		400		{object}	response.Response{}
// @Router			/admin/sellers/getsellers [get]
func (u *SellerHandler) GetSellers(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "1")

	sellers, count, err := u.usecase.GetAllSellers(page, limit)
	if err != nil {
		// message := fmt.Sprintf("total sellers  %d", *count)
		// finalReslt := response.Responses(http.StatusNotFound, message, "", err.Error())
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		message := fmt.Sprintf("total sellers  %d", *count)
		finalReslt := response.Responses(http.StatusOK, message, sellers, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// @Summary		Block Seller
// @Description	Using this handler, admin can block a seller
// @Tags			Admins
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Param			id	path		string	true	"Seller ID in the URL path"
// @Success		200	{object}	response.Response{}
// @Failure		400	{object}	response.Response{}
// @Router			/admin/seller/block [patch]
func (u *SellerHandler) BlockSeller(c *gin.Context) {
	userID := c.Query("id")
	id := strings.TrimSpace(userID)

	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": resCustomError.IDParamsEmpty})
		return
	}

	err := u.usecase.BlockSeller(userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusNotFound, "", "", err.Error())
		c.JSON(http.StatusNotFound, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Succesfully block", "", nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// @Summary		Block Seller
// @Description	Using this handler, admin can block a seller
// @Tags			Admins
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Param			id	path		string	true	"Seller ID in the URL path"
// @Success		200	{object}	response.Response{}
// @Failure		400	{object}	response.Response{}
// @Router			/admin/sellers/block [patch]
func (u *SellerHandler) UnblockSeller(c *gin.Context) {
	userID := c.Query("id")
	id := strings.TrimSpace(userID)

	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": resCustomError.IDParamsEmpty})
		return
	}

	err := u.usecase.ActiveSeller(userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusNotFound, "", "", err.Error())
		c.JSON(http.StatusNotFound, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Succesfully unblock", "", nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// @Summary		Get Pending Sellers
// @Description	Using this handler, admin can get a list of pending sellers
// @Tags			Admins
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Param			page	query		int	true	"Page number for pagination (default 1)"
// @Param			limit	query		int	true	"Number of items to return per page (default 5)"
// @Success		200		{object}	response.Response{}
// @Failure		400		{object}	response.Response{}
// @Router			/admin/sellers/pending [get]
func (u *SellerHandler) GetPendingSellers(c *gin.Context) {
	page := c.Query("page")
	limit := c.DefaultQuery("limit", "1")

	sellers, err := u.usecase.GetAllPendingSellers(page, limit)
	if err != nil {
		// finalReslt := response.Responses(http.StatusNotFound, "", "", err.Error())
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		finalReslt := response.Responses(http.StatusOK, "", sellers, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// @Summary		Get Single Seller Details
// @Description	Using this handler, admin can get details of a single seller
// @Tags			Admins
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Param			id	query		string	true	"Seller ID in the URL query"
// @Success		200	{object}	response.Response{}
// @Failure		400	{object}	response.Response{}
// @Router			/admin/sellers/singleview [get]
func (u *SellerHandler) FetchSingleSeller(c *gin.Context) {
	userID := c.Query("id")
	id := strings.TrimSpace(userID)

	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": resCustomError.IDParamsEmpty})
		return
	}

	sellerDetails, err := u.usecase.FetchSingleVender(userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusNotFound, "", "", err.Error())
		c.JSON(http.StatusNotFound, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", sellerDetails, nil)
		c.JSON(http.StatusOK, finalReslt)
	}

}

// @Summary		Verify Seller
// @Description	Using this handler, admin can Verify a seller
// @Tags			Admins
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Param			id	path		string	true	"Seller ID in the URL path"
// @Success		200	{object}	response.Response{}
// @Failure		400	{object}	response.Response{}
// @Router			/admin/sellers/verify [patch]
func (u *SellerHandler) VerifySeller(c *gin.Context) {
	userID := c.Query("id")
	id := strings.TrimSpace(userID)

	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": resCustomError.IDParamsEmpty})
		return
	}

	err := u.usecase.ActiveSeller(userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusNotFound, "", "", err.Error())
		c.JSON(http.StatusNotFound, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Verification Success", "", nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// ------------------------------------------Seller Profile------------------------------------\\

func (u *SellerHandler) GetSellerProfile(c *gin.Context) {

	userID, exist := c.MustGet("SellerID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	sellerProfile, err := u.usecase.GetSellerProfile(userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", sellerProfile, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *SellerHandler) EditSellerProfile(c *gin.Context) {

	var profile models.SellerEditProfile

	sellerID, exist := c.MustGet("SellerID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	profile.ID = sellerID

	if err := c.ShouldBind(&profile); err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, resCustomError.BindingConflict, nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	userProfile, err := u.usecase.UpdateSellerProfile(&profile)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Succesfully Edited", userProfile, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}
