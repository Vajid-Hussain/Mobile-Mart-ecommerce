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

func (r *JwtTokenUseCase) ValidateJwtToken(id string) (string, error) {
	status, err := r.repo.GetUserStatus(id)
	if err != nil {
		return "", err
	}

	if status != "active" {
		return "", errors.New("user is not active user")
	}

	return status, nil
}
