package usecase

import (
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/config"
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
	interfaces "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/repository/interface"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/service"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/utils/helper"
	"github.com/go-playground/validator/v10"
)

type inventoryUseCase struct {
	repo interfaces.IInventoryRepository
	s3   config.S3Bucket
}

func NewInventoryUseCase(repository interfaces.IInventoryRepository, s3aws *config.S3Bucket) interfaceUseCase.IInventoryUseCase {
	return &inventoryUseCase{repo: repository,
		s3: *s3aws}
}

func (d *inventoryUseCase) AddInventory(inventory *requestmodel.InventoryReq) (*responsemodel.InventoryRes, error) {

	sess := service.CreateSession(&d.s3)

	ImageURL, err := service.UploadImageToS3(inventory.Image, sess)
	if err != nil {
		return nil, err
	}

	inventory.ImageURL = ImageURL

	product, err := d.repo.CreateProduct(inventory)
	if err != nil {
		return nil, err
	}

	return product, nil
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

	offSet, limits, err := helper.Pagination(page, limit)
	if err != nil {
		return nil, err
	}

	inventories, err := r.repo.GetInventory(offSet, limits)
	if err != nil {
		return nil, err
	}

	return inventories, nil
}

func (r *inventoryUseCase) GetAInventory(productID string) (*responsemodel.InventoryRes, error) {
	inventory, err := r.repo.GetAInventory(productID)
	if err != nil {
		return nil, err
	}
	return inventory, nil
}

func (r *inventoryUseCase) GetSellerInventory(page string, limit string, sellerID string) (*[]responsemodel.InventoryShowcase, error) {

	offSet, limits, err := helper.Pagination(page, limit)
	if err != nil {
		return nil, err
	}

	inventories, err := r.repo.GetSellerInventory(offSet, limits, sellerID)
	if err != nil {
		return nil, err
	}

	return inventories, nil
}

func (r *inventoryUseCase) EditInventory(editInventory *requestmodel.EditInventory, invetoryID string) (*responsemodel.InventoryRes, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	inventory, err := r.repo.GetAInventory(invetoryID)
	if err != nil {
		return nil, err
	}

	err = validate.Struct(editInventory)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				fieldName := e.Field()
				switch fieldName {
				case "ID":
					editInventory.ID = inventory.ID
				case "Productname":
					editInventory.Productname = inventory.Productname
				case "Description":
					editInventory.Description = inventory.Description
				case "BrandID":
					editInventory.BrandID = inventory.BrandID
				case "CategoryID":
					editInventory.CategoryID = inventory.CategoryID
				case "SellerID":
					editInventory.SellerID = inventory.SellerID
				case "Mrp":
					editInventory.Mrp = inventory.Mrp
				case "Saleprice":
					editInventory.Saleprice = inventory.Saleprice
				case "Units":
					editInventory.Units = inventory.Units
				case "Os":
					editInventory.Os = inventory.Os
				case "CellularTechnology":
					editInventory.CellularTechnology = inventory.CellularTechnology
				case "Ram":
					editInventory.Ram = inventory.Ram
				case "Screensize":
					editInventory.Screensize = inventory.Screensize
				case "Batterycapacity":
					editInventory.Batterycapacity = inventory.Batterycapacity
				case "Processor":
					editInventory.Processor = inventory.Processor
				}
			}
		}
	}

	updatedData, err := r.repo.UpdateInventory(editInventory)
	if err != nil {
		return nil, err
	}

	return updatedData, nil
}
