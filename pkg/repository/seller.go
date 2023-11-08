package repository

import (
	interfaces "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/repository/interface"
	"gorm.io/gorm"
)

type sellerRepository struct{
	DB *gorm.DB
}

func NewSellerHandler(db *gorm.DB) interfaces.ISellerRepo{
	return &sellerRepository{DB: db}
}

