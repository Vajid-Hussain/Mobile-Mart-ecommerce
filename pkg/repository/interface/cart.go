package interfaces

import requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"

type ICartRepository interface {
	InsertToCart(*requestmodel.Cart) (*requestmodel.Cart, error)
	GetInventoryPrice(string) (uint, error)
	IsInventoryExistInCart(string, string) (int, error)
	DeleteInventoryFromCart(string, string) error
	GetSingleInverntory(string, string) (*requestmodel.Cart, error)
	UpdateQuantityAndPrice(*requestmodel.Cart) (*requestmodel.Cart, error)
}
