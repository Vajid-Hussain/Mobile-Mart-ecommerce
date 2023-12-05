package repository

import (
	"errors"

	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
	resCustomError "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel/custom_error"
	interfaces "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/repository/interface"
	"gorm.io/gorm"
)

type cartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) interfaces.ICartRepository {
	return &cartRepository{DB: db}
}

func (d *cartRepository) IsInventoryExistInCart(inventoryID string, userID string) (int, error) {
	var inventoryCount int

	query := "SELECT count(*) FROM carts WHERE inventory_id=? AND user_id=? AND status='active' "
	result := d.DB.Raw(query, inventoryID, userID).Scan(&inventoryCount)
	if result.Error != nil {
		return 0, errors.New("face some issue while finding inventory is exist in cart")
	}
	return inventoryCount, nil
}

func (d *cartRepository) InsertToCart(cart *requestmodel.Cart) (*requestmodel.Cart, error) {

	query := "INSERT INTO carts (user_id, inventory_id, quantity) VALUES (?, ?,  ?)   RETURNING *;"
	result := d.DB.Raw(query, cart.UserID, cart.InventoryID, cart.Quantity).Scan(&cart)

	if result.Error != nil {
		return nil, errors.New("face some issue while inventory insert to cart ")
	}
	if result.RowsAffected == 0 {

		return nil, resCustomError.ErrNoRowAffected
	}
	return cart, nil
}

func (d *cartRepository) GetInventoryPrice(inventoryID string) (uint, error) {
	var price uint
	query := "SELECT saleprice FROM inventories WHERE id= ? AND status = 'active'"
	result := d.DB.Raw(query, inventoryID).Scan(&price)
	if result.Error != nil {
		return 0, errors.New("face some issue while get user profile ")
	}
	if result.RowsAffected == 0 {
		return 0, resCustomError.ErrNoRowAffected
	}
	return price, nil
}

func (d *cartRepository) DeleteInventoryFromCart(inventoryID string, userID string) error {

	query := "UPDATE carts SET status='delete' WHERE inventory_id = ? AND user_id= ? AND status= 'active'"
	result := d.DB.Exec(query, inventoryID, userID)
	if result.Error != nil {
		return errors.New("face some issue while delete inventory in cart")
	}
	if result.RowsAffected == 0 {
		return resCustomError.ErrNoRowAffected
	}
	return nil
}

func (d *cartRepository) GetSingleInverntory(inventoryID string, userID string) (*requestmodel.Cart, error) {

	var singleInventory *requestmodel.Cart
	query := "SELECT * FROM carts WHERE user_id=? AND inventory_id=? AND status='active'"
	result := d.DB.Raw(query, userID, inventoryID).Scan(&singleInventory)
	if result.Error != nil {
		return nil, errors.New("face some issue while fetching inventory in cart")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrNoRowAffected
	}
	return singleInventory, nil
}

func (d *cartRepository) UpdateQuantity(cart *requestmodel.Cart) (*requestmodel.Cart, error) {

	var singleInventory *requestmodel.Cart
	query := "UPDATE carts SET quantity= ? WHERE user_id=? AND inventory_id= ? AND status='active' RETURNING* ;"
	result := d.DB.Raw(query, cart.Quantity, cart.UserID, cart.InventoryID).Scan(&singleInventory)
	if result.Error != nil {
		return nil, errors.New("face some issue while quantity updating cart")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrNoRowAffected
	}
	return singleInventory, nil
}

func (d *cartRepository) GetCart(userID string) (*[]responsemodel.CartInventory, error) {
	var cartView *[]responsemodel.CartInventory
	query := "SELECT * FROM carts INNER JOIN inventories ON id=inventory_id INNER JOIN categories ON categories.id=inventories.category_id WHERE carts.user_id=? AND carts.status='active'"
	result := d.DB.Raw(query, userID).Scan(&cartView)
	if result.Error != nil {
		return nil, errors.New("face some issue while  get cart")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("user have no cart")
	}
	return cartView, nil
}

func (d *cartRepository) GetNetAmoutOfCart(userID string, inventoryID string) (uint, error) {

	var NetCart uint
	query := "SELECT saleprice FROM carts INNER JOIN inventories ON id=inventory_id WHERE carts.user_id=? AND id= ? AND carts.status='active'"
	result := d.DB.Raw(query, userID, inventoryID).Scan(&NetCart)
	if result.Error != nil {
		return 0, errors.New("user have no Cart")
	}
	if result.RowsAffected == 0 {
		return 0, errors.New("user have no cart")
	}
	return NetCart, nil
}

func (d *cartRepository) GetCartCriteria(userID string) (uint, error) {

	var count uint
	query := "SELECT SUM(quantity) FROM carts WHERE user_id=? AND status='active'"
	result := d.DB.Raw(query, userID)
	result.Row().Scan(&count)
	if result.Error != nil {
		return 0, errors.New("face some issue while  get cart")
	}
	if result.RowsAffected == 0 {
		return 0, resCustomError.ErrNoRowAffected
	}
	return count, nil
}
