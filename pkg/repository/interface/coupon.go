package interfaces

import (
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
)

type ICouponRepository interface {
	CreateCoupon(*requestmodel.Coupon) (*responsemodel.Coupon, error)
	CheckCouponExpired(string) (*responsemodel.Coupon, error)
	GetCoupons() (*[]responsemodel.Coupon, error)
}
