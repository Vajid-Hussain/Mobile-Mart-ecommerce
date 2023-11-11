package interfaces

import requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"

type ISellerRepo interface {
	IsSellerExist(string) (int, error)
	CreateSeller(*requestmodel.SellerSignup) error
	GetHashPassAndStatus(string) (string, string, string, error)
	GetPasswordByMail(string) string
}
