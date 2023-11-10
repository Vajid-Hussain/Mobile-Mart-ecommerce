package interfaces

import requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"

type ISellerRepo interface {
	IsSellerExist(string) (int, error)
	CreateSeller(*requestmodel.SellerSignup) error
	GetEmailAndStatus(string) (string, string, error)
	GetPasswordByMail(string) string
}
