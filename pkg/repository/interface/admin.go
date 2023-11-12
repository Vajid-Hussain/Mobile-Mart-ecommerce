package interfaces

import responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"

type IAdminRepository interface {
	GetPassword(string) (string, error)
	AllUsers(int, int) (*[]responsemodel.UserDetails, error)
	UserCount(chan int)
	BlockUser(string) error
}
