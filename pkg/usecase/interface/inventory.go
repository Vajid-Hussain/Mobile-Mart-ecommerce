package interfaceUseCase

import (
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
)

type IInventoryUseCase interface {
	AddInventory(*requestmodel.InventoryReq) (*[]responsemodel.Errors, *responsemodel.InventoryRes, error)
}
