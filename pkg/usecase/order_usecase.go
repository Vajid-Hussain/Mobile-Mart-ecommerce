package usecase

import (
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
	interfaces "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/repository/interface"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
)

type orderUseCase struct {
	repo     interfaces.IOrderRepository
	cartrepo interfaces.ICartRepository
}

func NewOrderUseCase(repository interfaces.IOrderRepository, cartrepository interfaces.ICartRepository) interfaceUseCase.IOrderUseCase {
	return &orderUseCase{repo: repository, cartrepo: cartrepository}
}

func (r *orderUseCase) NewOrder(order *requestmodel.Order) (*[]responsemodel.OrderSuccess, error) {

	userCart, err := r.cartrepo.GetCart(order.UserID)
	if err != nil {
		return nil, err
	}

	order.Cart = *userCart

	orderResponse, err := r.repo.CreateOrder(order)
	if err != nil {
		return nil, err
	}

	for _, data := range order.Cart {
		err = r.cartrepo.DeleteInventoryFromCart(data.InventoryID, order.UserID)
		if err != nil {
			return nil, err
		}
	}

	return orderResponse, nil
}

func (r *orderUseCase) OrderShowcase(userID string) (*[]responsemodel.OrderShowcase, error) {
	abstractOrder, err := r.repo.GetOrderShowcase(userID)
	if err != nil {
		return nil, err
	}
	return abstractOrder, nil
}
