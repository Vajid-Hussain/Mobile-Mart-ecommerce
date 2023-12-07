package repository

import (
	"time"

	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
	resCustomError "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel/custom_error"
	interfaces "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/repository/interface"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type categoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) interfaces.ICategoryRepository {
	return &categoryRepository{DB: db}
}

func (d *categoryRepository) InsertCategory(categoryDetails *requestmodel.Category) error {
	query := "INSERT INTO categories(name) VALUES(?)"
	err := d.DB.Exec(query, categoryDetails.Name).Error
	if err != nil {
		return errors.New("canot make a new Category")
	}
	return nil
}

func (d *categoryRepository) GetAllCategory(offSet int, limit int) (*[]responsemodel.CategoryDetails, error) {
	var categories []responsemodel.CategoryDetails

	query := "SELECT * FROM categories WHERE status!='delete' ORDER BY name OFFSET ? LIMIT ?"
	err := d.DB.Raw(query, offSet, limit).Scan(&categories).Error
	if err != nil {
		return nil, errors.New("can't fetch categories from database")
	}

	return &categories, nil
}

func (d *categoryRepository) EditCategoryName(category *requestmodel.CategoryDetails) (*responsemodel.CategoryDetails, error) {
	var updatedCategory responsemodel.CategoryDetails
	query := "UPDATE categories SET name=? WHERE id=? RETURNING*"
	result := d.DB.Raw(query, category.Name, category.ID).Scan(&updatedCategory)

	if result.RowsAffected == 0 {
		return nil, errors.New("no category exist by id, do't duplicate brand")
	}
	if result.Error != nil {
		return nil, errors.New("some problem from database for update category")
	}
	return &updatedCategory, nil
}

func (d *categoryRepository) DeleteCategory(id string) error {
	query := "UPDATE categories SET status='delete' WHERE id= $1"
	err := d.DB.Exec(query, id).Error
	if err != nil {
		return errors.New("can't delete category from database")
	}
	return nil
}

// Brand
func (d *categoryRepository) InsertBrand(name *requestmodel.Brand) error {
	query := "INSERT INTO brands(name) VALUES(?)"
	err := d.DB.Exec(query, name.Name).Error
	if err != nil {
		return errors.New("can't make a new brand")
	}
	return nil
}

func (d *categoryRepository) GetAllBrand(offSet int, limit int) (*[]responsemodel.BrandRes, error) {
	var Brands []responsemodel.BrandRes

	query := "SELECT * FROM brands WHERE status!='delete' ORDER BY name OFFSET ? LIMIT ?"
	err := d.DB.Raw(query, offSet, limit).Scan(&Brands).Error
	if err != nil {
		return nil, errors.New("can't fetch brand from database")
	}

	return &Brands, nil
}

func (d *categoryRepository) EditBrandName(brand *requestmodel.BrandDetails) error {

	query := "UPDATE brands SET name=? WHERE id=?"
	result := d.DB.Exec(query, brand.Name, brand.ID)
	if result.RowsAffected == 0 {
		return errors.New("no brands exist by id")
	}
	if result.Error != nil {
		return errors.New("some problem from database for update brand")
	}
	return nil
}

func (d *categoryRepository) DeleteBrand(id string) error {

	query := "UPDATE brands SET status='delete' WHERE id= $1"
	err := d.DB.Exec(query, id).Error
	if err != nil {
		return errors.New("can't delete brand from database")
	}
	return nil
}

// inventory

func (d *categoryRepository) DeleteInventoryOfCategory(id string) error {
	query := "UPDATE inventories SET status= 'delete' WHERE category_id= $1 AND status!='delete'"
	err := d.DB.Exec(query, id).Error
	if err != nil {
		return errors.New("blocking product of blocked brand can't achive to change")
	}
	return nil
}

func (d *categoryRepository) DeleteInventoryOfBrand(id string) error {
	query := "UPDATE inventories SET status= 'delete' WHERE brand_id= $1 AND status!='delete'"
	err := d.DB.Exec(query, id).Error
	if err != nil {
		return errors.New("blocking product of blocked brand can't achive to change")
	}
	return nil
}

// category offer
func (d *categoryRepository) InsertCategoryOffer(categoryOffer *requestmodel.CategoryOffer) (*responsemodel.CategoryOffer, error) {

	var categoryOfferres responsemodel.CategoryOffer
	offerExpireTime := time.Now().Add(time.Duration(categoryOffer.Validity) * 24 * time.Hour)

	query := "INSERT INTO category_offers(title, category_id, seller_id, category_discount, start_date, end_date) VALUES (?, ?, ?, ?, now(), ?) RETURNING *"
	result := d.DB.Raw(query, categoryOffer.Title, categoryOffer.CategoryID, categoryOffer.SellerID, categoryOffer.CategoryDiscount, offerExpireTime).Scan(&categoryOfferres)
	if result.RowsAffected == 0 {
		return nil, errors.New("category offer, row is not inserted")
	}
	if result.Error != nil {
		return nil, errors.New("no catogary exist")
	}
	return &categoryOfferres, nil
}

func (d *categoryRepository) ChekSellerHaveCategoryOffer(sellerID, categoryID string) (*uint, error) {
	var exist uint

	query := "SELECT COUNT(*) FROM category_offers WHERE seller_id=? AND category_id= ? AND status= 'active' AND end_date> now()"
	result := d.DB.Raw(query, sellerID, categoryID).Scan(&exist)
	if result.Error != nil {
		return nil, result.Error
	}
	return &exist, nil
}

func (d *categoryRepository) ChangeStatus(status, categoryOfferID string) (*responsemodel.CategoryOffer, error) {

	var categoryOffer responsemodel.CategoryOffer

	query := "UPDATE category_offers SET status= ? WHERE id=?  RETURNING *"
	result := d.DB.Raw(query, status, categoryOfferID).Scan(&categoryOffer)
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrNoRowAffected
	}
	if result.Error != nil {
		return nil, errors.New("face issue on change status of category offer")
	}
	return &categoryOffer, nil
}

func (d *categoryRepository) GetAllCategoryOffers(sellerID string) (*[]responsemodel.CategoryOffer, error) {

	var categoryOffers *[]responsemodel.CategoryOffer
	query := "SELECT * FROM category_offers WHERE seller_id=? AND end_date>=now() AND status='active'"
	result := d.DB.Raw(query, sellerID).Scan(&categoryOffers)
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrNoRowAffected
	}
	if result.Error != nil {
		return nil, errors.New("face issue on change status of category offer")
	}
	return categoryOffers, nil
}

func (d *categoryRepository) UpdateCategoryOffer(updateData *requestmodel.EditCategoryOffer) (*responsemodel.CategoryOffer, error) {

	var updatedCategoryOffer responsemodel.CategoryOffer
	query := "UPDATE category_offers SET title=$1, category_discount=$2, end_date = end_date + $3 * INTERVAL '1 day' WHERE id=$4 AND seller_id=$5 AND status='active' RETURNING*"
	result := d.DB.Raw(query, updateData.Title, updateData.CategoryDiscount, updateData.Validity, updateData.ID, updateData.SellerID).Scan(&updatedCategoryOffer)
	if result.Error != nil {
		return nil, errors.New("face issue while update category offer")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrNoRowAffected
	}
	return &updatedCategoryOffer, nil
}
