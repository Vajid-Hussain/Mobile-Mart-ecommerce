package responsemodel

type CartInventory struct {
	Productname string `form:"productname" validate:"required,min=3,max=100"`
	InventoryID string `json:"inventoryid" validate:"required,number"`
	SellerID    string `json:"sellerID" validate:"required"`
	Quantity    uint   `json:"quantity"`
	Saleprice   uint   `form:"saleprice" validate:"required,min=0,number"`
	Price       uint   `json:"total-amout"`
	Units       uint64 `json:"units" validate:"required,min=0,number"`
	ImageURL    string `json:"imageURL"`
}

type UserCart struct {
	UserID         string          `json:"user_id" validate:"" gorm:"-"`
	TotalPrice     uint            `json:"total_price"`
	InventoryCount uint            `json:"inventory_count"`
	Cart           []CartInventory `json:"cart" gorm:"-"`
}
