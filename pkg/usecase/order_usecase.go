package usecase

import (
	"fmt"

	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
	interfaces "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/repository/interface"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
)

type orderUseCase struct {
	repo             interfaces.IOrderRepository
	cartrepo         interfaces.ICartRepository
	sellerRepository interfaces.ISellerRepo
}

func NewOrderUseCase(repository interfaces.IOrderRepository, cartrepository interfaces.ICartRepository, sellerRepository interfaces.ISellerRepo) interfaceUseCase.IOrderUseCase {
	return &orderUseCase{repo: repository, cartrepo: cartrepository, sellerRepository: sellerRepository}
}

func (r *orderUseCase) NewOrder(order *requestmodel.Order) (*responsemodel.OrderSuccess, error) {

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

	orderResponse.UserID = order.UserID
	orderResponse.Address = order.Address
	orderResponse.Payment = order.Payment
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

// ------------------------------------------Seller Control Orders------------------------------------\\

func (r *orderUseCase) GetSellerOrders(sellerID string, remainingQuery string) (*[]responsemodel.OrderDetails, error) {
	userOrders, err := r.repo.GetSellerOrders(sellerID, remainingQuery)
	if err != nil {
		return nil, err
	}
	return userOrders, nil
}

func (r *orderUseCase) ConfirmDeliverd(sellerID string, orderID string) (*responsemodel.OrderDetails, error) {

	err := r.repo.UpdateDeliveryTime(sellerID, orderID)
	if err != nil {
		return nil, err
	}

	orderDetails, err := r.repo.UpdateOrderDelivered(sellerID, orderID)
	if err != nil {
		fmt.Println("order", orderDetails)
		return nil, err
	}

	orderPrice, err := r.repo.GetOrderPrice(orderID)
	if err != nil {
		return nil, err
	}

	sellerCredit, err := r.sellerRepository.GetSellerCredit(sellerID)
	if err != nil {
		return nil, err
	}

	newSellerCredit := sellerCredit + orderPrice
	fmt.Println("$", newSellerCredit)

	err = r.sellerRepository.UpdateSellerCredit(sellerID, newSellerCredit)
	if err != nil {
		return nil, err
	}

	return orderDetails, nil
}
