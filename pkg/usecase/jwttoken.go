package usecase

import (
	"errors"

	interfaces "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/repository/interface"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
)

type JwtTokenUseCase struct {
	repo interfaces.IJwtTokenRepository
}

func NewJwtTokenUseCase(repo interfaces.IJwtTokenRepository) interfaceUseCase.IJwtTokenUseCase {
	return &JwtTokenUseCase{repo: repo}
}

func (r *JwtTokenUseCase) GetDataForCreteAccessToken(id string) (string, error) {
	status, err := r.repo.GetSellerStatus(id)
	if err != nil {
		return "", err
	}

	if status != "active" {
		return "", errors.New("not a active user")
	}

	return status, nil
}

func (r *JwtTokenUseCase) GetStatusOfUser(id string) (string, error) {
	status, err := r.repo.GetUserStatus(id)
	if err != nil {
		return "", err
	}

	if status != "active" {
		return "", errors.New("not a active user")
	}

	return status, nil
}
