package handler

import (
	"fmt"
	"net/http"
	"strings"

	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
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
		finalReslt := response.Responses(http.StatusBadRequest, "json is wrong can't bind", nil, err)
		c.JSON(http.StatusUnauthorized, finalReslt)
		return
	}

	result, err := u.AdminUseCase.AdminLogin(&loginCredential)
	if err != nil {
		finalReslt := response.Responses(http.StatusUnauthorized, "", result, err)
		c.JSON(http.StatusUnauthorized, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "succesfully login", result, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *AdminHandler) GetUser(c *gin.Context) {
	page := c.Query("page")
	limit := c.DefaultQuery("limit", "1")

	users, count, err := u.AdminUseCase.GetAllUsers(page, limit)
	if err != nil {
		message := fmt.Sprintf("total users  %d", *count)
		finalReslt := response.Responses(http.StatusNotFound, message, "", err.Error())
		c.JSON(http.StatusNotFound, finalReslt)
	} else {
		message := fmt.Sprintf("total users  %d", *count)
		finalReslt := response.Responses(http.StatusOK, message, users, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *AdminHandler) BlockUser(c *gin.Context) {
	userID := c.Query("id")
	id := strings.TrimSpace(userID)

	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": responsemodel.IDParamsEmpty})
		return
	}

	err := u.AdminUseCase.BlcokUser(userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusNotFound, "", "", err.Error())
		c.JSON(http.StatusNotFound, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Succesfully block", "", nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *AdminHandler) UnblockUser(c *gin.Context) {
	userID := c.Query("id")
	id := strings.TrimSpace(userID)

	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id parameter is empty"})
		return
	}

	err := u.AdminUseCase.UnblockUser(userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusNotFound, "", "", err.Error())
		c.JSON(http.StatusNotFound, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Succesfully unblock", "", nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

//Sellers

func (u *AdminHandler) GetSellers(c *gin.Context) {
	page := c.Query("page")
	limit := c.DefaultQuery("limit", "1")

	sellers, count, err := u.AdminUseCase.GetAllSellers(page, limit)
	if err != nil {
		message := fmt.Sprintf("total sellers  %d", *count)
		finalReslt := response.Responses(http.StatusNotFound, message, "", err.Error())
		c.JSON(http.StatusNotFound, finalReslt)
	} else {
		message := fmt.Sprintf("total sellers  %d", *count)
		finalReslt := response.Responses(http.StatusOK, message, sellers, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *AdminHandler) BlockSeller(c *gin.Context) {
	userID := c.Query("id")
	id := strings.TrimSpace(userID)

	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": responsemodel.IDParamsEmpty})
		return
	}

	err := u.AdminUseCase.BlockSeller(userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusNotFound, "", "", err.Error())
		c.JSON(http.StatusNotFound, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Succesfully block", "", nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *AdminHandler) UnblockSeller(c *gin.Context) {
	userID := c.Query("id")
	id := strings.TrimSpace(userID)

	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": responsemodel.IDParamsEmpty})
		return
	}

	err := u.AdminUseCase.UnblockSeller(userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusNotFound, "", "", err.Error())
		c.JSON(http.StatusNotFound, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Succesfully unblock", "", nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *AdminHandler) GetPendingSellers(c *gin.Context) {
	page := c.Query("page")
	limit := c.DefaultQuery("limit", "1")

	sellers, err := u.AdminUseCase.GetAllPendingSellers(page, limit)
	if err != nil {
		finalReslt := response.Responses(http.StatusNotFound, "", "", err.Error())
		c.JSON(http.StatusNotFound, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", sellers, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *AdminHandler) FetchSingleSeller(c *gin.Context) {
	userID := c.Query("id")
	id := strings.TrimSpace(userID)

	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": responsemodel.IDParamsEmpty})
		return
	}

	sellerDetails, err := u.AdminUseCase.FetchSingleVender(userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusNotFound, "", "", err.Error())
		c.JSON(http.StatusNotFound, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", sellerDetails, nil)
		c.JSON(http.StatusOK, finalReslt)
	}

}
