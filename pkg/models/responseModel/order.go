package responsemodel

type OrderShowcase struct {
	Productname string `json:"productname" validate:"required,min=3,max=100"`
	ID          string `gorm:"column:id" json:"orderID" validate:"required,number"`
	Price       uint   `json:"total-amout"`
	Quantity    uint   `json:"quantity"`
	ImageURL    string `json:"imageURL"`
}

type OrderSuccess struct {
	UserID      string `gorm:"column:user_id" json:"userid"`
	Address     string `gorm:"column:address_id" json:"address_id"`
	Payment     string `gorm:"column:payment_method" json:"payment"`
	InventoryID string `gorm:"column:inventory_id" json:"inventoryid"`
	Quantity    uint   `json:"quantity"`
	Price       uint   `json:"total-amout"`
}

type SingleOrder struct {
	Productname string `json:"productname" validate:"required,min=3,max=100"`
	ID          string `gorm:"column:id" json:"orderID" validate:"required,number"`
	Price       uint   `json:"total-amout"`
	Quantity    uint   `json:"quantity"`
	OrderStatus string `json:"order_status"`
	ImageURL    string `json:"imageURL"`
	FirstName   string `json:"firstName" validate:"required"`
	LastName    string `json:"lastName"`
	Street      string `json:"street" validate:"required,alpha"`
	City        string `json:"city" validate:"required,alpha"`
	State       string `json:"state" validate:"required,alpha"`
	Pincode     string `json:"pincode" validate:"min=6"`
	LandMark    string `json:"landmark" validate:"required"`
	PhoneNumber string `json:"phoneNumber" validate:"required,len=10,number"`
}
