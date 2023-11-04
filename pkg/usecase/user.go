package usecase

import (
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	interfaces "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/repository/interface"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
)

type userUseCase struct {
	repo interfaces.IUserRepo
}

func NewUserUseCase(userRepository interfaces.IUserRepo) interfaceUseCase.IuserUseCase {
	return &userUseCase{repo: userRepository}
}

func (u *userUseCase) UserSignup(userData requestmodel.UserDetails) {
	u.repo.CreateUser(userData)
}
