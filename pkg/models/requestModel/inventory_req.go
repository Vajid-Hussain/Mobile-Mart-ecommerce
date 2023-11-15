package requestmodel

type InventoryReq struct {
	Productname        string  `json:"productname" validate:"required,min=3,max=100"`
	Description        string  `json:"description" validate:"required,min=5"`
	BrandID            uint    `json:"brandID" validate:"required"`
	CategoryID         uint    `json:"categoryID" validate:"required"`
	SellerID           string  `json:"cellerID" validate:"required"`
	Mrp                int     `json:"mrp" validate:"required,min=0"`
	Saleprice          int     `json:"saleprice" validate:"required,min=0"`
	Units              int64   `json:"units" validate:"required,min=0"`
	Os                 string  `json:"os"`
	CellularTechnology string  `json:"cellularTechnology"`
	Ram                int     `json:"ram" validate:"required,min=0"`
	Screensize         float64 `json:"screensize" validate:"required,min=0"`
	Batterycapacity    int     `json:"batterycapacity" validate:"required,min=0"`
	Processor          string  `json:"processor" validate:"required"`
}
