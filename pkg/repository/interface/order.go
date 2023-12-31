package interfaces

import (
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
)

type IOrderRepository interface {
	CreateOrder(*requestmodel.Order) (*responsemodel.Order, error)
	GetOrderShowcase(string) (*[]responsemodel.OrderShowcase, error)
	GetSingleOrder(string, string) (*responsemodel.SingleOrder, error)
	GetInventoryUnits(string) (*uint, error)
	UpdateInventoryUnits(string, uint) error
	GetOrderPrice(string) (uint, error)
	UpdateUserOrderCancel(string, string) (*responsemodel.OrderDetails, error)
	GetPaymentType(string) (string, error)
	UpdateDeliveryTimeByUser(string, string) error
	GetOrderExistOfUser(string, string) error
	GetAddressExist(string, string) error
	AddProdutToOrderProductTable(*requestmodel.Order, *responsemodel.Order) (*responsemodel.Order, error)
	UpdateUserOrderReturn(string, string) (*responsemodel.OrderDetails, error)
	GetOrderFullDetails(string) (*responsemodel.Invoice, error)
	GetAddressForInvoice(string) (*requestmodel.Address, error)
	GetAInventoryForInvoice(id string) (*responsemodel.InventoryRes, error)
	GetOrderXlSalesReport(string) (*[]responsemodel.XlSalesReport, error)

	GetSellerOrders(string, string) (*[]responsemodel.OrderDetails, error)
	UpdateOrderDelivered(string, string) (*responsemodel.OrderDetails, error)
	UpdateDeliveryTime(string, string) error
	UpdateOrderCancel(string, string) (*responsemodel.OrderDetails, error)
	UpdateOrderPaymetSuccess(string, string) error
	GetOrderExistOfSeller(string, string) error
	CheckCouponAppliedOrNot(string, string) uint

	GetSalesReport(string, string, string, string) (*responsemodel.SalesReport, error)
	GetSalesReportByDays(string, string) (*responsemodel.SalesReport, error)

	GetCategoryOffers(string) uint
}
