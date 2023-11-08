package repository

import (
	"database/sql"
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
	d.DB.Exec(query, userDetails.Id, userDetails.Name, userDetails.Email, userDetails.Phone, userDetails.Password)
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

func (d *userRepository) ChangeUserStatusActive(phone string) error {
	fmt.Println(phone)
	query := "UPDATE user_details SET status = 'active' WHERE phone = ?"
	result := d.DB.Exec(query, phone)
	// count:=result.RowsAffected

	if result.Error != nil {
		return errors.New("no user Exist , phone number is wrong")
	} else {
		return nil
	}
}

func (d *userRepository) FetchUserID(phone string) (string, error) {
	var userID string

	query := "SELECT id FROM user_details WHERE phone=? AND status='active'"
	data := d.DB.Raw(query, phone).Row()

	if err := data.Scan(&userID); err != nil {
		return "", errors.New("fetching user id cause error")
	}
	return userID, nil
}

func (d *userRepository) FetchPasswordUsingPhone(phone string) (string, error) {
	var password string

	query := "SELECT password FROM user_details WHERE phone=? AND status='active'"
	row := d.DB.Raw(query, phone).Row()

	if row == nil {
		return "", errors.New("no rows returned from the query")
	}

	err := row.Scan(&password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("user does not exist")
		}
		return "", fmt.Errorf("error scanning row: %s", err)
	}
	return password, nil
}
