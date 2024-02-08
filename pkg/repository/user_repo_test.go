package repository

import (
	"errors"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGetUserDetails(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		stub    func(sqlmock.Sqlmock)
		want    *requestmodel.UserDetails
		wantErr error
	}{
		{
			name: "succesully got details",
			args: "1",
			stub: func(mockSql sqlmock.Sqlmock) {
				mockSql.ExpectQuery("SELECT id, name , email, phone, referal_code FROM users WHERE id= ?").
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "phone", "referal_code"}).AddRow(1, "vajid", "vajid23@gmail.com", "8765434564", "kj678"))
			},
			want: &requestmodel.UserDetails{
				Id:          "1",
				Name:        "vajid",
				Email:       "vajid23@gmail.com",
				Phone:       "8765434564",
				ReferalCode: "kj678",
			},
			wantErr: nil,
		}, {
			name: "no user found",
			args: "1",
			stub: func(mockSql sqlmock.Sqlmock) {
				mockSql.ExpectQuery("SELECT id, name , email, phone, referal_code FROM users WHERE id= ?").WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "phone", "referal_code"}))
			},
			want:    nil,
			wantErr: errors.New("no data matching the specified criteria was found in the database"),
		},
	}

	for _, tt := range tests {
		mockDB, mockSQL, _ := sqlmock.New()
		defer mockDB.Close()

		DB, _ := gorm.Open(postgres.New(postgres.Config{
			Conn: mockDB,
		}), &gorm.Config{})

		tt.stub(mockSQL)

		userRepository := NewUserRepository(DB)
		result, err := userRepository.GetProfile(tt.args)

		assert.Equal(t, tt.want, result)
		assert.Equal(t, tt.wantErr, err)
	}
}

func TestCreateUser(t *testing.T) {
	test := []struct {
		name    string
		args    *requestmodel.UserDetails
		stub    func(sqlmock.Sqlmock)
		want    *responsemodel.SignupData
		wantErr error
	}{
		{
			name: "successfully user signup",
			args: &requestmodel.UserDetails{
				Name:        "Ashik",
				Email:       "ashik55@gmail.com",
				Phone:       "9876543210",
				ReferalCode: "87hj9",
				Password:    "8ugfe4567ujki876tfde45tyhjiu7ytrd",
			},
			//"INSERT INTO users \(name, email, phone, password, referal_code\) VALUES\(\$1, \$2, \$3, \$4, \$5\) RETURNING \*"  regex quey using https://www.regex-escape.com/online-regex-escaper.php
			// https://stackoverflow.com/questions/59652031/sqlmock-is-not-matching-query-but-query-is-identical-and-log-output-shows-the-s           .also refer hear
			stub: func(mockSql sqlmock.Sqlmock) {
				mockSql.ExpectQuery(regexp.QuoteMeta("INSERT INTO users (name, email, phone, password, referal_code) VALUES($1, $2, $3, $4, $5) RETURNING *")).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "phone", "email", "referal_code"}).AddRow(1, "Ashik", "9876543210", "ashik55@gmail.com", "87hj9"))
			},
			want: &responsemodel.SignupData{
				ID:          "1",
				Name:        "Ashik",
				Email:       "ashik55@gmail.com",
				Phone:       "9876543210",
				ReferalCode: "87hj9",
			},
			wantErr: nil,
		},
	}

	for _, tt := range test {
		mockDB, mocksql, _ := sqlmock.New()
		defer mockDB.Close()

		DB, _ := gorm.Open(postgres.New(postgres.Config{
			Conn: mockDB,
		}), &gorm.Config{})

		tt.stub(mocksql)

		userRepository := NewUserRepository(DB)
		result, err := userRepository.CreateUser(tt.args)

		assert.Equal(t, tt.want, result)
		assert.Equal(t, err, tt.wantErr)
	}
}
