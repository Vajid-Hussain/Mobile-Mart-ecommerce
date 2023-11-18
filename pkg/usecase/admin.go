package usecase

import (
	"errors"
	"fmt"
	"mime/multipart"

	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/config"
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
	interfaces "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/repository/interface"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/service"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/utils/helper"
	// "github.com/go-playground/validator/v10"
)

type adminUsecase struct {
	repo             interfaces.IAdminRepository
	tokenSecurityKey config.Token
	s3               config.S3Bucket
}

func NewAdminUseCase(adminRepository interfaces.IAdminRepository, key *config.Token, s3aws *config.S3Bucket) interfaceUseCase.IAdminUseCase {
	return &adminUsecase{repo: adminRepository,
		tokenSecurityKey: *key,
		s3:               *s3aws}
}

func (r *adminUsecase) AdminLogin(adminData *requestmodel.AdminLoginData) (*responsemodel.AdminLoginRes, error) {
	var adminLoginRes responsemodel.AdminLoginRes

	// validate := validator.New(validator.WithRequiredStructEnabled())
	// err := validate.Struct(adminData)
	// if err != nil {
	// 	if ve, ok := err.(validator.ValidationErrors); ok {
	// 		for _, e := range ve {
	// 			switch e.Field() {
	// 			case "Email":
	// 				adminLoginRes.Email = "email id is wrong "
	// 			case "Password":
	// 				adminLoginRes.Password = "password have four or more digit"
	// 			}
	// 		}
	// 	}
	// 	return &adminLoginRes, errors.New("did't fullfill the login requirement ")
	// }

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

func (r *adminUsecase) ImageUpload(img *multipart.FileHeader) error {

	sess := service.CreateSession(&r.s3)

	s3Sess := service.CreateS3Session(sess)

	fmt.Println(*s3Sess)
	err := service.UploadObject(img, sess)
	if err != nil {
		return errors.New("can't upload images")
	}

	return nil
}
