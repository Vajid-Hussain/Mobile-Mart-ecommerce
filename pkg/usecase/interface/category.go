package interfaceUseCase

import (
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
)

type ICategoryUseCase interface {
	NewCategory(*requestmodel.Category) (*responsemodel.Category, error)
	GetAllCategory(string, string) (*[]responsemodel.CategoryDetails, error)
	EditCategory(*requestmodel.CategoryDetails) (*responsemodel.CategoryDetails, error)
	DeleteCategory(string) error

	CreateBrand(*requestmodel.Brand) (*responsemodel.BrandRes, error)
	GetAllBrand(string, string) (*[]responsemodel.BrandRes, error)
	EditBrand(*requestmodel.BrandDetails) (*responsemodel.BrandRes, error)
	DeleteBrand(string) error

	CategoryOffer(*requestmodel.CategoryOffer) (*responsemodel.CategoryOffer, error)
	ChangeStatusOfCategoryOffer(string, string) (*responsemodel.CategoryOffer, error)
	GetAllCategoryOffer(string) (*[]responsemodel.CategoryOffer, error)
	UpdateCategoryOffer(*requestmodel.EditCategoryOffer) (*responsemodel.CategoryOffer, error)
}
