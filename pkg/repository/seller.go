package repository

import (
	"errors"

	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
	interfaces "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/repository/interface"
	"gorm.io/gorm"
)

type sellerRepository struct {
	DB *gorm.DB
}

func NewSellerRepository(db *gorm.DB) interfaces.ISellerRepo {
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

	query := "INSERT INTO sellers (name, email, password, gst_no, description) VALUES($1, $2, $3, $4, $5)"
	err := d.DB.Exec(query, SellerData.Name, SellerData.Email, SellerData.Password, SellerData.GST_NO, SellerData.Description).Error
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

func (d *sellerRepository) AllSellers(offSet int, limit int) (*[]responsemodel.SellerDetails, error) {
	var sellers []responsemodel.SellerDetails

	query := "SELECT * FROM sellers ORDER BY name OFFSET ? LIMIT ?"
	err := d.DB.Raw(query, offSet, limit).Scan(&sellers).Error
	if err != nil {
		return nil, errors.New("can't get seller data from db")
	}

	return &sellers, nil
}

func (d *sellerRepository) SellerCount(ch chan int) {
	var count int

	query := "SELECT COUNT(email) FROM sellers WHERE status!='delete'"
	d.DB.Raw(query).Scan(&count)
	ch <- count
}

func (d *sellerRepository) BlockSeller(id string) error {
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

func (d *sellerRepository) UnblockSeller(id string) error {
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

func (d *sellerRepository) GetPendingSellers(offSet int, limit int) (*[]responsemodel.SellerDetails, error) {
	var sellers []responsemodel.SellerDetails

	query := "SELECT * FROM sellers WHERE status='pending' ORDER BY name OFFSET ? LIMIT ?"
	err := d.DB.Raw(query, offSet, limit).Scan(&sellers).Error
	if err != nil {
		return nil, errors.New("can't get pending list of sellers")
	}

	return &sellers, nil
}

func (d *sellerRepository) GetSingleSeller(id string) (*responsemodel.SellerDetails, error) {
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