package handler

import (
	"fmt"
	"net/http"
	"strings"

	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	resCustomError "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel/custom_error"
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

func (u *SellerHandler) GetSellers(c *gin.Context) {
	page := c.Query("page")
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

func (u *SellerHandler) UnblockSeller(c *gin.Context) {
	userID := c.Query("id")
	id := strings.TrimSpace(userID)

	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": resCustomError.IDParamsEmpty})
		return
	}

	err := u.usecase.UnblockSeller(userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusNotFound, "", "", err.Error())
		c.JSON(http.StatusNotFound, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Succesfully unblock", "", nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

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
