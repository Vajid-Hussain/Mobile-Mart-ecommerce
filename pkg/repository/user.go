package repository

import (
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	interfaces "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/repository/interface"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.IUserRepo {
	return &userRepository{DB: DB}
}

func (d *userRepository) CreateUser(userDetails requestmodel.UserDetails) {
	query := "INSERT INTO user_details (name, email, phone, password) VALUES($1, $2, $3, $4)"
	d.DB.Exec(query, userDetails.Name, userDetails.Email, userDetails.Phone, userDetails.Password)
}
