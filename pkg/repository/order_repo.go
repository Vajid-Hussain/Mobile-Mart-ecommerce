package repository

import (
	"errors"
	"fmt"
	"time"

	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
	resCustomError "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel/custom_error"
	interfaces "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/repository/interface"
	"gorm.io/gorm"
)

type orderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) interfaces.IOrderRepository {
	return &orderRepository{DB: db}
}

func (d *orderRepository) CreateOrder(order *requestmodel.Order) (*responsemodel.Order, error) {

	var orderSucess = &responsemodel.Order{}

	query := "INSERT INTO orders (user_id, address_id, payment_method, order_id_razopay, coupon_code) VALUES(?, ?, ?, ?, ?) RETURNING*"
	result := d.DB.Raw(query, order.UserID, order.Address, order.Payment, order.OrderIDRazopay, order.Coupon).Scan(&orderSucess)
	if result.Error != nil {
		return nil, errors.New("face some issue while creating order")
	}
	if result.RowsAffected == 0 {

		return nil, resCustomError.ErrNoRowAffected
	}
	return orderSucess, nil
}

func (d *orderRepository) AddProdutToOrderProductTable(order *requestmodel.Order, orderDetails *responsemodel.Order) (*responsemodel.Order, error) {

	var orderProduct responsemodel.OrderProducts
	today := time.Now().Format("2006-01-02 15:04:05")

	for _, data := range order.Cart {
		query := "INSERT INTO order_products (order_id, inventory_id, seller_id, quantity, order_date, order_status,payment_status, price, discount,payable_amount) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?) RETURNING*"
		d.DB.Raw(query, orderDetails.ID, data.InventoryID, data.SellerID, data.Quantity, today, order.OrderStatus, order.PaymentStatus, data.Price, data.Discount, data.FinalPrice).Scan(&orderProduct)
		orderDetails.Orders = append(orderDetails.Orders, orderProduct)
	}
	return orderDetails, nil
}

func (d *orderRepository) GetAddressExist(userID, addressesID string) error {
	query := "SELECT * FROM addresses WHERE userid= ? AND id= ?"
	result := d.DB.Exec(query, userID, addressesID)
	if result.Error != nil {
		return errors.New("face some issue while chcking address is exist of user")
	}
	if result.RowsAffected == 0 {
		return errors.New("user does not have specified address")
	}
	return nil
}

func (d *orderRepository) GetOrderShowcase(userID string) (*[]responsemodel.OrderShowcase, error) {
	fmt.Println("***", userID)
	var OrderShowcase []responsemodel.OrderShowcase
	query := "SELECT * FROM orders INNER JOIN order_products ON orders.id=order_products.order_id INNER JOIN inventories ON inventories.id=order_products.inventory_id WHERE orders.user_id=? ORDER BY order_products.item_id DESC "
	result := d.DB.Raw(query, userID).Scan(&OrderShowcase)
	if result.Error != nil {
		return nil, errors.New("face some issue while order showcase")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrNoRowAffected
	}
	return &OrderShowcase, nil
}

func (d *orderRepository) GetSingleOrder(orderID string, userID string) (*responsemodel.SingleOrder, error) {

	var OrderShowcase *responsemodel.SingleOrder
	query := "SELECT * FROM orders INNER JOIN order_products ON orders.id=order_products.order_id INNER JOIN inventories ON inventories.id=order_products.inventory_id INNER JOIN addresses ON addresses.id= orders.address_id WHERE order_products.item_id=? AND orders.user_id=? "
	result := d.DB.Raw(query, orderID, userID).Scan(&OrderShowcase)
	if result.Error != nil {
		return nil, errors.New("face some issue while get single order")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrNoRowAffected
	}
	return OrderShowcase, nil
}

func (d *orderRepository) GetInventoryUnits(inventoryID string) (*uint, error) {

	var units uint
	query := "SELECT units FROM inventories WHERE id=?"
	result := d.DB.Raw(query, inventoryID).Scan(&units)
	if result.Error != nil {
		return nil, errors.New("face some issue while get inventory units")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrNoRowAffected
	}
	return &units, nil
}

