package handler

import (
	"net/http"

	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type UserHandler struct{
	userUseCase	interfaceUseCase.IuserUseCase
}

func NewUserHandler(userUseCase interfaceUseCase.IuserUseCase) *UserHandler{
	return &UserHandler{userUseCase: userUseCase}
}

func (u *UserHandler) UserSignup(c *gin.Context){
	var userSignupData requestmodel.UserDetails
	if err:= c.BindJSON(userSignupData); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	}

	u.userUseCase.UserSignup(userSignupData)
}
