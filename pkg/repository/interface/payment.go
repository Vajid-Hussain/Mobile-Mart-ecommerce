package interfaces

import (
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
)

type IPaymentRepository interface {
	CreateOrUpdateWallet(string, uint) (uint, error)
	OnlinePayment(string, string) (*responsemodel.OnlinePayment, error)
	GetFinalPriceByorderID(string) (uint, error)
	UpdateOnlinePaymentSucess(string) (*[]responsemodel.OrderDetails, error)
	GetWallet(string) (*responsemodel.UserWallet, error)
	UpdateWalletReduceBalance(string, uint) error
	GetWalletbalance(userID string) (*uint, error)
	WalletTransaction(requestmodel.WalletTransaction) error
	GetWalletTransaction(string) (*[]responsemodel.WalletTransaction, error)
}
