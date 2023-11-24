package interfaces

import (
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
)

type ICartRepository interface {
	InsertToCart(*requestmodel.Cart) (*requestmodel.Cart, error)
	GetInventoryPrice(string) (uint, error)
	IsInventoryExistInCart(string, string) (int, error)
	DeleteInventoryFromCart(string, string) error
	GetSingleInverntory(string, string) (*requestmodel.Cart, error)
	UpdateQuantity(*requestmodel.Cart) (*requestmodel.Cart, error)
	GetCart(string) (*[]responsemodel.CartInventory, error)
	GetNetAmoutOfCart(string, string) (uint, error)
	GetCartCriteria(string) (uint, error)
}
