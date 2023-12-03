package interfaceUseCase

import (
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
)

type IPaymentUseCase interface {
	OnlinePayment(string, string) (*responsemodel.OnlinePayment, error)
	OnlinePaymentVerification(*requestmodel.OnlinePaymentVerification) (*[]responsemodel.OrderDetails, error)
}
