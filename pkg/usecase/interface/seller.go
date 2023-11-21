package interfaceUseCase

import (
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
)

type ISellerUseCase interface {
	SellerSignup(*requestmodel.SellerSignup) (*responsemodel.SellerSignupRes, error)
	SellerLogin(*requestmodel.SellerLogin) (*responsemodel.SellerLoginRes, error)
	GetAllSellers(string, string) (*[]responsemodel.SellerDetails, *int, error)
	BlockSeller(string) error
	ActiveSeller(string) error
	GetAllPendingSellers(string, string) (*[]responsemodel.SellerDetails, error)
	FetchSingleVender(string) (*responsemodel.SellerDetails, error)

	GetSellerProfile(string) (*responsemodel.SellerProfile, error)
	UpdateSellerProfile(*requestmodel.SellerEditProfile) (*responsemodel.SellerProfile, error)
}