func (d *orderRepository) GetPaymentType(orderItemID string) (string, error) {
	var paymentType string
	fmt.Println("@@", orderItemID)
	query := "SELECT payment_method FROM orders INNER JOIN order_products ON orders.id=order_products.order_id WHERE order_products.item_id=?"
	result := d.DB.Raw(query, orderItemID).Scan(&paymentType)
	if result.Error != nil {
		return "", errors.New("face some issue while get payment type of order")
	}
	if result.RowsAffected == 0 {
		return "", resCustomError.ErrNoRowAffected
	}
	return paymentType, nil

}

func (d *orderRepository) UpdateInventoryUnits(inventoryID string, units uint) error {

	query := "UPDATE inventories SET units= ? WHERE id =?"
	result := d.DB.Exec(query, units, inventoryID)
	if result.Error != nil {
		return errors.New("face some issue while updating inventory unit")
	}
	if result.RowsAffected == 0 {

		return resCustomError.ErrNoRowAffected
	}
	return nil
}

func (d *orderRepository) GetOrderPrice(orderID string) (uint, error) {
	fmt.Println("&&77", orderID)
	var price uint
	query := "SELECT price FROM orders WHERE id =?"
	result := d.DB.Raw(query, orderID).Scan(&price)
	if result.Error != nil {
		return 0, errors.New("face some issue while get credit from seller table")
	}
	if result.RowsAffected == 0 {
		return 0, resCustomError.ErrNoRowAffected
	}
	return price, nil
}

func (d *orderRepository) UpdateUserOrderCancel(orderItemID string, userID string) (*responsemodel.OrderDetails, error) {

	var cancelOrder responsemodel.OrderDetails
	today := time.Now().Format("2006-01-02 15:04:05")

	query := "UPDATE order_products SET order_status= 'cancelled', payment_status= 'refunded', end_date=? FROM orders WHERE orders.id=order_products.order_id AND item_id=? AND user_id= ? AND order_status='processing' RETURNING*"
	result := d.DB.Raw(query, today, orderItemID, userID).Scan(&cancelOrder)
	if result.Error != nil {
		return nil, errors.New("face some issue while order is cancel")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrProductOrderCompleted
	}
	return &cancelOrder, nil
}

func (d *orderRepository) UpdateUserOrderReturn(orderItemID string, userID string) (*responsemodel.OrderDetails, error) {

	var returnOrder responsemodel.OrderDetails
	today := time.Now().Format("2006-01-02 15:04:05")

	query := "UPDATE order_products SET order_status= 'return', payment_status= 'refunded', end_date=? FROM orders WHERE orders.id=order_products.order_id AND item_id=? AND user_id= ? AND order_status='delivered' RETURNING*"
	result := d.DB.Raw(query, today, orderItemID, userID).Scan(&returnOrder)
	if result.Error != nil {
		return nil, errors.New("face some issue while order is return")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("no deliverd order exist for the given order item id of the user")
	}
	return &returnOrder, nil
}

func (d *orderRepository) UpdateDeliveryTimeByUser(userID string, orderItemID string) error {

	delivaryTime := time.Now().Format("2006-01-02 15:04:05")

	query := "UPDATE orders SET delivery_date= ? WHERE user_id= ? AND id = ?"
	result := d.DB.Exec(query, delivaryTime, userID, orderItemID)
	if result.Error != nil {
		return errors.New("face some issue while updating delivary time")
	}
	if result.RowsAffected == 0 {
		return resCustomError.ErrNoRowAffected
	}
	return nil
}

func (d *orderRepository) GetOrderExistOfUser(orderItemID, userID string) error {

	query := "SELECT * FROM orders INNER JOIN order_products ON orders.id=order_products.order_id WHERE item_id= $1 AND user_id= $2"

	result := d.DB.Exec(query, orderItemID, userID)
	if result.Error != nil {
		return errors.New("encountered an issue while checking if the order exists")
	}
	if result.RowsAffected == 0 {
		return errors.New("no orders were found matching the specified criteria")
	}
	if result.RowsAffected != 0 {
		return nil
	}
	return nil
}

