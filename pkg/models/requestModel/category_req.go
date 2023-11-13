package requestmodel

type Category struct {
	Name string `json:"name" validate:"required"`
}
