package interfaceUseCase

import (
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
)

type ISellerUseCase interface {
	SellerSignup(*requestmodel.SellerSignup) (*responsemodel.SellerSignupRes, error)
	SellerLogin(*requestmodel.SellerLogin) (*responsemodel.SellerLoginRes, error)
}
