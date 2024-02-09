package repository

import (
	"errors"
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

func TestIsUserExist(t *testing.T) {
	mockDB, mockSQL, _ := sqlmock.New()
	defer mockDB.Close()

	db, _ := gorm.Open(postgres.New(postgres.Config{Conn: mockDB}), &gorm.Config{})

	argument := "9876543210"
	want := 1

	expectedquery := "SELECT COUNT(*) FROM users WHERE phone=$1 AND status!=$2"
	mockSQL.ExpectQuery(regexp.QuoteMeta(expectedquery)).WillReturnRows(sqlmock.NewRows([]string{"phone"}).AddRow("1"))

	userRepository := NewUserRepository(db)
	result := userRepository.IsUserExist(argument)

	assert.Equal(t, want, result)
}

func TestFetchPasswordUsingPhone(t *testing.T) {
	test := []struct {
		name    string
		args    string
		stub    func(sqlmock.Sqlmock)
		want    string
		wantErr error
	}{
		{
			name: "succesfully fetch password",
			args: "9876543210",
			stub: func(mocksql sqlmock.Sqlmock) {
				mocksql.ExpectQuery(regexp.QuoteMeta("SELECT password FROM users WHERE phone=$1 AND status='active'")).
					WillReturnRows(sqlmock.NewRows([]string{"password"}).AddRow("kjhgfds45678ijhbvc"))
			},
			want:    "kjhgfds45678ijhbvc",
			wantErr: nil,
		}, {
			name: "no user is exist",
			args: "9876543210",
			stub: func(mocksql sqlmock.Sqlmock) {
				mocksql.ExpectQuery(regexp.QuoteMeta("SELECT password FROM users WHERE phone=? AND status='active'")).
					WillReturnError(errors.New("no user exist or you are blocked by admin"))
			},
			want:    "",
			wantErr: errors.New("no user exist or you are blocked by admin"),
		},
	}

	for _, tt := range test {
		mockdb, mocksql, _ := sqlmock.New()
		defer mockdb.Close()

		DB, _ := gorm.Open(postgres.New(postgres.Config{
			Conn: mockdb,
		}), &gorm.Config{})

		tt.stub(mocksql)

		userRepository := NewUserRepository(DB)
		result, _ := userRepository.FetchPasswordUsingPhone(tt.args)

		assert.Equal(t, tt.want, result)
		// assert.Equal(t, tt.wantErr, err)
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
		}, {
			name: "Error at user creation",
			args: &requestmodel.UserDetails{
				Name:        "Ashik",
				Email:       "ashik55@gmail.com",
				Phone:       "9876543210",
				ReferalCode: "87hj9",
				Password:    "8ugfe4567ujki876tfde45tyhjiu7ytrd",
			},
			stub: func(mocksql sqlmock.Sqlmock) {
				mocksql.ExpectQuery(regexp.QuoteMeta("INSERT INTO users (name, email, phone, password, referal_code) VALUES($1, $2, $3, $4, $5) RETURNING *")).
					WillReturnError(errors.New("data missmatching can't store in database"))
			},
			want:    nil,
			wantErr: errors.New("data missmatching can't store in database"),
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
		// log.Println("succfully test complete", tt.name)
	}
}

func TestGetAllUser(t *testing.T) {
	test := []struct {
		name    string
		args1   int
		args2   int
		stub    func(sqlmock sqlmock.Sqlmock)
		want    *[]responsemodel.UserDetails
		wantErr error
	}{
		{
			name:  "Get All user datas",
			args1: 1,
			args2: 2,
			stub: func(sqlmock sqlmock.Sqlmock) {
				query := "SELECT * FROM users ORDER BY name OFFSET $1 LIMIT $2"
				sqlmock.ExpectQuery(regexp.QuoteMeta(query)).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "phone", "status"}).
						AddRow("1", "user1", "user1@example.com", "1234567890", "block").
						AddRow("2", "user2", "user2@example.com", "0987654321", "ACTIVE"))

			},
			want: &[]responsemodel.UserDetails{{
				ID:     "1",
				Name:   "user1",
				Email:  "user1@example.com",
				Phone:  "1234567890",
				Status: "block",
			}, {
				ID:     "2",
				Name:   "user2",
				Email:  "user2@example.com",
				Phone:  "0987654321",
				Status: "ACTIVE",
			}},
			wantErr: nil,
		}, {
			name:  "no user exist",
			args1: 1,
			args2: 2,
			stub: func(sqlmock sqlmock.Sqlmock) {
				sqlmock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM users ORDER BY name OFFSET $1 LIMIT $2")).
					WillReturnError(errors.New("can't get user data from db"))
			},
			want:    nil,
			wantErr: errors.New("can't get user data from db"),
		},
	}

	for _, tt := range test {
		mockDB, sqlmock, _ := sqlmock.New()
		defer mockDB.Close()
		DB, _ := gorm.Open(postgres.New(postgres.Config{
			Conn: mockDB,
		}), &gorm.Config{})

		userRepository := NewUserRepository(DB)

		tt.stub(sqlmock)

		result, err := userRepository.AllUsers(tt.args1, tt.args2)

		assert.Equal(t, tt.want, result)
		assert.Equal(t, tt.wantErr, err)
	}
}

