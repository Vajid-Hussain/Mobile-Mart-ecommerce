package usecase

import (
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
	interfaces "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/repository/interface"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
)

type couponUseCase struct {
	repo interfaces.ICouponRepository
}

func NewCouponUseCase(repository interfaces.ICouponRepository) interfaceUseCase.ICouponUseCase {
	return &couponUseCase{repo: repository}
}

func (r *couponUseCase) CreateCoupon(newCoupon *requestmodel.Coupon) (*responsemodel.Coupon, error) {
	coupon, err := r.repo.CreateCoupon(newCoupon)
	if err != nil {
		return nil, err
	}
	return coupon, nil
}

func (r *couponUseCase) GetCoupons() (*[]responsemodel.Coupon, error) {
	coupons, err := r.repo.GetCoupons()
	if err != nil {
		return nil, err
	}
	return coupons, nil
}

// func (r *couponUseCase) UpdateCouponStatus(couponID, active, block string) (*responsemodel.Coupon, error) {

// }
