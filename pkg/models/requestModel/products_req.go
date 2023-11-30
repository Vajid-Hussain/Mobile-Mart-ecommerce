package requestmodel

import "mime/multipart"

type InventoryReq struct {
	Productname        string                `form:"productname" validate:"required,min=3,max=100"`
	Description        string                `form:"description" validate:"required,min=5"`
	BrandID            uint                  `form:"brandID" validate:"required,number"`
	CategoryID         uint                  `form:"categoryID" validate:"required,number"`
	SellerID           uint                  `form:"sellerID" validate:"required,number"`
	Mrp                uint                  `form:"mrp" validate:"required,min=0,number"`
	Saleprice          uint                  `form:"saleprice" validate:"required,min=0,number"`
	Units              uint64                `form:"units" validate:"required,min=0,number"`
	Os                 string                `form:"os" validate:"required"`
	CellularTechnology string                `form:"cellularTechnology" validate:"required"`
	Ram                uint                  `form:"ram" validate:"required,min=1"`
	Screensize         float64               `form:"screensize" validate:"required,min=2"`
	Batterycapacity    uint                  `form:"batterycapacity" validate:"required,min=500"`
	Processor          string                `form:"processor" validate:"required"`
	Image              *multipart.FileHeader `form:"image"`
	ImageURL           string
}

type EditInventory struct {
	ID                 uint    `json:"id" validate:"required"`
	Productname        string  `json:"productname" validate:"required,min=3,max=100"`
	Description        string  `json:"description" validate:"required,min=5"`
	BrandID            uint    `json:"brandID" validate:"required"`
	CategoryID         uint    `json:"categoryID" validate:"required"`
	SellerID           string  `json:"cellerID" validate:"required"`
	Mrp                uint    `json:"mrp" validate:"required,min=0"`
	Saleprice          uint    `json:"saleprice" validate:"required,min=0"`
	Units              uint64  `json:"units" validate:"required,min=0"`
	Os                 string  `json:"os" validate:"required"`
	CellularTechnology string  `json:"cellularTechnology" validate:"required"`
	Ram                uint    `json:"ram" validate:"required,min=0"`
	Screensize         float64 `json:"screensize" validate:"required,min=0"`
	Batterycapacity    uint    `json:"batterycapacity" validate:"required,min=500"`
	Processor          string  `json:"processor" validate:"required"`
}
type FilterCriterion struct {
	Category string `json:"category" validate:"alpha"`
	Brand    string `json:"brand" validate:"alpha"`
	Product  string `json:"product" validate:"alpha"`
	MinPrice uint   `json:"minprice" validate:"numeric"`
	MaxPrice uint   `json:"maxprice" validate:"numeric"`
}
