package usecase

import (
	"errors"
	"fmt"

	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/config"
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
	interfaces "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/repository/interface"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/service"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
)

type paymentUseCase struct {
	repo    interfaces.IPaymentRepository
	razopay *config.Razopay
}

func NewPaymentUseCase(repository interfaces.IPaymentRepository, razopay *config.Razopay) interfaceUseCase.IPaymentUseCase {
	return &paymentUseCase{repo: repository, razopay: razopay}
}

func (r *paymentUseCase) OnlinePayment(userID, orderID string) (*responsemodel.OnlinePayment, error) {
	paymentDetails, err := r.repo.OnlinePayment(userID, orderID)
	if err != nil {
		return nil, err
	}
	fmt.Println("&&", paymentDetails)
	paymentDetails.FinalPrice, err = r.repo.GetFinalPriceByorderID(orderID)
	if err != nil {
		return nil, err
	}
	return paymentDetails, nil
}

func (r *paymentUseCase) OnlinePaymentVerification(details *requestmodel.OnlinePaymentVerification) (*[]responsemodel.OrderDetails, error) {
	result := service.VerifyPayment(details.OrderID, details.PaymentID, details.Signature, r.razopay.RazopaySecret)
	if !result {
		return nil, errors.New("payment is unsuccessful")
	}

	orders, err := r.repo.UpdateOnlinePaymentSucess(details.OrderID)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *paymentUseCase) GetUserWallet(userID string) (*responsemodel.UserWallet, error) {
	userWallet, err := r.repo.GetWallet(userID)
	if err != nil {
		return nil, err
	}
	return userWallet, err
}
