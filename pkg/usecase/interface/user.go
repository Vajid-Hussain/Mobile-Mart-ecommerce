package interfaceUseCase

import requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"

type IuserUseCase interface{
	UserSignup(requestmodel.UserDetails)
}