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

func (d *userRepository) IsUserExist(phone string) int {
	var userCount int

	query := "SELECT COUNT(*) FROM user_details WHERE phone=$1 AND status!=$2"
	err := d.DB.Raw(query, phone, "delete").Row().Scan(&userCount)
	if err != nil {
		fmt.Println("Error for user exist, using same phone in signup")
	}
	return userCount

}

func (d *userRepository) CheckUserByPhone(phone string)error{

	query:="UPDATE user_details SET status=$2 WHERE phone=$1"
	result:=d.DB.Raw(query, "active", phone).Error

	if result!=nil{
		return errors.New("no user Exist , phone number is wrong")
	}else{
		return nil
	}

}
