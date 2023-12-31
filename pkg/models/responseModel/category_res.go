package responsemodel

import "time"

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

type CategoryOffer struct {
	ID               string    `json:"id"`
	Title            string    `json:"title"`
	CategoryID       uint      `json:"category_id"`
	SellerID         uint      `json:"seller_id"`
	CategoryDiscount uint      `json:"discountPercentage"`
	StartDate        time.Time `json:"start_date"`
	EndDate          time.Time `json:"end_date"`
	Status           string    `json:"status"`
}
