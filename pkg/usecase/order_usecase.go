package usecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/config"
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
	interfaces "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/repository/interface"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/service"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/utils/helper"
)

type orderUseCase struct {
	repo             interfaces.IOrderRepository
	cartrepo         interfaces.ICartRepository
	sellerRepository interfaces.ISellerRepo
	paymentRepo      interfaces.IPaymentRepository
	couponrepo       interfaces.ICouponRepository
	razopay          *config.Razopay
}

func NewOrderUseCase(repository interfaces.IOrderRepository, cartrepository interfaces.ICartRepository, sellerRepository interfaces.ISellerRepo, paymentRepository interfaces.IPaymentRepository, coupon interfaces.ICouponRepository, razopay *config.Razopay) interfaceUseCase.IOrderUseCase {
	return &orderUseCase{repo: repository, cartrepo: cartrepository, sellerRepository: sellerRepository, paymentRepo: paymentRepository, couponrepo: coupon, razopay: razopay}
}

func (r *orderUseCase) NewOrder(order *requestmodel.Order) (*responsemodel.Order, error) {
	var couponData *responsemodel.Coupon

	if order.Payment == "COD" {
		order.OrderStatus = "processing"
		order.PaymentStatus = "pending"
	}
	if order.Payment == "ONLINE" {
		order.OrderStatus = "pending"
		order.PaymentStatus = "pending"
	}
	if order.Payment == "WALLET" {
		order.OrderStatus = "processing"
		order.PaymentStatus = "success"
	}

	err := r.repo.GetAddressExist(order.UserID, order.Address)
	if err != nil {
		return nil, err
	}

	// fetch products details from cart
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
			return nil, fmt.Errorf("sorry for inconvinent for insafishend stock , we have only %d units, your requirement is %d unit,of product id %s", *unit, data.Quantity, data.InventoryID)
		}

		newUnit := *unit - data.Quantity
		err = r.repo.UpdateInventoryUnits(data.InventoryID, newUnit)
		if err != nil {
			return nil, err
		}
	}

	for _, product := range order.Cart {
		inventotyPrice, err := r.cartrepo.GetInventoryPrice(product.InventoryID)
		if err != nil {
			return nil, err
		}
		order.FinalPrice += inventotyPrice
	}

	// verify coupon
	if order.Coupon != "" {
		couponData, err = r.couponrepo.CheckCouponExpired(order.Coupon)
		if err != nil {
			return nil, err
		}
		fmt.Println("##", order.FinalPrice)
		if order.FinalPrice < couponData.MinimumRequired || order.FinalPrice >= couponData.MaximumAllowed {
			return nil, errors.New("total price of order is not satisfying, for apply this coupon")
		}

		for i, data := range order.Cart {
			order.Cart[i].Price = helper.FindDiscount(float64(data.Price), float64(couponData.Discount))
		}

		order.FinalPrice = helper.FindDiscount(float64(order.FinalPrice), float64(couponData.Discount))

		rightNow := time.Now()
		if couponData.EndDate.Before(rightNow) {
			return nil, errors.New("coupon exeed the expiredata, better luck next time")
		}
	}

	// find total amount
	order.FinalPrice = 0
	for i, product := range order.Cart {
		inventotyPrice, err := r.cartrepo.GetInventoryPrice(product.InventoryID)
		if err != nil {
			return nil, err
		}

		discountedPrice := helper.FindDiscount(float64(inventotyPrice), float64(product.CategoryDiscount+couponData.Discount))

		order.Cart[i].Price = inventotyPrice * product.Quantity
		order.Cart[i].Discount = product.CategoryDiscount + couponData.Discount
		order.Cart[i].FinalPrice = discountedPrice
		fmt.Println("**", couponData.Discount, product.CategoryDiscount)
		order.FinalPrice += discountedPrice
	}

	// place order on payment is online
	if order.Payment == "ONLINE" {
		orderID, err := service.Razopay(order.FinalPrice, r.razopay.RazopayKey, r.razopay.RazopaySecret)
		if err != nil {
			return nil, err
		}
		order.OrderIDRazopay = orderID
	}

	// made payment using wallet
	if order.Payment == "WALLET" {
		userWallet, err := r.paymentRepo.GetWallet(order.UserID)
		if err != nil {
			return nil, err
		}

		if userWallet.Balance < order.FinalPrice {
			return nil, errors.New("no sufficient balance in the wallet")
		}

		err = r.paymentRepo.UpdateWalletReduceBalance(order.UserID, order.FinalPrice)
		if err != nil {
			return nil, err
		}
	}

	// order is creating
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

	orderResponse.TotalPrice = order.FinalPrice
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

