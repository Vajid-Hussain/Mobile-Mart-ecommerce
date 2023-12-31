package usecase

import (
	"errors"
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

func (r *categoryUseCase) CategoryOffer(categoryOffer *requestmodel.CategoryOffer) (*responsemodel.CategoryOffer, error) {
	categoryCount, err := r.repo.ChekSellerHaveCategoryOffer(categoryOffer.SellerID, categoryOffer.CategoryID)
	if err != nil {
		return nil, err
	}
	if *categoryCount > 0 {
		return nil, errors.New("the offer is currently live in the same category. Now, you can edit the category offer")
	}

	categoryOfferRes, err := r.repo.InsertCategoryOffer(categoryOffer)
	if err != nil {
		return nil, err
	}
	return categoryOfferRes, nil
}

func (r *categoryUseCase) ChangeStatusOfCategoryOffer(status, categoryOfferID string) (*responsemodel.CategoryOffer, error) {
	offer, err := r.repo.ChangeStatus(status, categoryOfferID)
	if err != nil {
		return nil, err
	}
	return offer, nil
}

func (r *categoryUseCase) GetAllCategoryOffer(sellerID string) (*[]responsemodel.CategoryOffer, error) {
	offers, err := r.repo.GetAllCategoryOffers(sellerID)
	if err != nil {
		return nil, err
	}
	return offers, nil
}

func (r *categoryUseCase) UpdateCategoryOffer(updateData *requestmodel.EditCategoryOffer) (*responsemodel.CategoryOffer, error) {
	newCategoryOffer, err := r.repo.UpdateCategoryOffer(updateData)
	if err != nil {
		return nil, err
	}
	return newCategoryOffer, nil
}
