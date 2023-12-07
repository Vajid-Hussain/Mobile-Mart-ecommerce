package usecase

import (
	"errors"

	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/config"
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
	resCustomError "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel/custom_error"
	interfaces "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/repository/interface"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/service"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/utils/helper"
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
	discountedPrice := helper.FindDiscount(float64(inventory.Mrp), float64(inventory.Discount))
	inventory.Saleprice = discountedPrice

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

	for i, product := range *inventories {
		if product.CategoryDiscount != 0 {
			(*inventories)[i].NetDiscount = product.Discount + product.CategoryDiscount
			(*inventories)[i].PriceAfterApplyCategoryDiscount = helper.FindDiscount(float64(product.Mrp), float64((*inventories)[i].NetDiscount))
		}
	}
	return inventories, nil
}

func (r *inventoryUseCase) GetAInventory(productID string) (*responsemodel.InventoryRes, error) {
	inventory, err := r.repo.GetAInventory(productID)
	if err != nil {
		return nil, err
	}
	if inventory.CategoryDiscount != 0 {
		inventory.NetDiscount = inventory.CategoryDiscount + inventory.Discount
		inventory.FinalPrice = helper.FindDiscount(float64(inventory.Mrp), float64(inventory.NetDiscount))
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

func (r *inventoryUseCase) EditInventory(editInventory *requestmodel.EditInventory) (*responsemodel.InventoryRes, error) {

	inventory, err := r.repo.GetAInventory(editInventory.ID)
	if err != nil {
		return nil, err
	}
	if inventory.SellerID != editInventory.SellerID {
		return nil, resCustomError.ErrNoRowAffected
	}

	// fill data if it's empty
	if editInventory.Units == 0 {
		editInventory.Units = inventory.Units
	}

	if editInventory.Discount == 0 {
		editInventory.Discount = inventory.Discount
	}

	if editInventory.Mrp == 0 {
		editInventory.Mrp = inventory.Mrp
	}

	if editInventory.Saleprice == 0 {
		editInventory.Saleprice = inventory.Saleprice
	}

	// modify data to reach my criteria
	if editInventory.Mrp != 0 {
		editInventory.Saleprice = helper.FindDiscount(float64(editInventory.Mrp), float64(inventory.Discount))
	}

	if editInventory.Discount != 0 {
		if editInventory.Discount > 99 {
			return nil, errors.New("discount must be 1 to 99")
		}
		editInventory.Saleprice = helper.FindDiscount(float64(inventory.Mrp), float64(editInventory.Discount))
	}

	if editInventory.Mrp != 0 && editInventory.Discount != 0 {
		if editInventory.Discount > 99 {
			return nil, errors.New("discount must be 0 to 99")
		}
		editInventory.Saleprice = helper.FindDiscount(float64(editInventory.Mrp), float64(editInventory.Discount))
	}

	updatedData, err := r.repo.UpdateInventory(editInventory)
	if err != nil {
		return nil, err
	}

	return updatedData, nil
}

func (r *inventoryUseCase) GetProductFilter(criterion *requestmodel.FilterCriterion) (*[]responsemodel.FilterProduct, error) {
	filteredProduct, err := r.repo.GetProductFilter(criterion)
	if err != nil {
		return nil, err
	}
	for i, product := range *filteredProduct {
		if product.CategoryDiscount != 0 {
			(*filteredProduct)[i].NetDiscount = product.Discount + product.CategoryDiscount
			(*filteredProduct)[i].PriceAfterApplyCategoryDiscount = helper.FindDiscount(float64(product.Mrp), float64((*filteredProduct)[i].NetDiscount))
		}
	}
	return filteredProduct, nil
}
