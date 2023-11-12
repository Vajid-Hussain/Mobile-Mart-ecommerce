package interfaceUseCase

import (
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
)

type IAdminUseCase interface {
	AdminLogin(*requestmodel.AdminLoginData) (*responsemodel.AdminLoginRes, error)
	GetAllUsers(string, string) (*[]responsemodel.UserDetails, *int, error)
	BlcokUser(string) error
	UnblockUser(string) error
	GetAllSellers(string, string) (*[]responsemodel.SellerDetails, *int, error)
	BlockSeller(string) error
	UnblockSeller(string) error
	GetAllPendingSellers(string, string) (*[]responsemodel.SellerDetails, error)
	FetchSingleVender(string) (*responsemodel.SellerDetails, error)
}
