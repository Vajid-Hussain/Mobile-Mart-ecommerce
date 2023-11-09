package usecase

import (
	"errors"
	"fmt"

	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
	interfaces "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/repository/interface"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
	"github.com/go-playground/validator/v10"
)

type sellerUseCase struct {
	repo interfaces.ISellerRepo
}

func NewSellerUseCase(sellerRepo interfaces.ISellerRepo) interfaceUseCase.IVenderUseCase {
	return &sellerUseCase{repo: sellerRepo}
}

func (r *sellerUseCase) SellerSignup(signData requestmodel.SellerSignup) (responsemodel.SellerSignupRes, error) {
	var SellerSignupRes responsemodel.SellerSignupRes

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(signData)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "Name":
					SellerSignupRes.Name = "don't empty fill with a name"
				case "Email":
					SellerSignupRes.Email = "email format is wrong"
				case "Password":
					SellerSignupRes.Password = "Password atleast 4 digit"
				case "ConfirmPassword":
					SellerSignupRes.ConfirmPassword = "must match with password"
				case "GST_NO":
					SellerSignupRes.GST_NO = "Must have fifteen digit number"
				}
			}
		}
		fmt.Println(SellerSignupRes)
		return SellerSignupRes, errors.New("login credential not obey")
	}


	return SellerSignupRes, nil
}

