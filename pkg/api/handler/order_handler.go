package handler

import (
	"fmt"
	"net/http"

	requestmodel "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/requestModel"
	resCustomError "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel/custom_error"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/models/responseModel/response"
	interfaceUseCase "github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/usecase/interface"
	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/utils/helper"
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

// @Summary Get Seller Sales Report
// @Description Retrieve the seller sales report for the specified year.
// @Tags Seller Sales Report
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Security Refreshtoken
// @Param year query int true "Year for which the report is requested"
// @Success 200 {object} response.Response "Seller sales report retrieved successfully"
// @Failure 400 {object} response.Response "Bad request. Please provide a valid year."
// @Router /seller/report [get]
func (u *OrderHandler) SalesReportByYear(c *gin.Context) {

	sellerID, exist := c.MustGet("SellerID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}
	year := c.Query("year")
	remainingQuery := " EXTRACT(YEAR FROM order_date)=" + year

	report, err := u.useCase.GetSalesReportByYear(sellerID, remainingQuery)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", report, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// @Summary Get Seller Sales Report for a Specific Month
// @Description Retrieve the seller sales report for the specified year and month.
// @Tags Seller Sales Report
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Security Refreshtoken
// @Param year query int true "Year for which the report is requested"
// @Param month query int true "Month for which the report is requested (1-12)"
// @Success 200 {object} response.Response "Seller sales report retrieved successfully"
// @Failure 400 {object} response.Response "Bad request. Please provide a valid year and month."
// @Router /seller/report/month [get]
func (u *OrderHandler) SalesReportByMonth(c *gin.Context) {

	sellerID, exist := c.MustGet("SellerID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}
	year := c.Query("year")
	month := c.Query("month")
	remainingQuery := " EXTRACT(YEAR FROM order_date)=" + year + " AND EXTRACT(Month FROM order_date)=" + month

	report, err := u.useCase.GetSalesReportByYear(sellerID, remainingQuery)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", report, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// @Summary Get Seller Sales Report for a Specific Week
// @Description Retrieve the seller sales report for the specified year and week.
// @Tags Seller Sales Report
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Security Refreshtoken
// @Param year query int true "Year for which the report is requested"
// @Param week query int true "Week for which the report is requested"
// @Success 200 {object} response.Response "Seller sales report retrieved successfully"
// @Failure 400 {object} response.Response "Bad request. Please provide a valid year and week."
// @Router /seller/report/week [get]
func (u *OrderHandler) SalesReportByWeek(c *gin.Context) {

	sellerID, exist := c.MustGet("SellerID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}
	year := c.Query("year")
	week := c.Query("week")
	remainingQuery := " EXTRACT(YEAR FROM order_date)=" + year + " AND EXTRACT(Week FROM order_date)=" + week

	report, err := u.useCase.GetSalesReportByYear(sellerID, remainingQuery)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", report, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

// @Summary Get Seller Sales Report for a Specific Day
// @Description Retrieve the seller sales report for the specified year, month, and day.
// @Tags Seller Sales Report
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Security Refreshtoken
// @Param year query int true "Year for which the report is requested"
// @Param month query int true "Month for which the report is requested (1-12)"
// @Param day query int true "Day for which the report is requested (1-31)"
// @Success 200 {object} response.Response "Seller sales report retrieved successfully"
// @Failure 400 {object} response.Response "Bad request. Please provide a valid year, month, and day."
// @Router /seller/report/day [get]
func (u *OrderHandler) SalesReportByDay(c *gin.Context) {

	sellerID, exist := c.MustGet("SellerID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}
	year := c.Query("year")
	month := c.Query("month")
	day := c.Query("day")
	remainingQuery := " EXTRACT(YEAR FROM order_date)=" + year + " AND EXTRACT(Month FROM order_date)=" + month + " AND EXTRACT(Day FROM order_date)=" + day

	report, err := u.useCase.GetSalesReportByYear(sellerID, remainingQuery)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", report, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *OrderHandler) SalesReportCustomDays(c *gin.Context) {
	sellerID, exist := c.MustGet("SellerID").(string)
	if !exist {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}
	day := c.Query("days")

	report, err := u.useCase.GetSalesReportByDays(sellerID, day)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", report, nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}

func (u *OrderHandler) OnlinePayment(c *gin.Context) {
	userID := c.Query("userID")
	fmt.Println("****", userID)
	orderDetails, err := u.useCase.OnlinePayment(userID)
	fmt.Println("###", orderDetails, err)
	if err != nil {
		c.HTML(http.StatusBadRequest, "razopay.html", gin.H{"badRequest": "Refine your request"})
	} else {
		c.HTML(http.StatusOK, "razopay.html", orderDetails)
	}
}

func (u *OrderHandler) VerifyOnlinePayment(c *gin.Context) {
	var onlinePaymentDetails requestmodel.OnlinePaymentVerification

	// userID, exist := c.MustGet("UserID").(string)
	// if !exist {
	// 	finalReslt := response.Responses(http.StatusBadRequest, "", nil, resCustomError.NotGetUserIdInContexr)
	// 	c.JSON(http.StatusBadRequest, finalReslt)
	// 	return
	// }

	if err := c.BindJSON(&onlinePaymentDetails); err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, resCustomError.BindingConflict, nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	data, err := helper.Validation(onlinePaymentDetails)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", data, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	err = u.useCase.OnlinePaymentVerification(&onlinePaymentDetails)
	if err != nil {
		finalReslt := response.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
	} else {
		finalReslt := response.Responses(http.StatusOK, "", "", nil)
		c.JSON(http.StatusOK, finalReslt)
	}
}
