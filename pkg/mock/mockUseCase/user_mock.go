// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/usecase/interface/user.go

// Package mockusecase is a generated GoMock package.
package mockusecase

import (
	reflect "reflect"

	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
	gomock "github.com/golang/mock/gomock"
)

// MockIuserUseCase is a mock of IuserUseCase interface.
type MockIuserUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockIuserUseCaseMockRecorder
}

// MockIuserUseCaseMockRecorder is the mock recorder for MockIuserUseCase.
type MockIuserUseCaseMockRecorder struct {
	mock *MockIuserUseCase
}

// NewMockIuserUseCase creates a new mock instance.
func NewMockIuserUseCase(ctrl *gomock.Controller) *MockIuserUseCase {
	mock := &MockIuserUseCase{ctrl: ctrl}
	mock.recorder = &MockIuserUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIuserUseCase) EXPECT() *MockIuserUseCaseMockRecorder {
	return m.recorder
}

// AddAddress mocks base method.
func (m *MockIuserUseCase) AddAddress(arg0 *requestmodel.Address) (*requestmodel.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAddress", arg0)
	ret0, _ := ret[0].(*requestmodel.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAddress indicates an expected call of AddAddress.
func (mr *MockIuserUseCaseMockRecorder) AddAddress(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAddress", reflect.TypeOf((*MockIuserUseCase)(nil).AddAddress), arg0)
}

// BlcokUser mocks base method.
func (m *MockIuserUseCase) BlcokUser(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlcokUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// BlcokUser indicates an expected call of BlcokUser.
func (mr *MockIuserUseCaseMockRecorder) BlcokUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlcokUser", reflect.TypeOf((*MockIuserUseCase)(nil).BlcokUser), arg0)
}

// DeleteAddress mocks base method.
func (m *MockIuserUseCase) DeleteAddress(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAddress", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAddress indicates an expected call of DeleteAddress.
func (mr *MockIuserUseCaseMockRecorder) DeleteAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAddress", reflect.TypeOf((*MockIuserUseCase)(nil).DeleteAddress), arg0, arg1)
}

// EditAddress mocks base method.
func (m *MockIuserUseCase) EditAddress(arg0 *requestmodel.EditAddress) (*requestmodel.EditAddress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditAddress", arg0)
	ret0, _ := ret[0].(*requestmodel.EditAddress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditAddress indicates an expected call of EditAddress.
func (mr *MockIuserUseCaseMockRecorder) EditAddress(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditAddress", reflect.TypeOf((*MockIuserUseCase)(nil).EditAddress), arg0)
}

// ForgotPassword mocks base method.
func (m *MockIuserUseCase) ForgotPassword(arg0 *requestmodel.ForgotPassword, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForgotPassword", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForgotPassword indicates an expected call of ForgotPassword.
func (mr *MockIuserUseCaseMockRecorder) ForgotPassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForgotPassword", reflect.TypeOf((*MockIuserUseCase)(nil).ForgotPassword), arg0, arg1)
}

// GetAddress mocks base method.
func (m *MockIuserUseCase) GetAddress(arg0, arg1, arg2 string) (*[]requestmodel.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddress", arg0, arg1, arg2)
	ret0, _ := ret[0].(*[]requestmodel.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAddress indicates an expected call of GetAddress.
func (mr *MockIuserUseCaseMockRecorder) GetAddress(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddress", reflect.TypeOf((*MockIuserUseCase)(nil).GetAddress), arg0, arg1, arg2)
}

// GetAllUsers mocks base method.
func (m *MockIuserUseCase) GetAllUsers(arg0, arg1 string) (*[]responsemodel.UserDetails, *int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers", arg0, arg1)
	ret0, _ := ret[0].(*[]responsemodel.UserDetails)
	ret1, _ := ret[1].(*int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockIuserUseCaseMockRecorder) GetAllUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockIuserUseCase)(nil).GetAllUsers), arg0, arg1)
}

// GetProfile mocks base method.
func (m *MockIuserUseCase) GetProfile(arg0 string) (*requestmodel.UserDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfile", arg0)
	ret0, _ := ret[0].(*requestmodel.UserDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfile indicates an expected call of GetProfile.
func (mr *MockIuserUseCaseMockRecorder) GetProfile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfile", reflect.TypeOf((*MockIuserUseCase)(nil).GetProfile), arg0)
}

// SendOtp mocks base method.
func (m *MockIuserUseCase) SendOtp(arg0 *requestmodel.SendOtp) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendOtp", arg0)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendOtp indicates an expected call of SendOtp.
func (mr *MockIuserUseCaseMockRecorder) SendOtp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendOtp", reflect.TypeOf((*MockIuserUseCase)(nil).SendOtp), arg0)
}

// UnblockUser mocks base method.
func (m *MockIuserUseCase) UnblockUser(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnblockUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnblockUser indicates an expected call of UnblockUser.
func (mr *MockIuserUseCaseMockRecorder) UnblockUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnblockUser", reflect.TypeOf((*MockIuserUseCase)(nil).UnblockUser), arg0)
}

// UpdateProfile mocks base method.
func (m *MockIuserUseCase) UpdateProfile(arg0 *requestmodel.UserEditProfile) (*requestmodel.UserDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProfile", arg0)
	ret0, _ := ret[0].(*requestmodel.UserDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProfile indicates an expected call of UpdateProfile.
func (mr *MockIuserUseCaseMockRecorder) UpdateProfile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfile", reflect.TypeOf((*MockIuserUseCase)(nil).UpdateProfile), arg0)
}

// UserLogin mocks base method.
func (m *MockIuserUseCase) UserLogin(arg0 requestmodel.UserLogin) (responsemodel.UserLogin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserLogin", arg0)
	ret0, _ := ret[0].(responsemodel.UserLogin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserLogin indicates an expected call of UserLogin.
func (mr *MockIuserUseCaseMockRecorder) UserLogin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserLogin", reflect.TypeOf((*MockIuserUseCase)(nil).UserLogin), arg0)
}

// UserSignup mocks base method.
func (m *MockIuserUseCase) UserSignup(arg0 *requestmodel.UserDetails) (*responsemodel.SignupData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserSignup", arg0)
	ret0, _ := ret[0].(*responsemodel.SignupData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserSignup indicates an expected call of UserSignup.
func (mr *MockIuserUseCaseMockRecorder) UserSignup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserSignup", reflect.TypeOf((*MockIuserUseCase)(nil).UserSignup), arg0)
}

// VerifyOtp mocks base method.
func (m *MockIuserUseCase) VerifyOtp(arg0 requestmodel.OtpVerification, arg1 string) (responsemodel.OtpValidation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyOtp", arg0, arg1)
	ret0, _ := ret[0].(responsemodel.OtpValidation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyOtp indicates an expected call of VerifyOtp.
func (mr *MockIuserUseCaseMockRecorder) VerifyOtp(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyOtp", reflect.TypeOf((*MockIuserUseCase)(nil).VerifyOtp), arg0, arg1)
}
