package repository

import (
	"errors"
	"fmt"

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

func (d *adminRepository) GetSellerDetailsForDashBord(criteria string) (uint, error) {
	var data uint

	query := "SELECT COUNT(*) FROM sellers WHERE status= $1 OR $1 = '' "
	result := d.DB.Raw(query, criteria).Scan(&data)

	if result.Error != nil {
		return 0, errors.New("face some issue while get report by days")
	}
	return data, nil
}

// func (d *adminRepository) TotalRevenue() (uint, uint, error) {

// 	var orders, netrevenue uint
// 	query := "SELECT COALESCE(COUNT(*),0), COALESCE(SUM(price),0) FROM orders"
// 	result := d.DB.Raw(query).Row().Scan(&orders, &netrevenue).Error()
// 	if result != "" {
// 		return 0, 0, errors.New("face some issue while get report by days")
// 	}
// 	return orders, netrevenue, nil
// }

func (d *adminRepository) TotalRevenue() (uint, uint, error) {
	var count, sum uint
	query := "SELECT COALESCE(COUNT(*), 0), COALESCE(SUM(price), 0) FROM orders WHERE order_status='delivered'"
	result := d.DB.Raw(query).Row().Scan(&count, &sum)
	if result != nil {
		return 0, 0, errors.New("face some issue while get report by days")
	}

	return count, sum, nil
}

// COALESCE

func (d *adminRepository) GetNetCredit() (uint, error) {
	var credit uint
	query := "SELECT COALESCE(SUM(seller_credit),0) FROM sellers"
	result := d.DB.Raw(query).Scan(&credit)
	if result.Error != nil {
		return 0, errors.New("face some issue while get report by days")
	}
	fmt.Println("$$", credit)
	return credit, nil
}
