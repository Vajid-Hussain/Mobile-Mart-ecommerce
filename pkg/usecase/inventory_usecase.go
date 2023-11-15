package usecase

import (
	"errors"
	"fmt"

	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
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