// ------------------------------------------Seller Control Orders------------------------------------\\

func (d *orderRepository) GetSellerOrders(sellerID string, remainingQuery string) (*[]responsemodel.OrderDetails, error) {

	var orderList *[]responsemodel.OrderDetails
	query := "SELECT * FROM orders INNER JOIN order_products ON orders.id= order_products.order_id WHERE seller_id=? AND order_status" + remainingQuery
	result := d.DB.Raw(query, sellerID).Scan(&orderList)
	if result.Error != nil {
		return nil, errors.New("face some issue while get user orders")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrNoRowAffected
	}
	return orderList, nil
}

func (d *orderRepository) UpdateDeliveryTime(sellerID string, orderItemID string) error {

	delivaryTime := time.Now().Format("2006-01-02 15:04:05")

	query := "UPDATE order_products SET end_date= ? FROM orders WHERE orders.id= order_products.order_id AND seller_id= ? AND order_products.item_id= ? AND order_status='processing'"
	result := d.DB.Exec(query, delivaryTime, sellerID, orderItemID)
	if result.Error != nil {
		return errors.New("face some issue while updating delivary time")
	}
	if result.RowsAffected == 0 {
		return resCustomError.ErrNoRowAffected
	}
	return nil
}

func (d *orderRepository) UpdateOrderDelivered(sellerID string, orderItemID string) (*responsemodel.OrderDetails, error) {
	var deliveryDetails responsemodel.OrderDetails
	query := "UPDATE order_products SET order_status='delivered' FROM orders WHERE orders.id= order_products.order_id AND seller_id= ? AND order_products.item_id= ? RETURNING*"
	result := d.DB.Raw(query, sellerID, orderItemID).Scan(&deliveryDetails)
	if result.Error != nil {
		return nil, errors.New("face some issue while update order delivered")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrNoRowAffected
	}
	return &deliveryDetails, nil
}

func (d *orderRepository) UpdateOrderPaymetSuccess(sellerID string, orderItemID string) error {

	query := "UPDATE order_products SET payment_status = 'success' FROM orders WHERE orders.id= order_products.order_id AND seller_id= ? AND order_products.item_id= ?"
	result := d.DB.Exec(query, sellerID, orderItemID)
	if result.Error != nil {
		return errors.New("face some issue while update payment status success")
	}
	if result.RowsAffected == 0 {
		return resCustomError.ErrNoRowAffected
	}
	return nil
}

func (d *orderRepository) UpdateOrderCancel(orderID string, sellerID string) (*responsemodel.OrderDetails, error) {

	var cancelOrder responsemodel.OrderDetails
	query := "UPDATE orders SET order_status= 'cancel', payment_status='cancel' WHERE id=? AND seller_id= ? AND order_status='processing' RETURNING*"
	result := d.DB.Raw(query, orderID, sellerID).Scan(&cancelOrder)
	if result.Error != nil {
		return nil, errors.New("face some issue while order is cancel")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrProductOrderCompleted
	}
	return &cancelOrder, nil
}

func (d *orderRepository) GetOrderExistOfSeller(orderID, sellerID string) error {

	query := "SELECT * FROM orders WHERE id= $1 AND seller_id=$2"

	result := d.DB.Exec(query, orderID, sellerID)
	if result.Error != nil {
		return errors.New("encountered an issue while checking if the order exists")
	}
	if result.RowsAffected == 0 {
		return errors.New("no orders were found matching the specified criteria")
	}
	if result.RowsAffected != 0 {
		return nil
	}
	return nil
}

// ------------------------------------------Sales Report------------------------------------\\

