package interfaces

import (
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
)

type IInventoryRepository interface {
	CreateProduct(*requestmodel.InventoryReq) (*responsemodel.InventoryRes, error)
}
