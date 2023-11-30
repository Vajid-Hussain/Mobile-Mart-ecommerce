package responsemodel

type OrderShowcase struct {
	Productname string `json:"productname" validate:"required,min=3,max=100"`
	ID          string `gorm:"column:id" json:"orderID" validate:"required,number"`
	UserID      string `gorm:"column:user_id" json:"userid"`
	SellerID    string `json:"seller id" gorm:"column:seller_id"`
	InventoryID string `gorm:"column:inventory_id" json:"productid"`
	Price       uint   `json:"total-amout"`
	OrderStatus string `json:"orderstatus,omitempty"`
	Quantity    uint   `json:"quantity"`
	ImageURL    string `json:"imageURL"`
}

type OrderDetails struct {
	ID            string `gorm:"id" json:"orderid"`
	UserID        string `gorm:"column:user_id" json:"userid"`
	Address       string `gorm:"column:address_id" json:"address_id"`
	Payment       string `gorm:"column:payment_method" json:"payment"`
	SellerID      string `json:"seller id" gorm:"column:seller_id"`
	InventoryID   string `gorm:"column:inventory_id" json:"productid"`
	Quantity      uint   `json:"quantity"`
	Price         uint   `json:"price"`
	OrderStatus   string `json:"orderstatus,omitempty"`
	PaymentStatus string `json:"paymentStatu,omitempty"`
}

type SingleOrder struct {
	Productname  string `json:"productname" validate:"required,min=3,max=100"`
	ID           string `gorm:"column:id" json:"orderID" validate:"required,number"`
	SingleUnit   uint   `json:"Price of a unit" gorm:"column:saleprice"`
	Price        uint   `json:"total-amout" `
	Quantity     uint   `json:"quantity"`
	OrderStatus  string `json:"order_status"`
	OrderDate    string `json:"orderdate"`
	DeliveryDate string `json:"delivary_date,omitempty"`
	ImageURL     string `json:"imageURL"`
	FirstName    string `json:"firstName" validate:"required"`
	LastName     string `json:"lastName,omitempty"`
	Street       string `json:"street" validate:"required,alpha"`
	City         string `json:"city" validate:"required,alpha"`
	State        string `json:"state" validate:"required,alpha"`
	Pincode      string `json:"pincode" validate:"min=6"`
	LandMark     string `json:"landmark" validate:"required"`
	PhoneNumber  string `json:"phoneNumber" validate:"required,len=10,number"`
}

type OrderSuccess struct {
	UserID     string `gorm:"column:user_id" json:"userid"`
	Address    string `gorm:"column:address_id" json:"address_id"`
	Payment    string `gorm:"column:payment_method" json:"payment"`
	TotalWorth uint   `json:"payable_amount"`
	Orders     []OrderDetails
}

type SalesReport struct {
	Orders   uint `json:"total -orders"`
	Quantity uint `json:"total-unit"`
	Price    uint `json:"total-price"`
}

type DashBord struct {
	TotalOrders        uint   `json:"totalOrders"`
	DeliveredOrders    uint   `json:"deliveredOrders"`
	OngoingOrders      uint   `json:"OngoingOrders"`
	CancelledOrders    uint   `json:"cancelledOrders"`
	TotalRevenue       uint   `json:"totalRevenue"`
	TotalSelledProduct uint   `json:"totalSelledProduct"`
	AdminCredit        uint   `json:"adminCredit"`
	LowStockProductID  []uint `json:"LowStockProductID"`
}

type OnlinePayment struct {
	OrderID     string `json:"orderID" `
	User        string `gorm:"column:first_name" json:"user"`
	FinalPrice  uint   `json:"finalPrice"`
	PhoneNumber uint   `json:"phoneNumber" `
}
