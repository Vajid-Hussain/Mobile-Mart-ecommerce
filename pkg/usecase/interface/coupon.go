package interfaceUseCase

import (
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
)

type ICouponUseCase interface {
	CreateCoupon(*requestmodel.Coupon) (*responsemodel.Coupon, error)
	GetCoupons() (*[]responsemodel.Coupon, error)
}
