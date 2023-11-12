package repository

import (
	"errors"

	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
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

// func (d adminRepository) GetUser(page string, limit string) error {
// 	offset := int(limit)
// 	query := "SELECT id, name, email, phone, status FROM user_details ORDER BY name OFFSET ? LIMIT ?"
// 	data, err := d.DB.Raw(query).Rows()
// 	if err != nil {
// 		return nil, errors.New("error for fetching user details")
// 	}

// }

func (d *adminRepository) AllUsers(offSet int, limit int) (*[]responsemodel.UserDetails, error) {
	var users []responsemodel.UserDetails

	query := "SELECT * FROM user_details ORDER BY name OFFSET ? LIMIT ?"
	err := d.DB.Raw(query, offSet, limit).Scan(&users).Error
	if err != nil {
		return nil, errors.New("can't get user data from db")
	}

	return &users, nil
}

func (d *adminRepository) UserCount(ch chan int) {
	var count int

	query := "SELECT COUNT(phone) FROM user_details WHERE status!='delete'"
	d.DB.Raw(query).Scan(&count)
	ch <- count
}

func (d *adminRepository) BlockUser(id string) error {
	query := "UPDATE user_details SET status = 'block' WHERE id=? "
	err := d.DB.Exec(query, id).Error
	if err != nil {
		return errors.New("block user process , is not satisfied")
	}
	return nil
}
