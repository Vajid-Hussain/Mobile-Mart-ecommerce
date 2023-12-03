package repository

import (
	"errors"
	"fmt"

	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
	resCustomError "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel/custom_error"
	interfaces "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/repository/interface"
	"gorm.io/gorm"
)

type paymentRepo struct {
	DB *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) interfaces.IPaymentRepository {
	return &paymentRepo{DB: db}
}

func (d *paymentRepo) CreateOrUpdateWallet(userID string, creditAmount uint) (*uint, error) {

	var currentBalance uint
	query := "INSERT INTO wallets (user_id, balance) VALUES ($1, $2) ON CONFLICT(user_id) DO UPDATE SET balance=wallets.balance + $2"
	result := d.DB.Exec(query, userID, creditAmount)
	if result.Error != nil {
		return nil, errors.New("face some issue while intract with wallet for made change in balance")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrNoRowAffected
	}
	return &currentBalance, nil
}

func (d *paymentRepo) GetWalletbalance(userID string) (*uint, error) {

	var currentBalance uint
	query := "SELECT balance FROM wallet WHERE user_id= ?"
	result := d.DB.Raw(query, userID).Scan(&currentBalance)
	if result.Error != nil {
		return nil, errors.New("face some issue while fetch user wallet balance")
	}
	if result.RowsAffected == 0 {

		return nil, resCustomError.ErrNoRowAffected
	}
	return &currentBalance, nil
}

// ------------------------------------------Online Payment------------------------------------\\

func (d *paymentRepo) OnlinePayment(userID, orderID string) (*responsemodel.OnlinePayment, error) {

	var orderDetails responsemodel.OnlinePayment
	query := "SELECT * FROM users INNER JOIN orders ON orders.user_id = users.id INNER JOIN addresses ON addresses.id = address_id INNER JOIN order_products ON order_products.order_id=orders.id WHERE orders.user_id=? AND payment_status = 'pending' AND payment_method= 'ONLINE'"
	result := d.DB.Raw(query, userID).Scan(&orderDetails)
	if result.Error != nil {
		return nil, errors.New("face some issue while processing online payment")
	}
	if result.RowsAffected == 0 {
		fmt.Println("no order matched")
		return nil, resCustomError.ErrNoRowAffected
	}
	return &orderDetails, nil
}

func (d *paymentRepo) GetFinalPriceByorderID(orderID string) (uint, error) {
	var finalPrice uint
	query := "SELECT sum(price) FROM orders INNER JOIN order_products ON order_products.order_id= orders.id WHERE orders.id= ?"
	result := d.DB.Raw(query, orderID).Scan(&finalPrice)
	if result.Error != nil {
		return 0, errors.New("face some issue while getting tatal amount of order by using order id")
	}
	if result.RowsAffected == 0 {
		return 0, resCustomError.ErrNoRowAffected
	}
	return finalPrice, nil
}

func (d *paymentRepo) UpdateOnlinePaymentSucess(orderID string) (*[]responsemodel.OrderDetails, error) {

	var orders []responsemodel.OrderDetails
	query := "UPDATE order_products SET payment_status='success', order_status='processing' FROM orders WHERE order_products.order_id=orders.id AND orders.order_id_razopay = ? RETURNING*"
	result := d.DB.Raw(query, orderID).Scan(&orders)
	if result.Error != nil {
		return nil, errors.New("face some issue while update order status and payment status on verify online payment success")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrProductOrderCompleted
	}
	return &orders, nil
}
