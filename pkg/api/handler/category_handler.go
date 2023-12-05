package handler

import (
	"net/http"

	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	resCustomError "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel/custom_error"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel/response"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/utils/helper"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryUseCase interfaceUseCase.ICategoryUseCase
}

func NewCategoryHandler(useCase interfaceUseCase.ICategoryUseCase) *CategoryHandler {
	return &CategoryHandler{categoryUseCase: useCase}
}

// @Summary		Add Category
// @Description	Using this handler, admin can add a new category
// @Tags			Category
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Param			name	query		string	true	"Name of the category"
// @Success		200		{object}	response.Response{}
// @Failure		400		{object}	response.Response{}
// @Router			/admin/category [post]
func (u *CategoryHandler) NewCategory(c *gin.Context) {

	var categoryDetails requestmodel.Category

	err := c.BindJSON(&categoryDetails)
	if err != nil {
		c.JSON(http.StatusBadRequest, resCustomError.BindingConflict)
	}

	data, err := helper.Validation(categoryDetails)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", data, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
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

// @Summary		Get All Categories
// @Description	Using this handler, admin can get a list of all categories
// @Tags			Category
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Param			page	query		int	true	"Page number for pagination (default 1)"
// @Param			limit	query		int	true	"Number of items to return per page (default 5)"
// @Success		200		{object}	response.Response{}
// @Failure		400		{object}	response.Response{}
// @Router			/admin/category [get]
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

// @Summary		Edit a Category by ID
// @Description	Edit an existing category using this handler.
// @Tags			Category
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Param			id			path		int					true	"ID of the category to edit"
// @Param			name		formData	string				true	"Updated name of the category"
// @Param			description	formData	string				false	"Updated description of the category"
// @Success		200			{object}	response.Response{}	"Category edited successfully"
// @Failure		400			{object}	response.Response{}	"Invalid input or validation error"
// @Failure		404			{object}	response.Response{}	"Category not found"
// @Router			/admin/category/{id} [patch]
func (u *CategoryHandler) UpdateCategory(c *gin.Context) {
	var categoryData requestmodel.CategoryDetails

	if err := c.BindJSON(&categoryData); err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, resCustomError.BindingConflict, nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	data, err := helper.Validation(categoryData)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", data, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	// categoryData.ID = c.Query("id")

	categoryRes, err := u.categoryUseCase.EditCategory(&categoryData)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "refine request", categoryRes, nil)
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "succesfully acomplish", categoryRes, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// @Summary		Delete a Category by ID
// @Description	Delete an existing category using this handler.
// @Tags			Category
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Param			id	query	int	true	"ID of the category to delete"
// @Success		204	"Category deleted successfully"
// @Failure		400	{object}	response.Response{}	"Invalid input or validation error"
// @Router			/admin/category [delete]
func (u *CategoryHandler) DeleteCategory(c *gin.Context) {

	id := c.Query("id")

	err := u.categoryUseCase.DeleteCategory(id)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "succesfully category deleted", nil, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// Brand

// @Summary		Create a Brand
// @Description	Create a new brand using this handler.
// @Tags			Brand
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Param			name		formData	string				true	"Name of the brand"
// @Param			description	formData	string				false	"Description of the brand"
// @Success		201			{object}	response.Response{}	"Brand created successfully"
// @Failure		400			{object}	response.Response{}	"Invalid input or validation error"
// @Router			/admin/brand [post]
func (u *CategoryHandler) CreateBrand(c *gin.Context) {
	var BrandDetails requestmodel.Brand

	err := c.ShouldBindJSON(&BrandDetails)
	if err != nil {
		c.JSON(http.StatusBadRequest, resCustomError.BindingConflict)
		return
	}

	data, err := helper.Validation(BrandDetails)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", data, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	result, err := u.categoryUseCase.CreateBrand(&BrandDetails)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", result, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "Brand succesfully added", result, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// @Summary		Get Paginated List of Brands
// @Description	Get a paginated list of brands using this handler.
// @Tags			Brand
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Param			page	query		int					true	"Page number for pagination (default 1)"
// @Param			limit	query		int					true	"Number of items to return per page (default 5)"
// @Success		200		{object}	response.Response{}	"Paginated list of brands"
// @Failure		400		{object}	response.Response{}	"Invalid input or validation error"
// @Router			/admin/brand [get]
func (u *CategoryHandler) FetchAllBrand(c *gin.Context) {
	page := c.Query("page")
	limit := c.DefaultQuery("limit", "1")

	brand, err := u.categoryUseCase.GetAllBrand(page, limit)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		finalReslt := response.Responses(http.StatusOK, "", brand, nil)
		c.JSON(http.StatusOK, finalReslt)
	}

}

// @Summary		Edit a Brand by ID
// @Description	Edit an existing brand using this handler.
// @Tags			Brand
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Param			id			path		int					true	"ID of the brand to edit"
// @Param			name		formData	string				true	"Updated name of the brand"
// @Param			description	formData	string				false	"Updated description of the brand"
// @Success		200			{object}	response.Response{}	"Brand edited successfully"
// @Failure		400			{object}	response.Response{}	"Invalid input or validation error"
// @Failure		404			{object}	response.Response{}	"Brand not found"
// @Router			/admin/brand/{id} [patch]
func (u *CategoryHandler) UpdateBrand(c *gin.Context) {
	var brandData requestmodel.BrandDetails

	if err := c.BindJSON(&brandData); err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, resCustomError.BindingConflict, nil, nil)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	data, err := helper.Validation(brandData)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", data, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	brandData.ID = c.Query("id")

	brandRes, err := u.categoryUseCase.EditBrand(&brandData)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "refine request", brandRes, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "succesfully acomplish", brandRes, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// @Summary		Delete a Brand by ID
// @Description	Delete an existing brand using this handler.
// @Tags			Brand
// @Accept			json
// @Produce		json
// @Security		BearerTokenAuth
// @Param			id	path	int	true	"ID of the brand to delete"
// @Success		204	"Brand deleted successfully"
// @Failure		400	{object}	response.Response{}	"Invalid input or validation error"
// @Router			/admin/brand/{id} [delete]
func (u *CategoryHandler) DeleteBrand(c *gin.Context) {

	id := c.Query("id")

	err := u.categoryUseCase.DeleteBrand(id)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "succesfully Brand deleted", nil, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}
