package usecase

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/config"
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
	interfaces "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/repository/interface"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/service"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/utils/helper"
	"github.com/go-playground/validator/v10"
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

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(adminData)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "Email":
					adminLoginRes.Email = "email id is wrong "
				case "Password":
					adminLoginRes.Password = "password have four or more digit"
				}
			}
		}
		return &adminLoginRes, errors.New("did't fullfill the login requirement ")
	}

	HashedPassword, err := r.repo.GetPassword(adminData.Email)
	if err != nil {
		fmt.Println(err, "---------", HashedPassword)

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

func (r *adminUsecase) GetAllUsers(page string, limit string) (*[]responsemodel.UserDetails, *int, error) {
	ch := make(chan int)

	go r.repo.UserCount(ch)
	count := <-ch

	pageNO, err := strconv.Atoi(page)
	if err != nil {
		return nil, nil, responsemodel.ConversionOFPageErr
	}

	limits, err := strconv.Atoi(limit)
	if err != nil {
		return nil, nil, responsemodel.ConversionOfLimitErr
	}

	if pageNO < 1 {
		return nil, nil, responsemodel.PaginationError
	}

	offSet := (pageNO * limits) - limits
	limits = pageNO * limits

	userDetails, err := r.repo.AllUsers(offSet, limits)
	if err != nil {
		return nil, nil, err
	}

	return userDetails, &count, nil
}

func (r *adminUsecase) BlcokUser(id string) error {
	err := r.repo.BlockUser(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *adminUsecase) UnblockUser(id string) error {
	err := r.repo.UnblockUser(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *adminUsecase) GetAllSellers(page string, limit string) (*[]responsemodel.SellerDetails, *int, error) {
	ch := make(chan int)

	go r.repo.SellerCount(ch)
	count := <-ch

	pageNO, err := strconv.Atoi(page)
	if err != nil {
		return nil, nil, responsemodel.ConversionOFPageErr
	}

	limits, err := strconv.Atoi(limit)
	if err != nil {
		return nil, nil, responsemodel.ConversionOfLimitErr
	}

	if pageNO < 1 {
		return nil, nil, responsemodel.PaginationError
	}

	offSet := (pageNO * limits) - limits
	limits = pageNO * limits

	SellerDetails, err := r.repo.AllSellers(offSet, limits)
	if err != nil {
		return nil, nil, err
	}

	return SellerDetails, &count, nil
}

func (r *adminUsecase) BlockSeller(id string) error {
	err := r.repo.BlockSeller(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *adminUsecase) UnblockSeller(id string) error {
	err := r.repo.UnblockSeller(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *adminUsecase) GetAllPendingSellers(page string, limit string) (*[]responsemodel.SellerDetails, error) {

	pageNO, err := strconv.Atoi(page)
	if err != nil {
		return nil, responsemodel.ConversionOFPageErr
	}

	limits, err := strconv.Atoi(limit)
	if err != nil {
		return nil, responsemodel.ConversionOfLimitErr
	}

	if pageNO < 1 {
		return nil, responsemodel.PaginationError
	}
	offSet := (pageNO * limits) - limits
	limits = pageNO * limits

	SellerDetails, err := r.repo.GetPendingSellers(offSet, limits)
	if err != nil {
		return nil, err
	}

	return SellerDetails, nil
}

func (r *adminUsecase) FetchSingleVender(id string) (*responsemodel.SellerDetails, error) {
	sellerData, err := r.repo.GetSingleSeller(id)
	if err != nil {
		return nil, err
	}
	return sellerData, nil
}
