package responsemodel

type CartInventory struct {
	Productname string `json:"productname" validate:"required,min=3,max=100"`
	InventoryID string `json:"productid" validate:"required,number"`
	SellerID    string `json:"sellerID" validate:"required"`
	Quantity    uint   `json:"quantity"`
	Discount    uint   `json:"discountPersentage"`
	Saleprice   uint   `json:"saleprice" validate:"required,min=0,number"`
	Price       uint   `json:"mrp" gorm:"column:mrp"`
	Units       uint64 `json:"available units" validate:"required,min=0,number"`
	ImageURL    string `json:"imageURL"`
}

type UserCart struct {
	UserID         string          `json:"user_id" validate:"" gorm:"-"`
	TotalPrice     uint            `json:"total_price"`
	InventoryCount uint            `json:"product_count"`
	Cart           []CartInventory `json:"cart" gorm:"-"`
}
