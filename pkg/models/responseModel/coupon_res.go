package responsemodel

import "time"

type Coupon struct {
	ID              string    `json:"couponID"`
	Name            string    `json:"name"`
	Type            string    `json:"type"`
	Discount        uint      `json:"discount"`
	MinimumRequired uint      `json:"minimum_required"`
	MaximumAllowed  uint      `json:"maximum_allowed"`
	StartDate       time.Time `json:"createTime,omitempty"`
	EndDate         time.Time `json:"expire_date"`
	Status          string    `json:"status"`
}
