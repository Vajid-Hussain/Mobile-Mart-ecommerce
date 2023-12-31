package responsemodel

type CartInventory struct {
	Productname      string `json:"productname" validate:"required,min=3,max=100"`
	InventoryID      string `json:"productid" validate:"required,number"`
	SellerID         string `json:"sellerID" validate:"required"`
	CategoryID       string `json:"categoryID"`
	Quantity         uint   `json:"quantity"`
	Price            uint   `json:"mrp" gorm:"column:mrp"`
	Discount         uint   `json:"productDiscount"`
	Saleprice        uint   `json:"saleprice" validate:"required,min=0,number"`
	CategoryDiscount uint   `json:"categoryDiscount,omitempty"`
	FinalPrice       uint   `json:"payedAmount"`
	Title            string `json:"categoryDiscountTitle,omitempty"`
	Units            uint64 `json:"available units" validate:"required,min=0,number"`
}

type UserCart struct {
	UserID         string          `json:"user_id" validate:"" gorm:"-"`
	TotalPrice     uint            `json:"total_price"`
	InventoryCount uint            `json:"product_count"`
	Cart           []CartInventory `json:"cart" gorm:"-"`
}
