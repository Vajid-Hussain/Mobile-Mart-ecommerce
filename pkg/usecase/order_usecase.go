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

type orderUseCase struct {
	repo             interfaces.IOrderRepository
	cartrepo         interfaces.ICartRepository
	sellerRepository interfaces.ISellerRepo
	razopay          *config.Razopay
}

func NewOrderUseCase(repository interfaces.IOrderRepository, cartrepository interfaces.ICartRepository, sellerRepository interfaces.ISellerRepo, razopay *config.Razopay) interfaceUseCase.IOrderUseCase {
	return &orderUseCase{repo: repository, cartrepo: cartrepository, sellerRepository: sellerRepository, razopay: razopay}
}

func (r *orderUseCase) NewOrder(order *requestmodel.Order) (*responsemodel.Order, error) {

	if order.Payment == "COD" {
		order.OrderStatus = "processing"
	} else {
		order.OrderStatus = "pending"
	}

	err := r.repo.GetAddressExist(order.UserID, order.Address)
	if err != nil {
		return nil, err
	}

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
			return nil, fmt.Errorf("sorry for inconvinent for less stock , we have only %d units, your requirement is %d unit,of product id %s", *unit, data.Quantity, data.InventoryID)
		}

		newUnit := *unit - data.Quantity
		err = r.repo.UpdateInventoryUnits(data.InventoryID, newUnit)
		if err != nil {
			return nil, err
		}
	}

	// find total amount
	for i, product := range order.Cart {
		inventotyPrice, err := r.cartrepo.GetInventoryPrice(product.InventoryID)
		if err != nil {
			return nil, err
		}
		order.Cart[i].Price = inventotyPrice * product.Quantity
		order.FinalPrice += order.Cart[i].Price
	}

	// place order on payment is online
	if order.Payment == "ONLINE" {
		orderID, err := service.Razopay(order.FinalPrice, r.razopay.RazopayKey, r.razopay.RazopaySecret)
		if err != nil {
			return nil, err
		}
		order.OrderID = orderID
	}

	orderResponse, err := r.repo.CreateOrder(order)
	if err != nil {
		return nil, err
	}

	OrderSuccessDetails, err := r.repo.AddProdutToOrderProductTable(order, orderResponse)
	if err != nil {
		return nil, err
	}
	// for _, data := range order.Cart {
	// 	err = r.cartrepo.DeleteInventoryFromCart(data.InventoryID, order.UserID)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	// orderResponse.UserID = order.UserID
	// orderResponse.Address = order.Address
	// orderResponse.Payment = order.Payment
	return OrderSuccessDetails, nil
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

func (r *orderUseCase) CancelUserOrder(orderID string, userID string) (*responsemodel.OrderDetails, error) {
	fmt.Println("**", orderID, userID)
	err := r.repo.GetOrderExistOfUser(orderID, userID)
	if err != nil {
		return nil, err
	}

	orderDetails, err := r.repo.UpdateUserOrderCancel(orderID, userID)
	if err != nil {
		return nil, err
	}

	units, err := r.repo.GetInventoryUnits(orderDetails.InventoryID)
	if err != nil {
		return nil, err
	}

	err = r.repo.UpdateDeliveryTimeByUser(userID, orderID)
	if err != nil {
		return nil, err
	}

	updatedUnit := *units + orderDetails.Quantity

	err = r.repo.UpdateInventoryUnits(orderDetails.InventoryID, updatedUnit)
	if err != nil {
		return nil, err
	}
	return orderDetails, nil
}

// ------------------------------------------Online Payment------------------------------------\\

func (r *orderUseCase) OnlinePayment(userID string) (*responsemodel.OnlinePayment, error) {
	paymentDetails, err := r.repo.OnlinePayment(userID)
	if err != nil {
		return nil, err
	}

	paymentDetails.FinalPrice, err = r.repo.GetFinalPriceByorderID(paymentDetails.OrderID)
	if err != nil {
		return nil, err
	}
	return paymentDetails, nil
}

func (r *orderUseCase) OnlinePaymentVerification(details *requestmodel.OnlinePaymentVerification) (*[]responsemodel.OrderDetails, error) {
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

	err = r.repo.UpdateOrderPaymetSuccess(sellerID, orderID)
	if err != nil {
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

func (r *orderUseCase) CancelOrder(orderID string, sellerID string) (*responsemodel.OrderDetails, error) {
	err := r.repo.GetOrderExistOfSeller(orderID, sellerID)
	if err != nil {
		return nil, err
	}
	orderDetails, err := r.repo.UpdateOrderCancel(orderID, sellerID)
	if err != nil {
		return nil, err
	}

	units, err := r.repo.GetInventoryUnits(orderDetails.InventoryID)
	if err != nil {
		return nil, err
	}

	err = r.repo.UpdateDeliveryTime(sellerID, orderID)
	if err != nil {
		return nil, err
	}

	updatedUnit := *units + orderDetails.Quantity

	err = r.repo.UpdateInventoryUnits(orderDetails.InventoryID, updatedUnit)
	if err != nil {
		return nil, err
	}

	return orderDetails, nil
}

// ------------------------------------------Seller Sales Report------------------------------------\\

func (r *orderUseCase) GetSalesReportByYear(sellerID string, year string) (*responsemodel.SalesReport, error) {
	report, err := r.repo.GetSalesReportByYear(sellerID, year)
	if err != nil {
		return nil, err
	}
	return report, nil
}

func (r *orderUseCase) GetSalesReportByDays(sellerID string, days string) (*responsemodel.SalesReport, error) {
	report, err := r.repo.GetSalesReportByDays(sellerID, days)
	if err != nil {
		return nil, err
	}
	return report, nil
}
