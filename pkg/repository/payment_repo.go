package repository

import (
	"errors"
	"fmt"

	resCustomError "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel/custom_error"
	interfaces "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/repository/interface"
	"gorm.io/gorm"
)

type paymentRepo struct {
	DB *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) interfaces.IPaymentRepository {
	return &paymentRepo{DB: db}
}

func (d *paymentRepo) CreateOrUpdateWallet(userID string, creditAmount uint) (*uint, error) {

	var currentBalance uint
	query := "INSERT INTO wallets (user_id, balance) VALUES ($1, $2) ON CONFLICT(user_id) DO UPDATE SET balance=wallets.balance + $2"
	result := d.DB.Exec(query, userID, creditAmount)
	if result.Error != nil {
		return nil, errors.New("face some issue while intract with wallet for made change in balance")
	}
	if result.RowsAffected == 0 {
		return nil, resCustomError.ErrNoRowAffected
	}
	return &currentBalance, nil
}

func (d *paymentRepo) GetWalletbalance(userID string) (*uint, error) {

	var currentBalance uint
	query := "SELECT balance FROM wallet WHERE user_id= ?"
	result := d.DB.Raw(query, userID).Scan(&currentBalance)
	if result.Error != nil {
		return nil, errors.New("face some issue while fetch user wallet balance")
	}
	if result.RowsAffected == 0 {

		return nil, resCustomError.ErrNoRowAffected
	}
	return &currentBalance, nil
}
