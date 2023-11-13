package handler

import (
	"net/http"

	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel/response"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryUseCase interfaceUseCase.ICategoryUseCase
}

func NewCategoryHandler(useCase interfaceUseCase.ICategoryUseCase) *CategoryHandler {
	return &CategoryHandler{categoryUseCase: useCase}
}

func (u *CategoryHandler) NewCategory(c *gin.Context) {

	var categoryDetails requestmodel.Category

	err := c.BindJSON(&categoryDetails)
	if err != nil {
		c.JSON(http.StatusBadRequest, "can't bind json with struct")
	}

	result, err := u.categoryUseCase.NewCategory(&categoryDetails)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", result, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Category succesfully added", result, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *CategoryHandler) FetchAllCatogry(c *gin.Context) {
	page := c.Query("page")
	limit := c.DefaultQuery("limit", "1")

	category, err := u.categoryUseCase.GetAllCategory(page, limit)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		finalReslt := response.Responses(http.StatusOK, "", category, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}
