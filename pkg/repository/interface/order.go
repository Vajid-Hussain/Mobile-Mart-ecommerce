package interfaces

import (
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
)

type IOrderRepository interface {
	CreateOrder(*requestmodel.Order) (*responsemodel.OrderSuccess, error)
	GetOrderShowcase(string) (*[]responsemodel.OrderShowcase, error)
	GetSingleOrder(string, string) (*responsemodel.SingleOrder, error)
	GetInventoryUnits(string) (*uint, error)
	UpdateInventoryUnits(string, uint) error
	GetOrderPrice(string) (uint, error)

	GetSellerOrders(string, string) (*[]responsemodel.OrderDetails, error)
	UpdateOrderDelivered(string, string) (*responsemodel.OrderDetails, error)
	UpdateDeliveryTime(string, string) error
}
