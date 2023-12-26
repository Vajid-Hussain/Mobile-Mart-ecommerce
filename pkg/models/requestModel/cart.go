package requestmodel

type Cart struct {
	UserID      string `json:"cartid userid" swaggerignore:"true"`
	InventoryID string `json:"inventoryid" validate:"required,number"`
	Quantity    uint   `json:"quantity" swaggerignore:"true"`
	Price       uint   `json:"price,omitempty" swaggerignore:"true"`
}
