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
	UpdateUserOrderCancel(string, string) (*responsemodel.OrderDetails, error)
	UpdateDeliveryTimeByUser(string, string) error
	GetOrderExistOfUser(string, string) error
	GetAddressExist(string, string) error

	GetSellerOrders(string, string) (*[]responsemodel.OrderDetails, error)
	UpdateOrderDelivered(string, string) (*responsemodel.OrderDetails, error)
	UpdateDeliveryTime(string, string) error
	UpdateOrderCancel(string, string) (*responsemodel.OrderDetails, error)
	UpdateOrderPaymetSuccess(string, string) error
	GetOrderExistOfSeller(string, string) error

	GetSalesReportByYear(string, string) (*responsemodel.SalesReport, error)
	GetSalesReportByDays(string, string) (*responsemodel.SalesReport, error)

	OnlinePayment(string) (*responsemodel.OnlinePayment, error)
	GetFinalPriceByorderID(string) (uint, error)
	UpdateOnlinePaymentSucess(string) (*[]responsemodel.OrderDetails, error)
}
