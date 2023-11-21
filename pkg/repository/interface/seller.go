package interfaces

import (
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
)

type ISellerRepo interface {
	IsSellerExist(string) (int, error)
	CreateSeller(*requestmodel.SellerSignup) error
	GetHashPassAndStatus(string) (string, string, string, error)
	GetPasswordByMail(string) string
	AllSellers(int, int) (*[]responsemodel.SellerDetails, error)
	SellerCount(chan int)
	BlockSeller(string) error
	UnblockSeller(string) error
	GetPendingSellers(int, int) (*[]responsemodel.SellerDetails, error)
	GetSingleSeller(string) (*responsemodel.SellerDetails, error)
	BlockInventoryOfSeller(string) error
	ActiveInventoryOfSeller(string) error

	GetSellerProfile(string) (*responsemodel.SellerProfile, error)
	UpdateSellerProfile(*requestmodel.SellerEditProfile) (*responsemodel.SellerProfile, error)
}
