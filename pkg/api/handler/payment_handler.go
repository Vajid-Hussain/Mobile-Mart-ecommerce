package handler

import interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"

type PaymentHandler struct {
	useCase interfaceUseCase.IPaymentUseCase
}

func NewPaymentHandler(useCase interfaceUseCase.IPaymentUseCase) *PaymentHandler {
	return &PaymentHandler{useCase: useCase}
}
