package middlewire

import (
	"fmt"
	"net/http"
	"os"

	responsemodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel/response"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/service"
	"github.com/gin-gonic/gin"
)

func UserTokenVerify(c *gin.Context) {
	var responses responsemodel.TokenVerificationMiddlewire
	token := c.Request.Header.Get("Authorization")

	securityKet := os.Getenv("user324567ytre")
	if securityKet == "" {
		fmt.Println("NO environmental variable available at this name")
	}

	_, err := service.VerifyToken(token, securityKet)
	if err != nil {
		responses.Error = err.Error()
		finalResponse:=response.Responses(http.StatusUnauthorized, "pls login", responses, nil)
		c.JSON(http.StatusUnauthorized, finalResponse)
	}
	c.Next()
}
