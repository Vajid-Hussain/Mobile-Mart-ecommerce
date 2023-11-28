package interfaceUseCase

import (
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
)

type IAdminUseCase interface {
	AdminLogin(*requestmodel.AdminLoginData) (*responsemodel.AdminLoginRes, error)
	GetSellerDetailsForDashBord() (*responsemodel.AdminDashBord, error)
}
