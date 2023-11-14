package requestmodel

type Category struct {
	Name string `json:"name" validate:"required"`
}

type CategoryDetails struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}
