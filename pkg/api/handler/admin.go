package handler

import (
	"net/http"

	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel/response"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/utils/helper"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	AdminUseCase interfaceUseCase.IAdminUseCase
}

func NewAdminHandler(useCase interfaceUseCase.IAdminUseCase) *AdminHandler {
	return &AdminHandler{AdminUseCase: useCase}
}

//	@Summary		Admin login
//	@Description	using this handler admins can login
//	@Tags			Admins
//	@Accept			json
//	@Produce		json
//	@Param			admin	body		requestmodel.AdminLoginData	true	"Admin login details"
//	@Success		200		{object}	response.Response{}
//	@Failure		400		{object}	response.Response{}
//	@Router			/admin/login/ [post]
func (u *AdminHandler) AdminLogin(c *gin.Context) {
	var loginCredential requestmodel.AdminLoginData

	err := c.BindJSON(&loginCredential)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "json is wrong can't bind", nil, err)
		c.JSON(http.StatusUnauthorized, finalReslt)
		return
	}

	data, err := helper.Validation(loginCredential)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", data, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	result, err := u.AdminUseCase.AdminLogin(&loginCredential)
	if err != nil {
		finalReslt := response.Responses(http.StatusUnauthorized, "", result, err)
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "succesfully login", result, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

//	@Summary		Get Admin Details
//	@Description	Retrieve details for the admin.
//	@Tags			Admins
//	@Accept			json
//	@Produce		json
//	@Security		BearerTokenAuth
//	@Success		200	{object}	response.Response	"Admin details retrieved successfully"
//	@Failure		401	{object}	response.Response	"Unauthorized. Authentication required."
//	@Router			/admin [get]
func (u *AdminHandler) AdminDashBord(c *gin.Context) {
	result, err := u.AdminUseCase.GetSellerDetailsForAdminDashBord()
	if err != nil {
		finalReslt := response.Responses(http.StatusUnauthorized, "", result, err)
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "succesfully login", result, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}
