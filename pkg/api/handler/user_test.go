package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	mockusecase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/mock/mockUseCase"
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserSignup(t *testing.T) {
	testCase := map[string]struct {
		input         requestmodel.UserDetails
		buildstub     func(useCaseMock *mockusecase.MockIuserUseCase, signupData requestmodel.UserDetails)
		checkResponse func(t *testing.T, responserecorder *httptest.ResponseRecorder)
	}{
		"success": {
			input: requestmodel.UserDetails{
				Name:            "vajid",
				Email:           "vajid44@gmail.com",
				Phone:           "9876543210",
				Password:        "jhy78ij",
				ConfirmPassword: "jhy78ij",
			},
			buildstub: func(useCaseMock *mockusecase.MockIuserUseCase, signupData requestmodel.UserDetails) {
				err := validator.New().Struct(signupData)
				if err != nil {
					fmt.Println("validation failed")
				}
				useCaseMock.EXPECT().UserSignup(&signupData).Times(1).Return(&responsemodel.SignupData{
					Name:          "vajid",
					Email:         "vajid44@gmail.com",
					Phone:         "9876543210",
					ID:            "3",
					ReferalCode:   "234rd34",
					WalletBelance: 100,
				}, nil)
			},
			checkResponse: func(t *testing.T, responserecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, responserecorder.Code)
			},
		},
		"bad requst": {
			input: requestmodel.UserDetails{
				Name:     "vajid",
				Email:    "vajid44@gmail.com",
				Phone:    "9876543210",
				Password: "jhy78ij",
				ConfirmPassword: "jhy78ij",
			},
			buildstub: func(useCaseMock *mockusecase.MockIuserUseCase, signupData requestmodel.UserDetails) {
				err := validator.New().Struct(signupData)
				if err != nil {
					fmt.Println("validation error")
				}
				useCaseMock.EXPECT().UserSignup(&signupData).Times(1).Return(nil, errors.New("user request data is not satisfying credential"))
			},
			checkResponse: func(t *testing.T, responserecorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, responserecorder.Code)
			},
		},
	}

	for testname, test := range testCase {
		test := test
		t.Run(testname, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			mockUseCase := mockusecase.NewMockIuserUseCase(ctrl)
			test.buildstub(mockUseCase, test.input)
			userHandler := NewUserHandler(mockUseCase)

			server := gin.Default()
			server.POST("/signup", userHandler.UserSignup)

			jsonData, err := json.Marshal(test.input)
			assert.NoError(t, err)
			body := bytes.NewBuffer(jsonData)

			mockRequst, err := http.NewRequest(http.MethodPost, "/signup", body)
			assert.NoError(t, err)
			responseRecord := httptest.NewRecorder()
			server.ServeHTTP(responseRecord, mockRequst)

			test.checkResponse(t, responseRecord)
		})
	}
}
