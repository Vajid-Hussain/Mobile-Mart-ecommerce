package usecase

import (
	"errors"
	"strconv"

	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
	resCustomError "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel/custom_error"
	interfaces "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/repository/interface"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
	"github.com/go-playground/validator/v10"
)

type categoryUseCase struct {
	repo interfaces.ICategoryRepository
}

func NewCategoryUseCase(repository interfaces.ICategoryRepository) interfaceUseCase.ICategoryUseCase {
	return &categoryUseCase{repo: repository}
}

func (r *categoryUseCase) NewCategory(categoryDetails *requestmodel.Category) (*responsemodel.Category, error) {
	var resCategory responsemodel.Category

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(categoryDetails)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "Name":
					resCategory.Name = "name is medetary"
				}
			}
		}
		return &resCategory, errors.New("don't fullfill the category requirement ")
	}

	err = r.repo.InsertCategory(categoryDetails)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *categoryUseCase) GetAllCategory(page string, limit string) (*[]responsemodel.CategoryDetails, error) {

	pageNO, err := strconv.Atoi(page)
	if err != nil {
		return nil, resCustomError.ConversionOFPageErr
	}

	limits, err := strconv.Atoi(limit)
	if err != nil {
		return nil, resCustomError.ConversionOfLimitErr
	}

	if pageNO < 1 {
		return nil, resCustomError.PaginationError
	}

	if limits <= 0 {
		return nil, resCustomError.PageLimitError
	}

	offSet := (pageNO * limits) - limits
	limits = pageNO * limits

	categoryDetails, err := r.repo.GetAllCategory(offSet, limits)
	if err != nil {
		return nil, err
	}

	return categoryDetails, nil
}

// func (r *categoryUseCase) EditCategory(categoryData requestmodel.Category) error{
// 	validate := validator.New(validator.WithRequiredStructEnabled())
// 	err := validate.Struct(categoryData)
// 	if err != nil {
// 		if ve, ok := err.(validator.ValidationErrors); ok {
// 			for _, e := range ve {
// 				switch e.Field() {
// 				case "Name":
// 					resCategory.Name = "name is medetary"
// 				}
// 			}
// 		}
// 		return &resCategory, errors.New("don't fullfill the category requirement ")
// 	}
// }
