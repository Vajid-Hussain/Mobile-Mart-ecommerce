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

func (d *orderRepository) CreateOrder(order *requestmodel.Order) (*responsemodel.OrderSuccess, error) {

	today := time.Now().Format("2006-01-02 15:04:05")
	var orderSucess = &responsemodel.OrderSuccess{}
	var orderData responsemodel.OrderDetails
	var result *gorm.DB

	for _, data := range order.Cart {
		query := `INSERT INTO orders (user_id, address_id, payment_method, inventory_id, seller_id, price, quantity,  order_date, order_status,  order_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?) RETURNING*`
		result = d.DB.Raw(query, order.UserID, order.Address, order.Payment, data.InventoryID, data.SellerID, data.Price, data.Quantity, today, order.OrderStatus, order.OrderID).Scan(&orderData)
		orderSucess.TotalWorth += orderData.Price
		orderSucess.Orders = append(orderSucess.Orders, orderData)
	}
	if result.Error != nil {
		return nil, errors.New("face some issue while creating order")
	}
	if result.RowsAffected == 0 {

		return nil, resCustomError.ErrNoRowAffected
	}
	return orderSucess, nil
}

func (d *orderRepository) GetOrderShowcase(userID string) (*[]responsemodel.OrderShowcase, error) {

	var OrderShowcase []responsemodel.OrderShowcase
	query := "SELECT * FROM inventories INNER JOIN  orders ON orders.inventory_id= inventories.id WHERE orders.user_id= ? ORDER BY orders.id DESC"
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
	query := "SELECT * FROM inventories INNER JOIN  orders ON orders.inventory_id= inventories.id INNER JOIN addresses ON addresses.id=orders.address_id WHERE orders.id= ? AND user_id= ?"
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

func (d *orderRepository) UpdateUserOrderCancel(orderID string, userID string) (*responsemodel.OrderDetails, error) {

	var cancelOrder responsemodel.OrderDetails
	query := "UPDATE orders SET order_status= 'cancel' WHERE id=? AND user_id= ? AND order_status='processing' RETURNING*"
	result := d.DB.Raw(query, orderID, userID).Scan(&cancelOrder)
	if result.Error != nil {
		return nil, errors.New("face some issue while order is cancel")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrProductOrderCompleted
	}
	return &cancelOrder, nil
}

func (d *orderRepository) UpdateDeliveryTimeByUser(userID string, orderID string) error {

	delivaryTime := time.Now().Format("2006-01-02 15:04:05")

	query := "UPDATE orders SET delivery_date= ? WHERE user_id= ? AND id = ?"
	result := d.DB.Exec(query, delivaryTime, userID, orderID)
	if result.Error != nil {
		return errors.New("face some issue while updating delivary time")
	}
	if result.RowsAffected == 0 {
		return resCustomError.ErrNoRowAffected
	}
	return nil
}

// ------------------------------------------Seller Control Orders------------------------------------\\

func (d *orderRepository) GetSellerOrders(sellerID string, remainingQuery string) (*[]responsemodel.OrderDetails, error) {

	var orderList *[]responsemodel.OrderDetails
	query := "SELECT * FROM orders WHERE seller_id=? AND order_status" + remainingQuery
	result := d.DB.Raw(query, sellerID).Scan(&orderList)
	if result.Error != nil {
		return nil, errors.New("face some issue while get user orders")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrNoRowAffected
	}
	return orderList, nil
}

func (d *orderRepository) UpdateDeliveryTime(sellerID string, orderID string) error {

	delivaryTime := time.Now().Format("2006-01-02 15:04:05")

	query := "UPDATE orders SET delivery_date= ? WHERE seller_id= ? AND id = ?"
	result := d.DB.Exec(query, delivaryTime, sellerID, orderID)
	if result.Error != nil {
		return errors.New("face some issue while updating delivary time")
	}
	if result.RowsAffected == 0 {
		return resCustomError.ErrNoRowAffected
	}
	return nil
}

func (d *orderRepository) UpdateOrderDelivered(sellerID string, orderID string) (*responsemodel.OrderDetails, error) {
	var deliveryDetails responsemodel.OrderDetails
	query := "UPDATE orders SET order_status = 'delivered' WHERE seller_id= ? AND id= ? AND order_status='processing' RETURNING*"
	result := d.DB.Raw(query, sellerID, orderID).Scan(&deliveryDetails)
	if result.Error != nil {
		return nil, errors.New("face some issue while order is delevered")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrProductOrderCompleted
	}
	return &deliveryDetails, nil
}

func (d *orderRepository) UpdateOrderPaymetSuccess(sellerID string, orderID string) error {

	query := "UPDATE orders SET payment_status='success' WHERE seller_id= ? AND id= ? AND order_status='delivered'"
	result := d.DB.Exec(query, sellerID, orderID)
	if result.Error != nil {
		return errors.New("face some issue while update payment status success")
	}
	if result.RowsAffected == 0 {
		return resCustomError.ErrProductOrderCompleted
	}
	return nil
}

func (d *orderRepository) UpdateOrderCancel(orderID string, sellerID string) (*responsemodel.OrderDetails, error) {

	var cancelOrder responsemodel.OrderDetails
	query := "UPDATE orders SET order_status= 'cancel' WHERE id=? AND seller_id= ? AND order_status='processing' RETURNING*"
	result := d.DB.Raw(query, orderID, sellerID).Scan(&cancelOrder)
	if result.Error != nil {
		return nil, errors.New("face some issue while order is cancel")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrProductOrderCompleted
	}
	return &cancelOrder, nil
}

// ------------------------------------------Sales Report------------------------------------\\

func (d *orderRepository) GetSalesReportByYear(sellerID string, balanceQuery string) (*responsemodel.SalesReport, error) {
	var report responsemodel.SalesReport
	query := "SELECT COUNT(*) AS Orders, SUM(quantity) AS Quantity, SUM(price) AS Price FROM orders WHERE seller_id= ? AND order_status='delivered' AND" + balanceQuery
	result := d.DB.Raw(query, sellerID).Scan(&report)
	if result.Error != nil {
		return nil, errors.New("face some issue while get report")
	}
	return &report, nil
}

func (d *orderRepository) GetSalesReportByDays(sellerID string, days string) (*responsemodel.SalesReport, error) {
	var report responsemodel.SalesReport
	remainingQuery := "(now() - interval '" + days + " day')"
	query := "SELECT COUNT(*) AS Orders, SUM(quantity) AS Quantity, SUM(price) AS Price FROM orders WHERE seller_id = ? AND order_status='delivered' AND order_date >= " + remainingQuery
	result := d.DB.Raw(query, sellerID).Scan(&report)

	if result.Error != nil {
		return nil, errors.New("face some issue while get report by days")
	}
	return &report, nil
}

// ------------------------------------------Payment------------------------------------\\

func (d *orderRepository) OnlinePayment(userID string) (*responsemodel.OnlinePayment, error) {

	var orderDetails responsemodel.OnlinePayment
	query := "SELECT * FROM users INNER JOIN orders ON orders.user_id = users.id INNER JOIN addresses ON addresses.id = address_id WHERE orders.user_id=? AND payment_status = 'pending' AND payment_method= 'ONLINE'"
	result := d.DB.Raw(query, userID).Scan(&orderDetails)
	if result.Error != nil {
		return nil, errors.New("face some issue while processing online payment")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrNoRowAffected
	}
	return &orderDetails, nil
}

func (d *orderRepository) GetFinalPriceByorderID(orderID string) (uint, error) {
	var finalPrice uint
	query := "SELECT SUM(price) FROM orders WHERE order_id= ?"
	result := d.DB.Raw(query, orderID).Scan(&finalPrice)
	if result.Error != nil {
		return 0, errors.New("face some issue while getting tatal amount of order by using order id")
	}
	if result.RowsAffected == 0 {
		return 0, resCustomError.ErrNoRowAffected
	}
	return finalPrice, nil
}