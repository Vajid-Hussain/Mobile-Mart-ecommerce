package handler

import (
	"net/http"

	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	resCustomError "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel/custom_error"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel/response"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	useCase interfaceUseCase.IOrderUseCase
}

func NewOrderHandler(orderUseCase interfaceUseCase.IOrderUseCase) *OrderHandler {
	return &OrderHandler{useCase: orderUseCase}
}

// @Summary Create User Order
// @Description Create a new order by the user.
// @Tags UserOrders
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Security Refreshtoken
// @Param order body requestmodel.Order true "Order details for creating"
// @Success 201 {object} response.Response "Order created successfully"
// @Failure 400 {object} response.Response "Bad request"
// @Router /order [post]
func (u *OrderHandler) NewOrder(c *gin.Context) {

	var order *requestmodel.Order

	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	if err := c.BindJSON(&order); err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, resCustomError.BindingConflict, nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	order.UserID = userID

	orderDetais, err := u.useCase.NewOrder(order)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", orderDetais, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// @Summary Get User Orders
// @Description Retrieve all orders for the user.
// @Tags UserOrders
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Security Refreshtoken
// @Success 200 {object} response.Response "Successfully retrieved user orders"
// @Failure 400 {object} response.Response "Bad request"
// @Router /order [get]
func (u *OrderHandler) ShowAbstractOrders(c *gin.Context) {

	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	result, err := u.useCase.OrderShowcase(userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", result, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// @Summary Get User Order Details
// @Description Retrieve details about a specific user order.
// @Tags UserOrders
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Security Refreshtoken
// @Param orderID path string true "Order ID in the URL path"
// @Success 200 {object} response.Response "Successfully retrieved user order details"
// @Failure 400 {object} response.Response "Bad request"
// @Router /order/{orderID} [get]
func (u *OrderHandler) SingleOrderDetails(c *gin.Context) {

	orderID, _ := c.Params.Get("orderID")

	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	orderDetais, err := u.useCase.SingleOrder(orderID, userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", orderDetais, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// @Summary Cancel User Order
// @Description Cancel an order for the user.
// @Tags UserOrders
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Security Refreshtoken
// @Param orderID query string true "Order ID in the query parameter"
// @Success 200 {object} response.Response "Order canceled successfully"
// @Failure 400 {object} response.Response "Bad request"
// @Router /order [patch]
func (u *OrderHandler) CancelUserOrder(c *gin.Context) {

	userID, exist := c.MustGet("UserID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}
	orderID := c.Query("orderID")

	orderDetais, err := u.useCase.CancelUserOrder(orderID, userID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", orderDetais, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// ------------------------------------------Seller Control Orders------------------------------------\\

// @Summary Get Seller Order
// @Description Retrieve a single order for the seller.
// @Tags SellerOrders
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Security Refreshtoken
// @Success 200 {object} response.Response "Successfully retrieved the seller order"
// @Failure 400 {object} response.Response "Bad request"
// @Router /seller/order [get]
func (u *OrderHandler) GetSellerOrders(c *gin.Context) {
	sellerID, exist := c.MustGet("SellerID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	remainingQuery := " IN ('processing','delivered')"
	orderDetais, err := u.useCase.GetSellerOrders(sellerID, remainingQuery)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", orderDetais, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// @Summary Get Seller Processing Orders
// @Description Retrieve still ongoing orders for the seller.
// @Tags SellerOrders
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Security Refreshtoken
// @Success 200 {object} response.Response "Successfully retrieved seller processing orders"
// @Failure 400 {object} response.Response "Bad request"
// @Router /seller/order/processing [get]
func (u *OrderHandler) GetSellerOrdersProcessing(c *gin.Context) {
	sellerID, exist := c.MustGet("SellerID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	remainingQuery := " IN ('processing')"
	orderDetais, err := u.useCase.GetSellerOrders(sellerID, remainingQuery)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", orderDetais, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// @Summary Get Seller Delivered Orders
// @Description Retrieve delivered orders for the seller.
// @Tags SellerOrders
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Security Refreshtoken
// @Success 200 {object} response.Response "Successfully retrieved seller delivered orders"
// @Failure 400 {object} response.Response "Bad request"
// @Router /seller/order/delivered [get]
func (u *OrderHandler) GetSellerOrdersDeliverd(c *gin.Context) {
	sellerID, exist := c.MustGet("SellerID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	remainingQuery := " IN ('delivered')"
	orderDetais, err := u.useCase.GetSellerOrders(sellerID, remainingQuery)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", orderDetais, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// @Summary Confirm Seller Order
// @Description Confirm an order for the seller.
// @Tags SellerOrders
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Security Refreshtoken
// @Param orderID query string true "Order ID in the query parameter"
// @Success 200 {object} response.Response "Order confirmed successfully"
// @Failure 400 {object} response.Response "Bad request"
// @Router /seller/order [patch]
func (u *OrderHandler) ConfirmDeliverd(c *gin.Context) {
	sellerID, exist := c.MustGet("SellerID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	orderID := c.Query("orderID")

	orderDetais, err := u.useCase.ConfirmDeliverd(sellerID, orderID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", orderDetais, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// @Summary Cancel Seller Order
// @Description Cancel an order for the seller.
// @Tags SellerOrders
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Security Refreshtoken
// @Param orderID path string true "Order ID in the URL path"
// @Success 200 {object} response.Response "Order canceled successfully"
// @Failure 400 {object} response.Response "Bad request"
// @Router /seller/order/{orderID}/cancel [patch]
func (u *OrderHandler) CancelOrder(c *gin.Context) {

	sellerID, exist := c.MustGet("SellerID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}
	orderID := c.Param("orderID")

	orderDetais, err := u.useCase.CancelOrder(orderID, sellerID)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", orderDetais, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// ------------------------------------------Sales Report------------------------------------\\

func (u *OrderHandler) SalesReportByYear(c *gin.Context) {

	sellerID, exist := c.MustGet("SellerID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}
	year := c.Query("year")
	partialQuery := " EXTRACT(YEAR FROM order_date)=" + year

	report, err := u.useCase.GetSalesReportByYear(sellerID, partialQuery)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", report, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *OrderHandler) SalesReportByMonth(c *gin.Context) {

	sellerID, exist := c.MustGet("SellerID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}
	year := c.Query("year")
	month := c.Query("month")
	partialQuery := " EXTRACT(YEAR FROM order_date)=" + year + " AND EXTRACT(Month FROM order_date)=" + month

	report, err := u.useCase.GetSalesReportByYear(sellerID, partialQuery)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", report, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *OrderHandler) SalesReportByWeek(c *gin.Context) {

	sellerID, exist := c.MustGet("SellerID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}
	year := c.Query("year")
	week := c.Query("week")
	partialQuery := " EXTRACT(YEAR FROM order_date)=" + year + " AND EXTRACT(Week FROM order_date)=" + week

	report, err := u.useCase.GetSalesReportByYear(sellerID, partialQuery)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", report, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *OrderHandler) SalesReportByDay(c *gin.Context) {

	sellerID, exist := c.MustGet("SellerID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}
	year := c.Query("year")
	day := c.Query("day")
	partialQuery := " EXTRACT(YEAR FROM order_date)=" + year + " AND EXTRACT(Day FROM order_date)=" + day

	report, err := u.useCase.GetSalesReportByYear(sellerID, partialQuery)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", report, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}
