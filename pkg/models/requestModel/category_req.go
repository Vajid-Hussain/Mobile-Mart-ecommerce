package requestmodel

type Category struct {
	Name string `json:"name" validate:"required"`
}

type CategoryDetails struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type BrandDetails struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type Brand struct {
	Name string `json:"name" validate:"required"`
}
