package usecase

import (
	"errors"
	"testing"

	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/config"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/mock/mockRepository"
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetAddress(t *testing.T) {
	ctrl := gomock.NewController(t)

	userRepo := mockRepository.NewMockIUserRepo(ctrl)
	// config, _ := config.LoadConfig()
	// fmt.Println("-------------", config)
	paymentRepo := mockRepository.NewMockIPaymentRepository(ctrl)
	userUseCase := NewUserUseCase(userRepo, paymentRepo, &config.Token{})

	testData := map[string]struct {
		userID    string
		addressID string
		stub      func(mockRepository.MockIPaymentRepository, mockRepository.MockIUserRepo, string, string)
		wantErr   error
	}{
		"success": {
			userID:    "1",
			addressID: "2",
			stub: func(mir1 mockRepository.MockIPaymentRepository, mir2 mockRepository.MockIUserRepo, s1, s2 string) {
				userRepo.EXPECT().DeleteAddress(s1, s2).Times(1).Return(nil)
			},
			wantErr: nil,
		},
		"failed": {
			userID:    "1",
			addressID: "2",
			stub: func(mir1 mockRepository.MockIPaymentRepository, mir2 mockRepository.MockIUserRepo, s1, s2 string) {
				userRepo.EXPECT().DeleteAddress(s1, s2).Times(1).Return(errors.New("no address exist"))
			},
			wantErr: errors.New("no address exist"),
		},
	}

	for _, tt := range testData {
		tt.stub(*paymentRepo, *userRepo, tt.userID, tt.addressID)
		err := userUseCase.DeleteAddress(tt.userID, tt.addressID)
		assert.Equal(t, err, tt.wantErr)
	}
}

func TestAddAddress(t *testing.T) {
	ctrl := gomock.NewController(t)

	userRepo := mockRepository.NewMockIUserRepo(ctrl)
	paymentRepo := mockRepository.NewMockIPaymentRepository(ctrl)
	userUseCase := NewUserUseCase(userRepo, paymentRepo, &config.Token{})

	testData := map[string]struct {
		arg     *requestmodel.Address
		stub    func(mockRepository.MockIPaymentRepository, mockRepository.MockIUserRepo, *requestmodel.Address)
		want    *requestmodel.Address
		wantErr error
	}{
		"success": {
			arg: &requestmodel.Address{
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
			stub: func(mir1 mockRepository.MockIPaymentRepository, mir2 mockRepository.MockIUserRepo, arg *requestmodel.Address) {
				userRepo.EXPECT().CreateAddress(arg).Times(1).Return(&requestmodel.Address{Userid: "123",
					FirstName:   "John",
					LastName:    "Doe",
					Street:      "kochi pallitheriv",
					City:        "bolgatti",
					State:       "kerala",
					Pincode:     "567843",
					LandMark:    "Near Park",
					PhoneNumber: "9876543210"}, nil)
			},
			want: &requestmodel.Address{
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
		},
		"fail":{
			arg: &requestmodel.Address{
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
			stub: func(mir1 mockRepository.MockIPaymentRepository, mir2 mockRepository.MockIUserRepo, a *requestmodel.Address) {
				userRepo.EXPECT().CreateAddress(a).Times(1).Return(nil,errors.New("missmath in addres data"))
			},
			want: nil,
			wantErr: errors.New("missmath in addres data"),
		},
	}

	for _, tt := range testData {
		tt.stub(*paymentRepo, *userRepo, tt.arg)
		result, err := userUseCase.AddAddress(tt.arg)
		assert.Equal(t, tt.want, result)
		assert.Equal(t, tt.wantErr, err)
	}
}
