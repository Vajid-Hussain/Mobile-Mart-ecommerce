package repository

import (
	"errors"

	interfaces "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/repository/interface"
	"gorm.io/gorm"
)

type adminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) interfaces.IAdminRepository {
	return &adminRepository{DB: db}
}

func (d *adminRepository) GetPassword(email string) (string, error) {
	var hashedPassword string

	query := "SELECT password FROM admins WHERE email=?"
	err := d.DB.Raw(query, email).Row().Scan(&hashedPassword)
	if err != nil {
		return "", errors.New("error at admin password fetch")
	}
	return hashedPassword, nil
}

// func (d *adminRepository) GetSellerDetailsForDashBord(criteria string) (uint, error) {
// 	var data uint

// 	query := "SELECT COUNT(*) FROM sellers WHERE status= "+ criteria
// 	d.DB.Raw(query).Scan(&data)

// 	if result.Error != nil {
// 		return nil, errors.New("face some issue while get report by days")
// 	}
// 	if result.RowsAffected == 0 {
// 		return nil, resCustomError.ErrNoRowAffected
// 	}
// }
