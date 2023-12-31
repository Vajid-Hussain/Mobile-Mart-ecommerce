package requestmodel

type Coupon struct {
	Name            string `json:"name" validate:"required"`
	Type            string `json:"type" validate:"required,alpha"`
	Discount        uint   `json:"discount" validate:"min=1,max=100"`
	MinimumRequired uint   `json:"minimum_required" validate:"min=0"`
	MaximumAllowed  uint   `json:"maximum_allowed" validate:"gtcsfield=MinimumRequired"`
	ExpireDate      uint   `json:"expire_date" validate:"min=1"`
}