func (d *orderRepository) GetSalesReportByYear(sellerID string, balanceQuery string) (*responsemodel.SalesReport, error) {
	var report responsemodel.SalesReport
	query := "SELECT COUNT(*) AS Orders, SUM(quantity) AS Quantity, SUM(price) AS Price FROM order_products WHERE seller_id= ? AND order_status='delivered' AND" + balanceQuery
	result := d.DB.Raw(query, sellerID).Scan(&report)
	if result.Error != nil {
		return nil, errors.New("face some issue while get report")
	}
	return &report, nil
}

func (d *orderRepository) GetSalesReportByDays(sellerID string, days string) (*responsemodel.SalesReport, error) {
	var report responsemodel.SalesReport
	remainingQuery := "(now() - interval '" + days + " day')"
	query := "SELECT COUNT(*) AS Orders, SUM(quantity) AS Quantity, SUM(price) AS Price FROM order_products WHERE seller_id = ? AND order_status='delivered' AND order_date >= " + remainingQuery
	result := d.DB.Raw(query, sellerID).Scan(&report)

	if result.Error != nil {
		return nil, errors.New("face some issue while get report by days")
	}
	return &report, nil
}

func (d *orderRepository) GetOrderXlSalesReport(sellerID string) (*[]responsemodel.XlSalesReport, error) {
	var order []responsemodel.XlSalesReport

	query := "SELECT * FROM orders INNER JOIN order_products ON order_products.order_id=orders.id INNER JOIN inventories ON inventories.id=order_products.inventory_id WHERE order_products.order_status='delivered' AND order_products.seller_id=? "
	result := d.DB.Raw(query, sellerID).Scan(&order)
	if result.Error != nil {
		return nil, errors.New("face some issue while order is cancel")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrProductOrderCompleted
	}
	return &order, nil
}

// ------------------------------------------category_offers------------------------------------\\

func (d *orderRepository) GetCategoryOffers(productID string) uint {
	var categoryDiscount uint
	query := "SELECT category_discount FROM category_offers RIGHT JOIN inventories ON inventories.seller_id=category_offers.seller_id AND category_offers.category_id=inventories.category_id AND category_offers.status='active' AND category_offers.end_date>now() WHERE inventories.status='active'  AND inventories.id=?"
	d.DB.Raw(query, productID).Scan(&categoryDiscount)
	return categoryDiscount
}

func (d *orderRepository) CheckCouponAppliedOrNot(userID, couponID string) uint {
	var exist uint
	query := "SELECT COUNT(*) FROM orders WHERE user_id=? AND coupon_code= ?"
	d.DB.Raw(query, userID, couponID).Scan(&exist)
	return exist
}

func (d *orderRepository) GetOrderFullDetails(orderItemID string) (*responsemodel.Invoice, error) {
	var orderDetails responsemodel.Invoice
	query := "SELECT * FROM orders INNER JOIN order_products ON orders.id=order_products.order_id WHERE order_products.item_id= ?;	"
	result := d.DB.Raw(query, orderItemID).Scan(&orderDetails)
	if result.Error != nil {
		return nil, errors.New("face some issue while get order details")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrNoRowAffected
	}
	return &orderDetails, nil
}

func (d *orderRepository) GetAddressForInvoice(addressID string) (*requestmodel.Address, error) {

	var address *requestmodel.Address
	query := "SELECT * FROM addresses WHERE id= ?"
	result := d.DB.Raw(query, addressID).Scan(&address)
	if result.Error != nil {
		return nil, errors.New("face some issue while address fetch")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrNoRowAffected
	}
	return address, nil
}

func (d *orderRepository) GetAInventoryForInvoice(id string) (*responsemodel.InventoryRes, error) {
	var inventory responsemodel.InventoryRes

	query := "SELECT * FROM category_offers RIGHT JOIN inventories ON category_offers.category_id= inventories.category_id AND inventories.seller_id=category_offers.seller_id AND category_offers.status='active' AND category_offers.end_date>=now() WHERE inventories.id=? AND inventories.status='active'"
	result := d.DB.Raw(query, id).Scan(&inventory)
	if result.Error != nil {
		return nil, errors.New("can't get inventory data from db or inventory is not active state")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrNoRowAffected
	}
	return &inventory, nil
}
