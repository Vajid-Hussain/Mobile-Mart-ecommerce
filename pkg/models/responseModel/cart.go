package responsemodel

type CartInventory struct {
	Productname string `form:"productname" validate:"required,min=3,max=100"`
	InventoryID string `json:"inventoryid" validate:"required,number"`
	Quantity    uint   `json:"quantity"`
	Saleprice   uint   `form:"saleprice" validate:"required,min=0,number"`
	Price       uint   `json:"total-amout"`
	Units       uint64 `form:"units" validate:"required,min=0,number"`
	ImageURL    string
}

type UserCart struct {
	UserID         string          `json:"user_id" validate:""`
	Price          uint            `json:"tatalprice"`
	InventoryCount uint            `json:"inventorycount"`
	Cart           []CartInventory `json:"cart"`
}
