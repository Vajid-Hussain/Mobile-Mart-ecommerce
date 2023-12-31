package requestmodel

import (
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
)

type Order struct {
	ID             uint                          `json:"-"`
	UserID         string                        `json:"userid" swaggerignore:"true"`
	Address        string                        `json:"address" validate:"required,numeric"`
	Payment        string                        `json:"payment" validate:"required,alpha,uppercase"`
	Coupon         string                        `json:"coupon"`
	OrderIDRazopay string                        `json:"-"`
	FinalPrice     uint                          `json:"-"`
	CouponDiscount uint                          `json:"-"`
	OrderStatus    string                        `json:"-"`
	PaymentStatus  string                        `json:"-"`
	Cart           []responsemodel.CartInventory `json:"-"`
}

type OnlinePaymentVerification struct {
	PaymentID string `json:"payment_id" validate:"required"`
	OrderID   string `json:"order_id" validate:"required"`
	Signature string `json:"signature" validate:"required"`
}
