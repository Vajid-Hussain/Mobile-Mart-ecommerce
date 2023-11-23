package requestmodel

import responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"

type Order struct {
	UserID  string                        `json:"userid" validate:"required"`
	Address string                        `json:"address" validate:"required"`
	Payment string                        `json:"payment" validate:"required"`
	Cart    []responsemodel.CartInventory `json:"-"`
}
