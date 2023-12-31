package requestmodel

type Category struct {
	Name string `json:"name" validate:"required,alpha"`
}

type CategoryDetails struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type BrandDetails struct {
	ID   string `json:"id" validate:"required,number"`
	Name string `json:"name" validate:"required"`
}

type Brand struct {
	Name string `json:"name" validate:"required,alpha"`
}

type CategoryOffer struct {
	Title            string `json:"title" validate:"required"`
	CategoryID       string `json:"category_id" validate:"required"`
	SellerID         string `json:"seller_id"`
	CategoryDiscount uint   `json:"category_discount" validate:"required,min=1,max=99"`
	Validity         uint   `json:"validity" validate:"required,min=0"`
}

type EditCategoryOffer struct {
	ID               string `json:validate:"required"`
	SellerID         string
	Title            string `json:"title" validate:"required"`
	CategoryDiscount uint   `json:"category_discount" validate:"required,min=1,max=99"`
	Validity         uint   `json:"validity"`
}
