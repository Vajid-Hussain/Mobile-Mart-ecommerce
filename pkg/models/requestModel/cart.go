package requestmodel

type Cart struct {
	UserID      string `json:"cartid userid" validate:""`
	InventoryID string `json:"inventoryid" validate:"required,number"`
	Quantity    uint   `json:"quantity"`
	Price       uint   `json:"price,omitempty"`
}
