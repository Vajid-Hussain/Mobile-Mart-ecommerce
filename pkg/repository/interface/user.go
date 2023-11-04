package interfaces

import requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"

type IUserRepo interface {
	CreateUser(requestmodel.UserDetails)
}
