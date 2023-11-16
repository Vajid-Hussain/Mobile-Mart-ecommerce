package interfaceUseCase

import (
	"mime/multipart"

	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
)

type IAdminUseCase interface {
	AdminLogin(*requestmodel.AdminLoginData) (*responsemodel.AdminLoginRes, error)
	ImageUpload(*multipart.FileHeader) error
}
