package repository

import (
	"errors"

	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
	interfaces "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/repository/interface"
	"gorm.io/gorm"
)

type inventoryRepository struct {
	DB *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) interfaces.IInventoryRepository {
	return &inventoryRepository{DB: db}
}

func (d *inventoryRepository) CreateProduct(inventory *requestmodel.InventoryReq) (*responsemodel.InventoryRes, error) {
	var insertedData responsemodel.InventoryRes

	query := `INSERT INTO inventories (productname, description, brand_id, category_id, mrp, saleprice, units,os, cellular_technology, ram, screensize, Batterycapacity, processor, seller_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) RETURNING *;`
	result := d.DB.Raw(query,
		inventory.Productname, inventory.Description, inventory.BrandID, inventory.CategoryID,
		inventory.Mrp, inventory.Saleprice, inventory.Units,
		inventory.Os, inventory.CellularTechnology, inventory.Ram,
		inventory.Screensize, inventory.Batterycapacity, inventory.Processor, inventory.SellerID,
	).Scan(&insertedData)

	if result.Error != nil {
		return nil, errors.New("inventory is not added into database")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("inventory is not added in database , face some error")
	}
	return &insertedData, nil
}
