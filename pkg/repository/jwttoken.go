package repository

import (
	"errors"

	interfaces "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/repository/interface"
	"gorm.io/gorm"
)

type jwtTokenRepository struct {
	DB *gorm.DB
}

func NewJwtTokenRepository(db *gorm.DB) interfaces.IJwtTokenRepository {
	return &jwtTokenRepository{DB: db}
}

func (d *jwtTokenRepository) GetUserStatus(id string) (string, error) {
	var status string
	query := "SELECT status FROM seller WHERE id=?"
	err := d.DB.Raw(query, id).Row().Scan(&status)
	if err != nil {
		return "", errors.New("can't fetch data from database for crating jwt token")
	}
	return status, nil
}
