package usecase

import (
	"errors"

	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/config"
	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
	interfaces "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/repository/interface"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/service"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
	"github.com/go-playground/validator/v10"

	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/utils/helper"
)

type sellerUseCase struct {
	repo  interfaces.ISellerRepo
	token config.Token
}

func NewSellerUseCase(sellerRepo interfaces.ISellerRepo, token *config.Token) interfaceUseCase.ISellerUseCase {
	return &sellerUseCase{repo: sellerRepo,
		token: *token}
}

func (r *sellerUseCase) SellerSignup(sellerSignupData *requestmodel.SellerSignup) (*responsemodel.SellerSignupRes, error) {
	var SellerSignupRes responsemodel.SellerSignupRes

	count, err := r.repo.IsSellerExist(sellerSignupData.Email)
	if err != nil {
		return &SellerSignupRes, err
	} else {
		if count >= 1 {
			return &SellerSignupRes, errors.New("seller exist with same email id, ")
		}
	}

	hashPassword := helper.HashPassword(sellerSignupData.Password)
	sellerSignupData.Password = hashPassword

	err = r.repo.CreateSeller(sellerSignupData)
	if err != nil {
		return &SellerSignupRes, err
	}

	SellerSignupRes.Result = "Registeration saved ! Your request is now in processing. You will receive a confirmation once you have been admitted and granted access to start selling."
	return &SellerSignupRes, nil
}

func (r *sellerUseCase) SellerLogin(loginData *requestmodel.SellerLogin) (*responsemodel.SellerLoginRes, error) {
	var loginResponse responsemodel.SellerLoginRes

	hashedPassword, sellerID, status, err := r.repo.GetHashPassAndStatus(loginData.Email)
	if err != nil {
		return &loginResponse, err
	}

	if status == "block" {
		return &loginResponse, errors.New("vender blocked by admin")
	}

	if status == "pending" {
		return &loginResponse, errors.New("your request under process pls whait ")
	}

	err = helper.CompairPassword(hashedPassword, loginData.Password)
	if err != nil {
		return &loginResponse, err
	}

	accessToken, err := service.GenerateAcessToken(r.token.SellerSecurityKey, sellerID)
	if err != nil {
		return &loginResponse, err
	}

	refreshToken, err := service.GenerateRefreshToken(r.token.SellerSecurityKey)
	if err != nil {
		return &loginResponse, err
	}

	loginResponse.AccessToken = accessToken
	loginResponse.RefreshToken = refreshToken

	return &loginResponse, nil
}

func (r *sellerUseCase) GetAllSellers(page string, limit string) (*[]responsemodel.SellerDetails, *int, error) {
	ch := make(chan int)

	go r.repo.SellerCount(ch)
	count := <-ch

	offSet, limits, err := helper.Pagination(page, limit)
	if err != nil {
		return nil, &count, err
	}

	SellerDetails, err := r.repo.AllSellers(offSet, limits)
	if err != nil {
		return nil, nil, err
	}

	return SellerDetails, &count, nil
}

func (r *sellerUseCase) BlockSeller(id string) error {
	err := r.repo.BlockSeller(id)
	if err != nil {
		return err
	}
	err = r.repo.BlockInventoryOfSeller(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *sellerUseCase) ActiveSeller(id string) error {
	err := r.repo.UnblockSeller(id)
	if err != nil {
		return err
	}
	err = r.repo.ActiveInventoryOfSeller(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *sellerUseCase) GetAllPendingSellers(page string, limit string) (*[]responsemodel.SellerDetails, error) {

	offSet, limits, err := helper.Pagination(page, limit)
	if err != nil {
		return nil, err
	}

	SellerDetails, err := r.repo.GetPendingSellers(offSet, limits)
	if err != nil {
		return nil, err
	}

	return SellerDetails, nil
}

func (r *sellerUseCase) FetchSingleVender(id string) (*responsemodel.SellerDetails, error) {
	sellerData, err := r.repo.GetSingleSeller(id)
	if err != nil {
		return nil, err
	}
	return sellerData, nil
}

// ------------------------------------------Seller Profile------------------------------------\\

func (r *sellerUseCase) GetSellerProfile(userID string) (*responsemodel.SellerProfile, error) {
	sellerDetails, err := r.repo.GetSellerProfile(userID)
	if err != nil {
		return nil, err
	}
	return sellerDetails, nil
}

func (r *sellerUseCase) UpdateSellerProfile(editedProfile *requestmodel.SellerEditProfile) (*responsemodel.SellerProfile, error) {

	exist, err := r.repo.IsSellerExist(editedProfile.Email)
	if err != nil {
		return nil, err
	}

	if exist > 0 {
		return nil, errors.New("seller exist with same email id , change email")
	}

	if editedProfile.Password != editedProfile.ConfirmPassword {
		return nil, errors.New("password and confirmpassword is not match")
	}

	if editedProfile.Password != "" {
		editedProfile.Password = helper.HashPassword(editedProfile.Password)
	}

	SellerProfile, err := r.repo.GetSellerProfile(editedProfile.ID)
	if err != nil {
		return nil, err
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(editedProfile)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				fieldName := e.Field()
				switch fieldName {
				case "ID":
					editedProfile.ID = SellerProfile.ID
				case "Name":
					editedProfile.Name = SellerProfile.Name
				case "Email":
					editedProfile.Email = SellerProfile.Email
				case "Password":
					editedProfile.Password = SellerProfile.Password
				case "Description":
					editedProfile.Description = SellerProfile.Description
				}
			}
		}

	}

	sellerEdittedProfile, err := r.repo.UpdateSellerProfile(editedProfile)
	if err != nil {
		return nil, err
	}
	return sellerEdittedProfile, nil
}

func (r *sellerUseCase) GetSellerDashbord(sellerID string) (*responsemodel.DashBord, error) {
	var dashBord responsemodel.DashBord
	var err error
	dashBord.SellerID = sellerID

	dashBord.DeliveredOrders, err = r.repo.GetDashBordOrderCount(sellerID, "delivered")
	if err != nil {
		return nil, err
	}

	dashBord.OngoingOrders, err = r.repo.GetDashBordOrderCount(sellerID, "processing")
	if err != nil {
		return nil, err
	}

	dashBord.CancelledOrders, err = r.repo.GetDashBordOrderCount(sellerID, "cancelled")
	if err != nil {
		return nil, err
	}

	dashBord.TotalOrders, err = r.repo.GetDashBordOrderCount(sellerID, "")
	if err != nil {
		return nil, err
	}

	dashBord.TotalRevenue, err = r.repo.GetDashBordOrderSum(sellerID, "price")
	if err != nil {
		return nil, err
	}

	dashBord.TotalSelledProduct, err = r.repo.GetDashBordOrderSum(sellerID, "quantity")
	if err != nil {
		return nil, err
	}

	dashBord.AdminCredit, err = r.repo.GetSellerCredit(sellerID)
	if err != nil {
		return nil, err
	}

	dashBord.LowStockProductID, err = r.repo.GetLowStokesProduct(sellerID)
	if err != nil {
		return nil, err
	}
	return &dashBord, nil
}
