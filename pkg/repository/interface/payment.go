package interfaces

import responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"

type IPaymentRepository interface {
	CreateOrUpdateWallet(string, uint) (*uint, error)
	OnlinePayment(string, string) (*responsemodel.OnlinePayment, error)
	GetFinalPriceByorderID(string) (uint, error)
	UpdateOnlinePaymentSucess(string) (*[]responsemodel.OrderDetails, error)
}