func TestCreateAddress(t *testing.T) {
	test := []struct {
		name    string
		args    *requestmodel.Address
		stub    func(sqlmock.Sqlmock)
		want    *requestmodel.Address
		wantErr error
	}{
		{
			name: "Create user address",
			args: &requestmodel.Address{
				Userid:      "123",
				FirstName:   "John",
				LastName:    "Doe",
				Street:      "kochi pallitheriv",
				City:        "bolgatti",
				State:       "kerala",
				Pincode:     "567843",
				LandMark:    "Near Park",
				PhoneNumber: "9876543210",
			},
			stub: func(sqlmock sqlmock.Sqlmock) {
				sqlmock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO addresses ( userid, first_name, last_name, street, city, state, pincode, land_mark, phone_number) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *;`)).
					WillReturnRows(sqlmock.NewRows([]string{"id", "userid", "first_name", "last_name", "street", "city", "state", "pincode", "land_mark", "phone_numbe"}).AddRow("1", "123", "John", "Doe", "kochi pallitheriv", "bolgatti", "kerala", "567843", "Near Park", "9876543210"))
			},
			want: &requestmodel.Address{
				ID:          "1",
				Userid:      "123",
				FirstName:   "John",
				LastName:    "Doe",
				Street:      "kochi pallitheriv",
				City:        "bolgatti",
				State:       "kerala",
				Pincode:     "567843",
				LandMark:    "Near Park",
				PhoneNumber: "9876543210",
			},
			wantErr: nil,
		},{
			name: "error at addres creation of user",
			args:  &requestmodel.Address{
				Userid:      "123",
				FirstName:   "John",
				LastName:    "Doe",
				Street:      "kochi pallitheriv",
				City:        "bolgatti",
				State:       "kerala",
				Pincode:     "567843",
				LandMark:    "Near Park",
				PhoneNumber: "9876543210",
			},
			stub: func(sqlmock sqlmock.Sqlmock) {
				sqlmock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO addresses ( userid, first_name, last_name, street, city, state, pincode, land_mark, phone_number) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *;`)).
					WillReturnRows(sqlmock.NewRows([]string{"id", "userid", "first_name", "last_name", "street", "city", "state", "pincode", "land_mark", "phone_numbe"}))
			},
			want: nil,
			wantErr: errors.New("no data matching the specified criteria was found in the database"),
		},
	}

	for _, tt := range test {
		mockDB, mocksql, _ := sqlmock.New()
		defer mockDB.Close()

		DB, _ := gorm.Open(postgres.New(postgres.Config{
			Conn: mockDB,
		}), &gorm.Config{})

		userRepository := NewUserRepository(DB)

		tt.stub(mocksql)

		result, err := userRepository.CreateAddress(tt.args)

		assert.Equal(t, tt.want, result)
		assert.Equal(t, tt.wantErr, err)
	}
}
