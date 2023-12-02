package responsemodel

import "time"

type OrderShowcase struct {
	Productname string `json:"productname" validate:"required,min=3,max=100"`
	ID          string `gorm:"column:order_id" json:"orderID" validate:"required,number"`
	UserID      string `gorm:"column:user_id" json:"userid"`
	SellerID    string `json:"seller id" gorm:"column:seller_id"`
	InventoryID string `gorm:"column:inventory_id" json:"productid"`
	Price       uint   `json:"total-amout"`
	Saleprice   uint   `json:"singleUnitCost"`
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

type OrderProducts struct {
	OrderID     string `json:"orderID"`
	InventoryID string `json:"productID"`
	SellerID    string `json:"sellerID"`
	Quantity    uint   `json:"quantity"`
	Price       uint   `json:"price"`
	ImageURL    string `json:"imageURL"`
}

type Order struct {
	ID             string    `json:"id"`
	UserID         string    `gorm:"column:user_id" json:"userid"`
	Address        string    `gorm:"column:address_id" json:"address_id"`
	Payment        string    `gorm:"column:payment_method" json:"payment"`
	TotalPrice     uint      `json:"payable_amount"`
	OrderDate      time.Time `json:"orderDate"`
	DeliveryDate   string    `json:"delivaryDate,omitempty"`
	OrderStatus    string    `json:"omitempty"`
	PaymentStatus  string    `json:"paymentStatus,omitempty"`
	OrderIDRazopay string    `json:"razopayOrderID,omitempty"`
	Orders         []OrderProducts
}

type SingleOrder struct {
	Productname  string `json:"productname" validate:"required,min=3,max=100"`
	ID           string `gorm:"column:order_id" json:"orderID" validate:"required,number"`
	SingleUnit   uint   `json:"PriceOfAUnit" gorm:"column:saleprice"`
	Price        uint   `json:"totalAmout" `
	Quantity     uint   `json:"quantity"`
	OrderStatus  string `json:"orderStatus"`
	OrderDate    string `json:"orderdate"`
	DeliveryDate string `json:"delivaryDate,omitempty"`
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
