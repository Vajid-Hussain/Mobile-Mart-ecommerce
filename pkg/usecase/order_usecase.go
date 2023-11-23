package usecase

import (
	"fmt"

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

	for _, data := range order.Cart {
		unit, err := r.repo.GetInventoryUnits(data.InventoryID)
		if err != nil {
			return nil, err
		}

		if *unit < data.Quantity {
			return nil, fmt.Errorf("sorry for inconvinent for less stock , we have only %d units, your requirement is %d unit", *unit, data.Quantity)
		}

		newUnit := *unit - data.Quantity
		err = r.repo.UpdateInventoryUnits(data.InventoryID, newUnit)
		if err != nil {
			return nil, err
		}
	}

	for i, product := range order.Cart {
		inventotyPrice, err := r.cartrepo.GetInventoryPrice(product.InventoryID)
		if err != nil {
			return nil, err
		}
		order.Cart[i].Price = inventotyPrice * product.Quantity
	}

	fmt.Println("@@@@@@@@@@", order.Cart)

	orderResponse, err := r.repo.CreateOrder(order)
	if err != nil {
		return nil, err
	}

	// for _, data := range order.Cart {
	// 	err = r.cartrepo.DeleteInventoryFromCart(data.InventoryID, order.UserID)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	return orderResponse, nil
}

func (r *orderUseCase) OrderShowcase(userID string) (*[]responsemodel.OrderShowcase, error) {
	abstractOrder, err := r.repo.GetOrderShowcase(userID)
	if err != nil {
		return nil, err
	}
	return abstractOrder, nil
}

func (r *orderUseCase) SingleOrder(orderID string, userID string) (*responsemodel.SingleOrder, error) {
	singleOrder, err := r.repo.GetSingleOrder(orderID, userID)
	if err != nil {
		return nil, err
	}
	return singleOrder, nil
}
