package db

import (
	"fmt"

	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/config"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/domain"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/utils/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(config config.DataBase) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", config.DBHost, config.DBUser, config.DBName, config.DBPort, config.DBPassword)
	DB, dberr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if dberr != nil {
		return DB, dberr
	}

	// Table Creation
	if err := DB.AutoMigrate(&domain.Users{}); err != nil {
		return DB, err
	}
	if err := DB.AutoMigrate(&domain.Seller{}); err != nil {
		return DB, err
	}
	if err := DB.AutoMigrate(&domain.Admin{}); err != nil {
		return DB, err
	}
	if err := DB.AutoMigrate(&domain.Category{}, &domain.Brand{}); err != nil {
		return DB, err
	}
	if err := DB.AutoMigrate(&domain.Inventories{}); err != nil {
		return DB, err
	}
	if err := DB.AutoMigrate(&domain.Address{}); err != nil {
		return DB, err
	}
	if err := DB.AutoMigrate(&domain.Cart{}); err != nil {
		return DB, err
	}
	if err := DB.AutoMigrate(&domain.Order{}); err != nil {
		return DB, err
	}
	if err := DB.AutoMigrate(&domain.OrderProducts{}); err != nil {
		return DB, err
	}
	if err := DB.AutoMigrate(&domain.Wallet{}); err != nil {
		return DB, err
	}
	if err := DB.AutoMigrate(&domain.CategoryOffer{}); err != nil {
		return DB, err
	}
	if err := DB.AutoMigrate(&domain.WalletTransaction{}); err != nil {
		return DB, err
	}
	if err := DB.AutoMigrate(&domain.Coupons{}); err != nil {
		return DB, err
	}

	CheckAndCreateAdmin(DB)

	return DB, nil
}

func CheckAndCreateAdmin(DB *gorm.DB) {
	var count int
	var (
		Name     = "mobileMart"
		Email    = "mobilemart@gmail.com"
		Password = "buyMobiles"
	)
	HashedPassword := helper.HashPassword(Password)

	query := "SELECT COUNT(*) FROM admins"
	DB.Raw(query).Row().Scan(&count)
	if count <= 0 {
		query = "INSERT INTO admins(name, email, password) VALUES(?, ?, ?)"
		DB.Exec(query, Name, Email, HashedPassword).Row().Err()
	}
}
