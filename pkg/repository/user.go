package repository

import (
	"errors"
	"fmt"

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

//user Repository

func (d *userRepository) CreateUser(userDetails *requestmodel.UserDetails) {
	query := "INSERT INTO user_details (id, name, email, phone, password) VALUES($1, $2, $3, $4, $5)"
	d.DB.Exec(query,userDetails.Id, userDetails.Name, userDetails.Email, userDetails.Phone, userDetails.Password)
}

func (d *userRepository) IsUserExist(userDetails *requestmodel.UserDetails) int {
	var userCount int

	query := "SELECT COUNT(*) FROM user_details WHERE email=?"
	err := d.DB.Raw(query, userDetails.Email).Row().Scan(&userCount)
	if err != nil {
		fmt.Println("Error for user exist using email in signup")
	}

	return userCount
}

func (d *userRepository) CheckUserByPhone(phone string)error{
	var count int

	query:="SELECT COUNT(*) FROM user_details WHERE Phone=$1"
	result:=d.DB.Raw(query,phone).Row()
	result.Scan(&count)
	if count>=1{
		return errors.New("no user Exist , phone number is wrong")
	}else{
		return nil
	}
}