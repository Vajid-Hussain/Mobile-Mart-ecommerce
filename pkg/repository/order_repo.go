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

func (d *orderRepository) CreateOrder(order *requestmodel.Order) (*[]responsemodel.OrderSuccess, error) {

	today := time.Now().Format("2006-01-02 15:04:05")
	var orderSucess *[]responsemodel.OrderSuccess
	var result *gorm.DB

	// var paymentStatus string
	// if order.Payment == "COD" {
	// 	paymentStatus = "active"
	// } else {
	// 	paymentStatus = "pending"
	// }

	for _, data := range order.Cart {
		fmt.Println("$$$$$", data, data.Saleprice, data.Price)
		query := `INSERT INTO orders (user_id, address_id, payment_method, inventory_id, seller_id, price, quantity,  order_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?) RETURNING*`
		result = d.DB.Raw(query, order.UserID, order.Address, order.Payment, data.InventoryID, data.SellerID, data.Price, data.Quantity, today).Scan(&orderSucess)
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
	query := "SELECT * FROM inventories INNER JOIN  orders ON orders.inventory_id= inventories.id WHERE orders.user_id= ?"
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
