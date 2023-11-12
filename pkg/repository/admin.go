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
	err := d.DB.Exec(query, id)
	if err.Error != nil {
		return errors.New("block user process , is not satisfied")
	}
	count := err.RowsAffected
	if count <= 0 {
		return errors.New("no user exist by id ")
	}
	return nil
}

func (d *adminRepository) UnblockUser(id string) error {
	query := "UPDATE user_details SET status = 'active' WHERE id=?"
	err := d.DB.Exec(query, id)
	if err.Error != nil {
		return errors.New("active user process , is not satisfied")
	}

	if err.RowsAffected <= 0 {
		return errors.New("no user exist by id ")
	}
	return nil
}

// Sellers

func (d *adminRepository) AllSellers(offSet int, limit int) (*[]responsemodel.SellerDetails, error) {
	var sellers []responsemodel.SellerDetails

	query := "SELECT * FROM sellers ORDER BY name OFFSET ? LIMIT ?"
	err := d.DB.Raw(query, offSet, limit).Scan(&sellers).Error
	if err != nil {
		return nil, errors.New("can't get seller data from db")
	}

	return &sellers, nil
}

func (d *adminRepository) SellerCount(ch chan int) {
	var count int

	query := "SELECT COUNT(email) FROM sellers WHERE status!='delete'"
	d.DB.Raw(query).Scan(&count)
	ch <- count
}

func (d *adminRepository) BlockSeller(id string) error {
	query := "UPDATE sellers SET status = 'block' WHERE id=? "
	err := d.DB.Exec(query, id)
	if err.Error != nil {
		return errors.New("block seller process , is not satisfied")
	}
	count := err.RowsAffected
	if count <= 0 {
		return errors.New("no seller exist by id ")
	}
	return nil
}

func (d *adminRepository) UnblockSeller(id string) error {
	query := "UPDATE sellers SET status = 'active' WHERE id=?"
	err := d.DB.Exec(query, id)
	if err.Error != nil {
		return errors.New("active seller process , is not satisfied")
	}

	if err.RowsAffected <= 0 {
		return errors.New("no seller exist by id ")
	}
	return nil
}

func (d *adminRepository) GetPendingSellers(offSet int, limit int) (*[]responsemodel.SellerDetails, error) {
	var sellers []responsemodel.SellerDetails

	query := "SELECT * FROM sellers WHERE status='pending' ORDER BY name OFFSET ? LIMIT ?"
	err := d.DB.Raw(query, offSet, limit).Scan(&sellers).Error
	if err != nil {
		return nil, errors.New("can't get pending list of sellers")
	}

	return &sellers, nil
}

func (d *adminRepository) GetSingleSeller(id string) (*responsemodel.SellerDetails, error) {
	var seller responsemodel.SellerDetails
	query := "SELECT * FROM sellers WHERE id= ?"
	err := d.DB.Raw(query, id).Scan(&seller)
	if err.Error != nil {
		return nil, errors.New("something wrong at fetching seller details")
	}

	if err.RowsAffected <= 0 {
		return nil, errors.New("no seller exist by id ")
	}

	return &seller, nil
}
