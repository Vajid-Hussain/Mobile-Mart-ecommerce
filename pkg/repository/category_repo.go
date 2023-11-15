package repository

import (
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
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

func (d *categoryRepository) InsertCategory(name *requestmodel.Category) error {
	query := "INSERT INTO categories(name) VALUES(?)"
	err := d.DB.Exec(query, name.Name).Error
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

func (d *categoryRepository) EditCategoryName(category *requestmodel.CategoryDetails) error {
	query := "UPDATE categories SET name=? WHERE id=?"
	result := d.DB.Exec(query, category.Name, category.ID)

	if result.RowsAffected == 0 {
		return errors.New("no category exist by id, do't duplicate brand")
	}
	if result.Error != nil {
		return errors.New("some problem from database for update category")
	}
	return nil
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
