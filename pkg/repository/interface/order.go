package interfaces

import (
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
)

type IOrderRepository interface {
	CreateOrder(*requestmodel.Order) (*[]responsemodel.OrderSuccess, error)
	GetOrderShowcase(string) (*[]responsemodel.OrderShowcase, error)
}
