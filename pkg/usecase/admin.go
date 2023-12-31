package usecase

import (
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/config"
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
	interfaces "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/repository/interface"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/service"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/utils/helper"
)

type adminUsecase struct {
	repo             interfaces.IAdminRepository
	tokenSecurityKey config.Token
}

func NewAdminUseCase(adminRepository interfaces.IAdminRepository, key *config.Token) interfaceUseCase.IAdminUseCase {
	return &adminUsecase{repo: adminRepository,
		tokenSecurityKey: *key}
}

func (r *adminUsecase) AdminLogin(adminData *requestmodel.AdminLoginData) (*responsemodel.AdminLoginRes, error) {
	var adminLoginRes responsemodel.AdminLoginRes

	HashedPassword, err := r.repo.GetPassword(adminData.Email)
	if err != nil {
		return nil, err
	}

	err = helper.CompairPassword(HashedPassword, adminData.Password)
	if err != nil {
		return nil, err
	}

	token, err := service.GenerateRefreshToken(r.tokenSecurityKey.AdminSecurityKey)
	if err != nil {
		return nil, err
	}

	adminLoginRes.Token = token
	return &adminLoginRes, nil
}

func (r *adminUsecase) GetSellerDetailsForAdminDashBord() (*responsemodel.AdminDashBord, error) {
	var dashBord responsemodel.AdminDashBord
	var err error

	dashBord.TotalSellers, err = r.repo.GetSellerDetailsForDashBord("")
	if err != nil {
		return nil, err
	}

	dashBord.ActiveSellers, err = r.repo.GetSellerDetailsForDashBord("active")
	if err != nil {
		return nil, err
	}

	dashBord.BlockedSellers, err = r.repo.GetSellerDetailsForDashBord("block")
	if err != nil {
		return nil, err
	}

	dashBord.PendingSellers, err = r.repo.GetSellerDetailsForDashBord("pending")
	if err != nil {
		return nil, err
	}

	dashBord.TotalOrders, dashBord.TotalRevenue, err = r.repo.TotalRevenue()
	if err != nil {
		return nil, err
	}

	dashBord.TotalCredit, err = r.repo.GetNetCredit()
	if err != nil {
		return nil, err
	}

	return &dashBord, nil
}
