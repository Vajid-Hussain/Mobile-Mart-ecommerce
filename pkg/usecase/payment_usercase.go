package usecase

import (
	interfaces "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/repository/interface"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
)

type paymentUseCase struct {
	repo interfaces.IPaymentRepository
}

func NewPaymentUseCase(repository interfaces.IPaymentRepository) interfaceUseCase.IPaymentUseCase {
	return &paymentUseCase{repo: repository}
}