func (r *orderUseCase) CancelUserOrder(orderItemID string, userID string) (*responsemodel.OrderDetails, error) {

	err := r.repo.GetOrderExistOfUser(orderItemID, userID)
	if err != nil {
		return nil, err
	}

	orderDetails, err := r.repo.UpdateUserOrderCancel(orderItemID, userID)
	if err != nil {
		return nil, err
	}

	paymentType, err := r.repo.GetPaymentType(orderItemID)
	if err != nil {
		return nil, err
	}

	if paymentType == "ONLINE" || paymentType == "WALLET" {
		orderDetails.WalletBalance, err = r.paymentRepo.CreateOrUpdateWallet(userID, orderDetails.Price)
		if err != nil {
			return nil, err
		}
	}

	units, err := r.repo.GetInventoryUnits(orderDetails.InventoryID)
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

func (r *orderUseCase) ReturnUserOrder(orderItemID, userID string) (*responsemodel.OrderDetails, error) {

	orderDetails, err := r.repo.UpdateUserOrderReturn(orderItemID, userID)
	if err != nil {
		return nil, err
	}

	orderDetails.WalletBalance, err = r.paymentRepo.CreateOrUpdateWallet(userID, orderDetails.Price)
	if err != nil {
		return nil, err
	}

	units, err := r.repo.GetInventoryUnits(orderDetails.InventoryID)
	if err != nil {
		return nil, err
	}

	updatedUnit := *units + orderDetails.Quantity

	err = r.repo.UpdateInventoryUnits(orderDetails.InventoryID, updatedUnit)
	if err != nil {
		return nil, err
	}

	sellerCredit, err := r.sellerRepository.GetSellerCredit(orderDetails.SellerID)
	if err != nil {
		return nil, err
	}

	err = r.sellerRepository.UpdateSellerCredit(orderDetails.SellerID, sellerCredit-orderDetails.Price)
	if err != nil {
		return nil, err
	}

	return orderDetails, nil
}

// ------------------------------------------Seller Control Orders------------------------------------\\

func (r *orderUseCase) GetSellerOrders(sellerID string, remainingQuery string) (*[]responsemodel.OrderDetails, error) {
	userOrders, err := r.repo.GetSellerOrders(sellerID, remainingQuery)
	if err != nil {
		return nil, err
	}
	return userOrders, nil
}

func (r *orderUseCase) ConfirmDeliverd(sellerID string, orderItemID string) (*responsemodel.OrderDetails, error) {

	err := r.repo.UpdateDeliveryTime(sellerID, orderItemID)
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	fmt.Println("hiii")

	orderDetails, err := r.repo.UpdateOrderDelivered(sellerID, orderItemID)
	if err != nil {
		return nil, err
	}

	err = r.repo.UpdateOrderPaymetSuccess(sellerID, orderItemID)
	if err != nil {
		return nil, err
	}

	sellerCredit, err := r.sellerRepository.GetSellerCredit(sellerID)
	if err != nil {
		return nil, err
	}

	err = r.sellerRepository.UpdateSellerCredit(sellerID, sellerCredit+orderDetails.Price)
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

	// err = r.repo.UpdateDeliveryTime(sellerID, orderID)
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
