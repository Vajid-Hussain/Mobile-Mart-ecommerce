package interfaces

import (
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
)

type IInventoryRepository interface {
	CreateProduct(*requestmodel.InventoryReq) (*responsemodel.InventoryRes, error)
	BlockSingleInventoryBySeller(string, string) error
	UNBlockSingleInventoryBySeller(string, string) error
	DeleteInventoryBySeller(string, string) error
	GetInventory(int, int) (*[]responsemodel.InventoryShowcase, error)
	GetAInventory(string) (*[]responsemodel.InventoryRes, error)
	GetSellerInventory(int, int, string) (*[]responsemodel.InventoryShowcase, error)
}
