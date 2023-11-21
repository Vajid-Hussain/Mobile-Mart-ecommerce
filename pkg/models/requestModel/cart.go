package requestmodel

type Cart struct {
	UserID      string `json:"user_id" validate:""`
	InventoryID string `json:"inventoryid" validate:"required,number"`
	Quantity    uint   `json:"quantity"`
	Price       uint   `json:"price"`
}
