package responsemodel

type Category struct {
	Name string `json:"name,omitempty"`
}

type CategoryDetails struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type BrandRes struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
