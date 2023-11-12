package interfaces

import responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"

type IAdminRepository interface {
	GetPassword(string) (string, error)
	AllUsers(int, int) (*[]responsemodel.UserDetails, error)
	UserCount(chan int)
	BlockUser(string) error
	UnblockUser(string) error
	AllSellers(int, int) (*[]responsemodel.SellerDetails, error)
	SellerCount(chan int)
	BlockSeller(string) error
	UnblockSeller(string) error
	GetPendingSellers(int, int) (*[]responsemodel.SellerDetails, error)
	GetSingleSeller(string) (*responsemodel.SellerDetails, error)
}
