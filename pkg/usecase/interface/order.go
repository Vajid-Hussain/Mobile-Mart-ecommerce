package interfaceUseCase

import (
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
)

type IOrderUseCase interface {
	NewOrder(*requestmodel.Order) (*responsemodel.OrderSuccess, error)
	OrderShowcase(string) (*[]responsemodel.OrderShowcase, error)
	SingleOrder(string, string) (*responsemodel.SingleOrder, error)

	GetSellerOrders(string, string) (*[]responsemodel.OrderDetails, error)
	ConfirmDeliverd(string, string) (*responsemodel.OrderDetails, error)
}
