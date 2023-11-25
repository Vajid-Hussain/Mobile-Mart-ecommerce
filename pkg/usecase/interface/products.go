package interfaceUseCase

import (
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
)

type IInventoryUseCase interface {
	AddInventory(*requestmodel.InventoryReq) (*responsemodel.InventoryRes, error)
	BlockInventory(string, string) error
	UNBlockInventory(string, string) error
	DeleteInventory(string, string) error
	GetAllInventory(string, string) (*[]responsemodel.InventoryShowcase, error)
	GetAInventory(string) (*responsemodel.InventoryRes, error)
	GetSellerInventory(string, string, string) (*[]responsemodel.InventoryShowcase, error)
	EditInventory(*requestmodel.EditInventory, string) (*responsemodel.InventoryRes, error)
}
