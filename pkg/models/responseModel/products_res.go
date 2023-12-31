package responsemodel

type InventoryRes struct {
	ID                 uint    `json:"id"`
	Productname        string  `json:"productname" validate:"required,min=3,max=100"`
	Description        string  `json:"description" validate:"required,min=5"`
	BrandID            uint    `json:"brandID" validate:"required"`
	CategoryID         uint    `json:"categoryID" validate:"required"`
	SellerID           string  `json:"sellerID" validate:"required"`
	Mrp                uint    `json:"mrp" validate:"required,min=0"`
	Discount           uint    `form:"discount" validate:"required,min=0,max=99,number"`
	Saleprice          uint    `json:"saleprice" validate:"required,min=0"`
	CategoryDiscount   uint    `json:"categoryDiscount,omitempty"`
	NetDiscount        uint    `json:"netDiscount,omitempty"`
	FinalPrice         uint    `json:"priceApplyCategoryDiscount,omitempty"`
	Units              uint64  `json:"units" validate:"required,min=0"`
	Os                 string  `json:"os"`
	CellularTechnology string  `json:"cellularTechnology"`
	Ram                uint    `json:"ram" validate:"required,min=0"`
	Screensize         float64 `json:"screensize" validate:"required,min=0"`
	Batterycapacity    uint    `json:"batterycapacity" validate:"required,min=0"`
	Processor          string  `json:"processor" validate:"required"`
	ImageURL           string  `json:"imageURL" validate:"required"`
}

type Errors struct {
	Err string
}

type InventoryShowcase struct {
	ID                              uint   `json:"productID"`
	Productname                     string `json:"productname"`
	Mrp                             int    `json:"mrp" `
	Discount                        uint   `form:"discount" `
	Saleprice                       int    `json:"saleprice" `
	CategoryDiscount                uint   `json:"categoryDiscount,omitempty"`
	NetDiscount                     uint   `json:"netDiscount,omitempty"`
	PriceAfterApplyCategoryDiscount uint   `json:"priceApplyCategoryDiscount,omitempty"`
	SellerID                        string `json:"sellerID" `
	Units                           uint   `json:"units"`
}

type FilterProduct struct {
	ID                              uint   `json:"product" gorm:"column:productid"`
	Productname                     string `json:"productname"`
	Mrp                             int    `json:"mrp" `
	Discount                        uint   `form:"discount" validate:"required,min=0,max=99,number"`
	Saleprice                       int    `json:"saleprice" `
	CategoryDiscount                uint   `json:"categoryDiscount,omitempty"`
	NetDiscount                     uint   `json:"netDiscount,omitempty"`
	PriceAfterApplyCategoryDiscount uint   `json:"priceApplyCategoryDiscount,omitempty"`
	Title                           string `json:"categoryOfferTitle"`
	SellerID                        string `json:"sellerID" `
	Category                        string `json:"categoty" gorm:"column:name"`
	Brand                           string `json:"brand" gorm:"column:name"`
	Units                           uint   `json:"units"`
}
