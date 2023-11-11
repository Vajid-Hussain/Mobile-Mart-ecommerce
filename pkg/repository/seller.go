package repository

import (
	"errors"

	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	interfaces "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/repository/interface"
	"gorm.io/gorm"
)

type sellerRepository struct {
	DB *gorm.DB
}

func NewSellerHandler(db *gorm.DB) interfaces.ISellerRepo {
	return &sellerRepository{DB: db}
}

func (d *sellerRepository) IsSellerExist(email string) (int, error) {

	var sellerCount int

	query := "SELECT COUNT(*) FROM sellers WHERE email=$1 AND status!=$2"
	err := d.DB.Raw(query, email, "delete").Row().Scan(&sellerCount)
	if err != nil {
		return 0, errors.New("error at fetching seller count using email")
	}
	return sellerCount, nil
}

func (d *sellerRepository) CreateSeller(SellerData *requestmodel.SellerSignup) error {

	query := "INSERT INTO sellers (id, name, email, password, gst_no, description) VALUES($1, $2, $3, $4, $5, $6)"
	err := d.DB.Exec(query, SellerData.ID, SellerData.Name, SellerData.Email, SellerData.Password, SellerData.GST_NO, SellerData.Description).Error
	if err != nil {
		return errors.New("seller data not save in database face some issue")
	}

	return nil
}

func (d *sellerRepository) GetHashPassAndStatus(email string) (string, string, string, error) {
	var password, status, sellerID string
	query := "SELECT password, id, status FROM sellers WHERE email=? AND status!='delete'"
	err := d.DB.Raw(query, email).Row().Scan(&password, &sellerID, &status)
	if err != nil {
		return "", "", "", errors.New("feching password and status, can't make action in database")
	}
	return password, sellerID, status, nil
}

func (d *sellerRepository) GetPasswordByMail(email string) string {
	var hashedPassword string
	query := "SELECT password FROM sellers WHERE email=? AND status='active'"
	err := d.DB.Raw(query, email).Row().Scan(&hashedPassword)
	if err != nil {
		return ""
	}
	return hashedPassword
}
