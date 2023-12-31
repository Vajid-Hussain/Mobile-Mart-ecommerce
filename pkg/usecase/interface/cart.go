package interfaceUseCase

import (
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
)

type ICartUseCase interface {
	CreateCart(*requestmodel.Cart) (*requestmodel.Cart, error)
	DeleteInventoryFromCart(string, string) error
	QuantityIncriment(string, string) (*requestmodel.Cart, error)
	QuantityDecrease(string, string) (*requestmodel.Cart, error)
	ShowCart(string) (*responsemodel.UserCart, error)
}
