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

// @Summary         Add Category
// @Description     Using this handler, admin can add a new category
// @Tags            Admins
// @Accept          json
// @Produce         json
// @Security        BearerTokenAuth
// @Param           name    query   string  true    "Name of the category"
// @Success         200     {object}    response.Response{}
// @Failure         400     {object}    response.Response{}
// @Router          /admin/category/add [post]
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

// @Summary         Get All Categories
// @Description     Using this handler, admin can get a list of all categories
// @Tags            Admins
// @Accept          json
// @Produce         json
// @Security        BearerTokenAuth
// @Param           page    query   int     true    "Page number for pagination (default 1)"
// @Param           limit   query   int     true    "Number of items to return per page (default 5)"
// @Success         200     {object}    response.Response{}
// @Failure         400     {object}    response.Response{}
// @Router          /admin/category/all [get]
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

func (u *CategoryHandler) UpdateCategory(c *gin.Context) {
	var categoryData requestmodel.CategoryDetails

	if err := c.BindJSON(&categoryData); err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "can't bind json with struct", nil, nil)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	categoryData.ID = c.Query("id")

	categoryRes, err := u.categoryUseCase.EditCategory(&categoryData)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "refine request", categoryRes, nil)
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "succesfully acomplish", categoryRes, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}
