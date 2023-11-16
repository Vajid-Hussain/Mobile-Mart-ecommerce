package usecase

import (
	"errors"
	"fmt"
	"strconv"

	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
	resCustomError "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel/custom_error"
	interfaces "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/repository/interface"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
	"github.com/go-playground/validator/v10"
)

type inventoryUseCase struct {
	repo interfaces.IInventoryRepository
}

func NewInventoryUseCase(repository interfaces.IInventoryRepository) interfaceUseCase.IInventoryUseCase {
	return &inventoryUseCase{repo: repository}
}

func (d *inventoryUseCase) AddInventory(inventory *requestmodel.InventoryReq) (*[]responsemodel.Errors, *responsemodel.InventoryRes, error) {
	var afterErrorCorection []responsemodel.Errors
	var result responsemodel.Errors
	validate := validator.New()

	err := validate.Struct(inventory)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Tag() {
				case "required":
					err := fmt.Sprintf("%s is required", e.Field())
					result = responsemodel.Errors{Err: err}
				case "min":
					err := fmt.Sprintf("%s should be at least %s characters", e.Field(), e.Param())
					result = responsemodel.Errors{Err: err}
				case "max":
					err := fmt.Sprintf("%s should be at most %s characters", e.Field(), e.Param())
					result = responsemodel.Errors{Err: err}
				}
				afterErrorCorection = append(afterErrorCorection, result)
			}
		}
		return &afterErrorCorection, nil, errors.New("doesn't fulfill the inventory requirements")
	}

	product, err := d.repo.CreateProduct(inventory)
	if err != nil {
		return nil, nil, err
	}

	return &afterErrorCorection, product, nil
}

func (r *inventoryUseCase) BlockInventory(sellerID string, productID string) error {
	err := r.repo.BlockSingleInventoryBySeller(sellerID, productID)
	if err != nil {
		return err
	}
	return nil
}

func (r *inventoryUseCase) UNBlockInventory(sellerID string, productID string) error {
	err := r.repo.UNBlockSingleInventoryBySeller(sellerID, productID)
	if err != nil {
		return err
	}
	return nil
}

func (r *inventoryUseCase) DeleteInventory(sellerID string, productID string) error {
	err := r.repo.DeleteInventoryBySeller(sellerID, productID)
	if err != nil {
		return err
	}
	return nil
}

func (r *inventoryUseCase) GetAllInventory(page string, limit string) (*[]responsemodel.InventoryShowcase, error) {

	pageNO, err := strconv.Atoi(page)
	if err != nil {
		return nil, resCustomError.ErrConversionOFPage
	}

	limits, err := strconv.Atoi(limit)
	if err != nil {
		return nil, resCustomError.ErrConversionOfLimit
	}

	if pageNO < 1 {
		return nil, resCustomError.ErrPagination
	}

	if limits <= 0 {
		return nil, resCustomError.ErrPageLimit
	}

	offSet := (pageNO * limits) - limits
	limits = pageNO * limits

	inventories, err := r.repo.GetInventory(offSet, limits)
	if err != nil {
		return nil, err
	}

	return inventories, nil
}

func (r *inventoryUseCase) GetAInventory(productID string) (*[]responsemodel.InventoryRes, error) {
	inventory, err := r.repo.GetAInventory(productID)
	if err != nil {
		return nil, err
	}
	return inventory, nil
}

func (r *inventoryUseCase) GetSellerInventory(page string, limit string, sellerID string) (*[]responsemodel.InventoryShowcase, error) {

	pageNO, err := strconv.Atoi(page)
	if err != nil {
		return nil, resCustomError.ErrConversionOFPage
	}

	limits, err := strconv.Atoi(limit)
	if err != nil {
		return nil, resCustomError.ErrConversionOfLimit
	}

	if pageNO < 1 {
		return nil, resCustomError.ErrPagination
	}

	if limits <= 0 {
		return nil, resCustomError.ErrPageLimit
	}

	offSet := (pageNO * limits) - limits
	limits = pageNO * limits

	inventories, err := r.repo.GetSellerInventory(offSet, limits, sellerID)
	if err != nil {
		return nil, err
	}

	return inventories, nil
}
