package repository

import "gorm.io/gorm"

type paymentRepo struct {
	DB *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) {
	return paymentRepo{DB: db}
}

func (d *paymentRepo) CreateOrUpdateWallet(userID string, amount uint) {
	query := "INSERT INTO wallet (user_id, belance) VALUES ($1, $2) "
}
