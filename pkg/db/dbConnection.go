package db

import (
	"fmt"

	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/config"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(config config.DataBase) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", config.DBHost, config.DBUser, config.DBName, config.DBPort, config.DBPassword)
	DB, dberr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if dberr != nil {
		return DB, nil
	}

	if err:=DB.AutoMigrate(&domain.UserDetails{}); err!=nil{
		return DB, err
	}
	if err:=DB.AutoMigrate(&domain.Seller{}); err!=nil{
		return DB, err
	}

	return DB, nil
}
