package requestmodel

type Category struct {
	Name             string `json:"name" validate:"required,alpha"`
	CategoryDiscount uint   `json:"discount" validate:"min=1,max=100"`
}

type CategoryDetails struct {
	ID               string `json:"id" validate:"required"`
	Name             string `json:"name" validate:"required"`
	CategoryDiscount uint   `json:"discount" validate:"min=1,max=100"`
}

type BrandDetails struct {
	ID   string `json:"id" validate:"required,number"`
	Name string `json:"name" validate:"required"`
}

type Brand struct {
	Name string `json:"name" validate:"required,alpha"`
}
