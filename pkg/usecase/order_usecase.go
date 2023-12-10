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
	"github.com/jung-kurt/gofpdf"
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

		if order.FinalPrice < couponData.MinimumRequired || order.FinalPrice >= couponData.MaximumAllowed {
			return nil, fmt.Errorf("total price of order is %d not satisfying, for apply this coupon code %s of maximum allowed %d", order.FinalPrice, order.Coupon, couponData.MaximumAllowed)
		}

		rightNow := time.Now()
		if couponData.EndDate.Before(rightNow) {
			return nil, errors.New("coupon exeed the expiredata, better luck next times")
		}

		exist := r.repo.CheckCouponAppliedOrNot(order.UserID, order.Coupon)
		if exist > 0 {
			return nil, fmt.Errorf("you are alredy apply %s coupon for %d time", order.Coupon, exist)
		}

		order.CouponDiscount = couponData.Discount
	}

	// find total amount
	order.FinalPrice = 0
	for i, product := range order.Cart {
		order.Cart[i].Price = helper.FindDiscount(float64(product.Price), float64(product.CategoryDiscount+product.Discount))
		order.Cart[i].Discount = product.Discount + product.CategoryDiscount
		order.Cart[i].FinalPrice = helper.FindDiscount(float64(product.Price), float64(product.Discount+product.CategoryDiscount+order.CouponDiscount))
		order.FinalPrice += order.Cart[i].FinalPrice
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

// ------------------------------------------Order Invoice------------------------------------\\

func (r *orderUseCase) OrderInvoiceCreation(orderItemID string) (*gofpdf.Fpdf, error) {
	orderDetails, err := r.repo.GetOrderFullDetails(orderItemID)
	if err != nil {
		return nil, err
	}

	sellerDetails, err := r.sellerRepository.GetSingleSeller(orderDetails.SellerID)
	if err != nil {
		return nil, err
	}

	userAddresses, err := r.repo.GetAddressForInvoice(orderDetails.AddressID)
	if err != nil {
		return nil, err
	}

	product, err := r.repo.GetAInventoryForInvoice(orderDetails.InventoryID)
	if err != nil {
		return nil, err
	}

	//create pdf
	marginX := 10.0
	marginY := 20.0
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(marginX, marginY, marginX)
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 16)
	_, lineHeight := pdf.GetFontSize()
	currentY := pdf.GetY() + lineHeight
	pdf.SetY(currentY)
	pdf.Cell(40, 10, "Tax Invoice")
	pdf.Cell(40, 10, "|  Order ID: "+orderItemID)
	pdf.Cell(40, 10, "|  Order Date: "+orderDetails.OrderDate.Format("2006-01-02 15:04:05"))
	pdf.Ln(15)

	pdf.Cell(40, 10, "Seller: "+sellerDetails.Name)
	pdf.Ln(10)

	// address
	pdf.Cell(20, 10, "Address")
	pdf.Ln(10)

	pdf.SetFont("Helvetica", "", 12)
	pdf.Cell(40, 10, fmt.Sprintf("Name: %s %s", userAddresses.FirstName, userAddresses.LastName))
	pdf.Ln(10)
	pdf.Cell(0, 10, fmt.Sprintf("Address: %s, %s, %s - %s", userAddresses.Street, userAddresses.City, userAddresses.State, userAddresses.Pincode))
	pdf.Ln(10)
	pdf.Cell(0, 10, fmt.Sprintf("Landmark: %s", userAddresses.LandMark))
	pdf.Ln(10)
	pdf.Cell(0, 10, fmt.Sprintf("Phone Number: %s", userAddresses.PhoneNumber))
	pdf.Ln(20)

	lineHt := 10.0
	const colNumber = 5
	header := [colNumber]string{"No", "Product", "Quantity", "Mrp", "Final-Price"}
	colWidth := [colNumber]float64{10.0, 75.0, 25.0, 40.0, 40.0}

	// Headers
	// pdf.SetFont("Arial", "B", 16)
	pdf.SetFontStyle("B")
	pdf.SetFillColor(200, 200, 200)
	for colJ := 0; colJ < colNumber; colJ++ {
		pdf.CellFormat(colWidth[colJ], lineHt, header[colJ], "1", 0, "CM", true, 0, "")
	}

	pdf.Ln(10)

	// Table data
	pdf.CellFormat(colWidth[0], lineHt, fmt.Sprintf("%d", 1), "1", 0, "CM", false, 0, "")
	pdf.CellFormat(colWidth[1], lineHt, product.Productname, "1", 0, "LM", false, 0, "")
	pdf.CellFormat(colWidth[2], lineHt, fmt.Sprintf("%d", orderDetails.Quantity), "1", 0, "CM", false, 0, "")
	pdf.CellFormat(colWidth[3], lineHt, fmt.Sprintf("%d", product.Mrp), "1", 0, "CM", false, 0, "")
	pdf.CellFormat(colWidth[4], lineHt, fmt.Sprintf("%d", orderDetails.PayableAmount), "1", 0, "CM", false, 0, "")
	pdf.Ln(-1)

	leftIndent := 0.0
	for i := 0; i < 3; i++ {
		leftIndent += colWidth[i]
	}

	pdf.SetX(marginX + leftIndent)
	pdf.CellFormat(colWidth[3], lineHt, "Grand total", "1", 0, "CM", false, 0, "")
	pdf.CellFormat(colWidth[4], lineHt, fmt.Sprintf("%d", orderDetails.PayableAmount), "1", 0, "CM", false, 0, "")
	pdf.Ln(20)

	pdf.Cell(40, 10, "Mobile-mart: Thanks for shopping!")

	err = pdf.OutputFileAndClose("invoice.pdf")
	if err != nil {
		fmt.Println("Error:", err)
	}

	return pdf, nil
}
