package interfaceUseCase

import (
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
	"github.com/jung-kurt/gofpdf"
)

type IOrderUseCase interface {
	NewOrder(*requestmodel.Order) (*responsemodel.Order, error)
	OrderShowcase(string) (*[]responsemodel.OrderShowcase, error)
	SingleOrder(string, string) (*responsemodel.SingleOrder, error)
	CancelUserOrder(string, string) (*responsemodel.OrderDetails, error)
	ReturnUserOrder(string, string) (*responsemodel.OrderDetails, error)

	GetSellerOrders(string, string) (*[]responsemodel.OrderDetails, error)
	ConfirmDeliverd(string, string) (*responsemodel.OrderDetails, error)
	CancelOrder(string, string) (*responsemodel.OrderDetails, error)

	GetSalesReportByYear(string, string) (*responsemodel.SalesReport, error)
	GetSalesReportByDays(string, string) (*responsemodel.SalesReport, error)

	OrderInvoiceCreation(string) (*gofpdf.Fpdf, error)
}
