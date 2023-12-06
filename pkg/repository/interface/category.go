package interfaces

import (
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
)

type ICategoryRepository interface {
	InsertCategory(*requestmodel.Category) error
	GetAllCategory(int, int) (*[]responsemodel.CategoryDetails, error)
	EditCategoryName(*requestmodel.CategoryDetails) (*responsemodel.CategoryDetails, error)
	DeleteCategory(string) error
	DeleteInventoryOfCategory(string) error

	InsertBrand(*requestmodel.Brand) error
	GetAllBrand(int, int) (*[]responsemodel.BrandRes, error)
	EditBrandName(*requestmodel.BrandDetails) error
	DeleteBrand(string) error
	DeleteInventoryOfBrand(string) error

	InsertCategoryOffer(*requestmodel.CategoryOffer) (*responsemodel.CategoryOffer, error)
	ChekSellerHaveCategoryOffer(string, string) (*uint, error)
	ChangeStatus(string, string, string, string) (*responsemodel.CategoryOffer, error)
	GetAllCategoryOffers(string) (*[]responsemodel.CategoryOffer, error)
	UpdateCategoryOffer(*requestmodel.EditCategoryOffer) (*responsemodel.CategoryOffer, error)
}
