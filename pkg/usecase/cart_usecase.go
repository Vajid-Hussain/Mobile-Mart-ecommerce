package usecase

import (
	"errors"
	"fmt"

	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
	interfaces "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/repository/interface"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
)

type cartUseCase struct {
	repo interfaces.ICartRepository
}

func NewCartUseCase(repository interfaces.ICartRepository) interfaceUseCase.ICartUseCase {
	return &cartUseCase{repo: repository}
}

func (r *cartUseCase) CreateCart(cart *requestmodel.Cart) (*requestmodel.Cart, error) {

	count, err := r.repo.IsInventoryExistInCart(cart.InventoryID, cart.UserID)
	if err != nil {
		return nil, err
	}

	if count >= 1 {
		return nil, errors.New("inverntory alrady exist in cart now you can purchase")
	}

	productPrice, err := r.repo.GetInventoryPrice(cart.InventoryID)
	if err != nil {
		return nil, err
	}

	cart.Price = productPrice
	cart.Quantity = 1

	inserCart, err := r.repo.InsertToCart(cart)
	if err != nil {
		return nil, err
	}
	return inserCart, nil
}

func (r *cartUseCase) DeleteInventoryFromCart(inventoryID string, userID string) error {
	err := r.repo.DeleteInventoryFromCart(inventoryID, userID)
	if err != nil {
		return err
	}
	return nil
}

func (r *cartUseCase) QuantityIncriment(inventoryID string, userID string) (*requestmodel.Cart, error) {

	singleInventory, err := r.repo.GetSingleInverntory(inventoryID, userID)
	if err != nil {
		return nil, err
	}

	price := singleInventory.Price / singleInventory.Quantity
	currentQuantity := singleInventory.Quantity

	singleInventory.Quantity = currentQuantity + 1
	singleInventory.Price = singleInventory.Quantity * price

	singleInventory, err = r.repo.UpdateQuantityAndPrice(singleInventory)
	if err != nil {
		return nil, err
	}
	return singleInventory, nil
}

func (r *cartUseCase) QuantityDecrease(inventoryID string, userID string) (*requestmodel.Cart, error) {

	singleInventory, err := r.repo.GetSingleInverntory(inventoryID, userID)
	if err != nil {
		return nil, err
	}

	if singleInventory.Quantity == 1 {
		return singleInventory, errors.New("reach the maximum limit")
	}

	price := singleInventory.Price / singleInventory.Quantity
	currentQuantity := singleInventory.Quantity

	singleInventory.Quantity = currentQuantity - 1
	singleInventory.Price = singleInventory.Quantity * price

	singleInventory, err = r.repo.UpdateQuantityAndPrice(singleInventory)
	if err != nil {
		return nil, err
	}
	return singleInventory, nil
}

func (r *cartUseCase) ShowCart(userID string) (*responsemodel.UserCart, error) {

	cart, err := r.repo.GetCartCriteria(userID)
	if err != nil {
		return cart, err
	}
	fmt.Println("&&%%%%%%%%&", cart)

	// cartInventories, err := r.repo.GetCart(userID)
	// if err != nil {
	// 	return cart, err
	// }
	// fmt.Println("&$$$$$$&", cart)

	// cart.Cart = *cartInventories
	return cart, nil
}
