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
		return nil, nil, errors.New("attempt to convert string to int , page")
	}

	limits, err := strconv.Atoi(limit)
	if err != nil {
		return nil, nil, errors.New("attempt to convert string to int , page limit")
	}

	if pageNO < 1 {
		return nil, nil, errors.New("page must start from one")
	}

	offSet := (pageNO * limits) - limits

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
