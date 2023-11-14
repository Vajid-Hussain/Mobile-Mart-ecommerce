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

	query := "SELECT * FROM categories ORDER BY id OFFSET ? LIMIT ?"
	err := d.DB.Raw(query, offSet, limit).Scan(&categories).Error
	if err != nil {
		return nil, errors.New("can't fetch categories from database")
	}

	return &categories, nil
}

// func (d *categoryRepository) EditCategory(id int, name string) error {
// 	query := "UPDATE categories SET name= ? WHERE id= ?"
// 	err := d.DB.Exec(query, name, id).Error
// 	if err != nil {
// 		return errors.New("updating category,  facing an error, confirm id of category too")
// 	}
// 	return nil
// }

func (d *categoryRepository) EditCategoryName(category *requestmodel.CategoryDetails) error {
	query := "UPDATE categories SET name=? WHERE id=?"
	result := d.DB.Exec(query, category.Name, category.ID)

	if result.RowsAffected == 0 {
		return errors.New("no category exist by id")
	}
	if result.Error != nil {
		return errors.New("some problem from database for update category")
	}
	return nil
}
