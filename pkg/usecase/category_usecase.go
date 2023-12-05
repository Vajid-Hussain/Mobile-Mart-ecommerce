package usecase

import (
	"strconv"

	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
	resCustomError "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel/custom_error"
	interfaces "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/repository/interface"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/utils/helper"
)

type categoryUseCase struct {
	repo interfaces.ICategoryRepository
}

func NewCategoryUseCase(repository interfaces.ICategoryRepository) interfaceUseCase.ICategoryUseCase {
	return &categoryUseCase{repo: repository}
}

func (r *categoryUseCase) NewCategory(categoryDetails *requestmodel.Category) (*responsemodel.Category, error) {

	err := r.repo.InsertCategory(categoryDetails)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *categoryUseCase) GetAllCategory(page string, limit string) (*[]responsemodel.CategoryDetails, error) {

	offSet, limits, err := helper.Pagination(page, limit)
	if err != nil {
		return nil, err
	}

	categoryDetails, err := r.repo.GetAllCategory(offSet, limits)
	if err != nil {
		return nil, err
	}

	return categoryDetails, nil
}

func (r *categoryUseCase) EditCategory(categoryData *requestmodel.CategoryDetails) (*responsemodel.CategoryDetails, error) {

	category, err := r.repo.EditCategoryName(categoryData)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (r *categoryUseCase) DeleteCategory(id string) error {

	ID, err := strconv.Atoi(id)
	if err != nil {
		return resCustomError.ErrNegativeID
	}

	if ID < 1 {
		return resCustomError.ErrNegativeID
	}

	err = r.repo.DeleteInventoryOfCategory(id)
	if err != nil {
		return err
	}

	err = r.repo.DeleteCategory(id)
	if err != nil {
		return err
	}
	return nil
}

// Brand
func (r *categoryUseCase) CreateBrand(brandDetails *requestmodel.Brand) (*responsemodel.BrandRes, error) {
	// var resBrand responsemodel.BrandRes

	// validate := validator.New(validator.WithRequiredStructEnabled())
	// err := validate.Struct(brandDetails)
	// if err != nil {
	// 	if ve, ok := err.(validator.ValidationErrors); ok {
	// 		for _, e := range ve {
	// 			switch e.Field() {
	// 			case "Name":
	// 				resBrand.Name = "name is medetary"
	// 			}
	// 		}
	// 	}
	// 	return &resBrand, errors.New("don't fullfill the brand requirement ")
	// }

	err := r.repo.InsertBrand(brandDetails)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *categoryUseCase) GetAllBrand(page string, limit string) (*[]responsemodel.BrandRes, error) {

	offSet, limits, err := helper.Pagination(page, limit)
	if err != nil {
		return nil, err
	}

	brandDetails, err := r.repo.GetAllBrand(offSet, limits)
	if err != nil {
		return nil, err
	}

	return brandDetails, nil
}

func (r *categoryUseCase) EditBrand(brandData *requestmodel.BrandDetails) (*responsemodel.BrandRes, error) {
	// var brandRes responsemodel.BrandRes

	// brandData.ID = strings.TrimSpace(brandData.ID)

	// validate := validator.New(validator.WithRequiredStructEnabled())
	// err := validate.Struct(brandData)
	// if err != nil {
	// 	if ve, ok := err.(validator.ValidationErrors); ok {
	// 		for _, e := range ve {
	// 			switch e.Field() {
	// 			case "Name":
	// 				brandRes.Name = "name is medetary"
	// 			case "ID":
	// 				brandRes.ID = "id is required,as query"
	// 			}
	// 		}
	// 	}
	// 	return &brandRes, errors.New("don't fullfill the brand requirement ")
	// }

	err := r.repo.EditBrandName(brandData)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *categoryUseCase) DeleteBrand(id string) error {

	ID, err := strconv.Atoi(id)
	if err != nil {
		return resCustomError.ErrNegativeID
	}

	if ID < 1 {
		return resCustomError.ErrNegativeID
	}

	err = r.repo.DeleteInventoryOfBrand(id)
	if err != nil {
		return err
	}

	err = r.repo.DeleteBrand(id)
	if err != nil {
		return err
	}
	return nil
}
