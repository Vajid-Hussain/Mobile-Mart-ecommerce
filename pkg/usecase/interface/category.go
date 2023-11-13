package interfaceUseCase

import (
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
)

type ICategoryUseCase interface {
	NewCategory(*requestmodel.Category) (*responsemodel.Category, error)
	GetAllCategory(string, string) (*[]responsemodel.CategoryDetails, error)
}
