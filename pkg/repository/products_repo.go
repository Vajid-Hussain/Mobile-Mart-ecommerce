package repository

import (
	"errors"
	"fmt"

	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
	resCustomError "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel/custom_error"
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

	query := `INSERT INTO inventories (productname, description, brand_id, category_id, mrp, saleprice, units,os, cellular_technology, ram, screensize, Batterycapacity, processor, seller_id, image_url, discount) VALUES (?, ?,  ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) RETURNING *;`
	result := d.DB.Raw(query,
		inventory.Productname, inventory.Description, inventory.BrandID, inventory.CategoryID,
		inventory.Mrp, inventory.Saleprice, inventory.Units,
		inventory.Os, inventory.CellularTechnology, inventory.Ram,
		inventory.Screensize, inventory.Batterycapacity, inventory.Processor, inventory.SellerID, inventory.ImageURL, inventory.Discount,
	).Scan(&insertedData)

	if result.Error != nil {
		return nil, errors.New("inventory is not added into database")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("inventory is not added in database , face some error")
	}
	return &insertedData, nil
}

func (d *inventoryRepository) BlockSingleInventoryBySeller(SellerID string, productID string) error {
	query := "UPDATE inventories SET status='block' WHERE id= $1"
	err := d.DB.Exec(query, productID).Error
	if err != nil {
		return errors.New("can't change the status of product")
	}
	return nil
}

func (d *inventoryRepository) UNBlockSingleInventoryBySeller(SellerID string, productID string) error {
	query := "UPDATE inventories SET status='active' WHERE id= $1"
	err := d.DB.Exec(query, productID).Error
	if err != nil {
		return errors.New("can't change the status of product in inverntories")
	}
	return nil
}

func (d *inventoryRepository) DeleteInventoryBySeller(SellerID string, productID string) error {
	query := "UPDATE inventories SET status='delete' WHERE id= $1"
	result := d.DB.Exec(query, productID)
	if result.Error != nil {
		return errors.New("can't change the status of product in inverntories")
	}
	if result.RowsAffected == 0 {
		return errors.New("no inventory exist in table for deletion")
	}
	return nil
}

func (d *inventoryRepository) GetInventory(offSet int, limit int) (*[]responsemodel.InventoryShowcase, error) {
	var inventory []responsemodel.InventoryShowcase

	query := "SELECT * FROM inventories WHERE status = 'active' ORDER BY id OFFSET ? LIMIT ?"
	err := d.DB.Raw(query, offSet, limit).Scan(&inventory).Error
	if err != nil {
		return nil, errors.New("can't get inventory data from db")
	}

	return &inventory, nil
}

func (d *inventoryRepository) GetAInventory(id string) (*responsemodel.InventoryRes, error) {
	var inventory responsemodel.InventoryRes

	query := "SELECT * FROM inventories WHERE id=? AND status='active'"
	result := d.DB.Raw(query, id).Scan(&inventory)
	if result.Error != nil {
		return nil, errors.New("can't get inventory data from db or inventory is not active state")
	}

	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrNoRowAffected
	}

	return &inventory, nil
}

func (d *inventoryRepository) GetSellerInventory(offSet int, limit int, sellerID string) (*[]responsemodel.InventoryShowcase, error) {
	var inventory []responsemodel.InventoryShowcase

	query := "SELECT * FROM inventories WHERE seller_id= ? AND status = 'active' ORDER BY id OFFSET ? LIMIT ?"
	err := d.DB.Raw(query, sellerID, offSet, limit).Scan(&inventory).Error
	if err != nil {
		return nil, errors.New("can't get inventory data from db")
	}

	return &inventory, nil
}

func (d *inventoryRepository) UpdateInventory(inventory *requestmodel.EditInventory) (*responsemodel.InventoryRes, error) {
	var updatedData responsemodel.InventoryRes

	// query := `UPDATE inventories
	// SET productname = ?, description = ?, brand_id = ?, category_id = ?,
	// 	mrp = ?, saleprice = ?, units = ?, os = ?, cellular_technology = ?,
	// 	ram = ?, screensize = ?, batterycapacity = ?, processor = ?
	// WHERE id = ? RETURNING *;`

	// result := d.DB.Raw(query,
	// 	inventory.Productname, inventory.Description, inventory.BrandID, inventory.CategoryID,
	// 	inventory.Mrp, inventory.Saleprice, inventory.Units,
	// 	inventory.Os, inventory.CellularTechnology, inventory.Ram,
	// 	inventory.Screensize, inventory.Batterycapacity, inventory.Processor,
	// 	inventory.ID,
	// ).Scan(&updatedData)

	query := "UPDATE inventories SET mrp=?, discount= ?, saleprice= ?, units= ? WHERE id=? RETURNING*;"
	result := d.DB.Raw(query, inventory.Mrp, inventory.Discount, inventory.Saleprice, inventory.Units, inventory.ID).Scan(&updatedData)
	if result.Error != nil {
		return nil, errors.New("inventory is not updated into database")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("inventory is not updated in database , face some error")
	}
	return &updatedData, nil
}

func (d *inventoryRepository) GetProductFilter(criterion *requestmodel.FilterCriterion) (*[]responsemodel.FilterProduct, error) {
	fmt.Println("##", criterion.MinPrice)
	var sortedProduct []responsemodel.FilterProduct

	query := "SELECT inventories.id AS productID, * FROM inventories INNER JOIN categories ON categories.id= inventories.category_id INNER JOIN brands ON brands.id= inventories.brand_id WHERE categories.name ILIKE '%' || $1 || '%' AND brands.name ILIKE '%' || $2 || '%' AND inventories.productname ILIKE '%' || $3 || '%' AND ($4 = 0 OR $4 < inventories.saleprice AND ($5 = 0 OR $5 >= inventories.saleprice))"
	result := d.DB.Raw(query, criterion.Category, criterion.Brand, criterion.Product, criterion.MinPrice, criterion.MaxPrice).Scan(&sortedProduct)
	if result.Error != nil {
		return nil, errors.New("face some issue while filter product")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrNoRowAffected
	}
	return &sortedProduct, nil
}
