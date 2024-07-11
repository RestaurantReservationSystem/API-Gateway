package handlers

import (
	pb "api_get_way/genproto"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var allowedPaymentMethods = []string{"card", "payment", "cash"}

func isValidPaymentMethod(method string) bool {
	for _, m := range allowedPaymentMethods {
		if m == method {
			return true
		}
	}
	return false
}

// CreatePaymentHandler handles the creation of a new payment.
// @Summary Create Payment
// @Description Create a new payment
// @Tags Payment
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param Create body genproto.CreatePaymentRequest true "Create Payment"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/payment/create [post]
func IsValidOffset(offset string) (int, error) {
	if offset == "" {
		offset += "0"
	}
	offset1, err := strconv.Atoi(offset)
	if err != nil {

		return 0, err
	}
	return offset1, nil

}

func (h *Handler) CreatePaymentHandler(ctx *gin.Context) {
	payment := pb.CreatePaymentRequest{}
	err := ctx.ShouldBindJSON(&payment)
	if err != nil {
		h.Log.Error("error")
		BadRequest(ctx, err)
		return
	}

	if Parse(payment.ReservationId) {
		h.Log.Error("error")
		BadRequest(ctx, fmt.Errorf("invalid reservation ID"))
		return
	}

	if !isValidPaymentMethod(payment.PaymentMethod) {
		h.Log.Error("error")
		BadRequest(ctx, fmt.Errorf("invalid payment method"))
		return
	}

	_, err = h.PaymentService.CreatePayment(ctx, &payment)
	if err != nil {
		h.Log.Error("error")
		InternalServerError(ctx, err)
		return
	}
	h.Log.Info("iahkadu")
	Created(ctx)
}

// UpdatePaymentHandler handles the update of a payment.
// @Summary Update Payment
// @Description Update an existing payment
// @Tags Payment
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "Payment ID"
// @Param Update body genproto.UpdatePaymentRequest true "Update Payment"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/payment/update/{id} [put]
func (h *Handler) UpdatePaymentHandler(ctx *gin.Context) {
	payment := pb.UpdatePaymentRequest{}
	err := ctx.ShouldBindJSON(&payment)
	if err != nil {
		h.Log.Error("error")
		BadRequest(ctx, err)
		return
	}

	payment.Id = ctx.Param("id")
	if Parse(payment.Id) {
		h.Log.Error("error")
		BadRequest(ctx, fmt.Errorf("invalid payment ID"))
		return
	}

	if Parse(payment.ReservationId) {
		h.Log.Error("error")
		BadRequest(ctx, fmt.Errorf("invalid reservation ID"))
		return
	}

	if !isValidPaymentMethod(payment.PaymentMethod) {
		h.Log.Error("error")
		BadRequest(ctx, fmt.Errorf("invalid payment method"))
		return
	}

	_, err = h.PaymentService.UpdatePayment(ctx, &payment)
	if err != nil {
		h.Log.Error("error")
		InternalServerError(ctx, err)
		return
	}

	OK(ctx)
}

// DeletePaymentHandler handles the deletion of a payment.
// @Summary Delete Payment
// @Description Delete an existing payment
// @Tags Payment
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "Payment ID"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/payment/delete/{id} [delete]
func (h *Handler) DeletePaymentHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if Parse(id) {
		h.Log.Error("error")
		BadRequest(ctx, fmt.Errorf("invalid payment ID"))
		return
	}

	_, err := h.PaymentService.DeletePayment(ctx, &pb.IdRequest{Id: id})
	if err != nil {
		h.Log.Error("error")
		InternalServerError(ctx, err)
		return
	}

	OK(ctx)
}

// GetByIdPaymentHandler handles fetching a payment by its ID.
// @Summary Get Payment by ID
// @Description Get a payment by its ID
// @Tags Payment
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "Payment ID"
// @Success 200 {object} genproto.PaymentResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/payment/get_id/{id} [get]
func (h *Handler) GetByIdPaymentHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if Parse(id) {
		h.Log.Error("error")
		BadRequest(ctx, fmt.Errorf("invalid payment ID"))
		return
	}

	resp, err := h.PaymentService.GetByIdPayment(ctx, &pb.IdRequest{Id: id})
	if err != nil {
		h.Log.Error("error")
		InternalServerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// GetAllPaymentHandler  filtering and pagination.
// @Summary Get All Payment
// @Description Retrieve  filtering and pagination.
// @Tags Order
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param payment_status query string false "Filter by order item payment_status"
// @Param payment_method query string false "Filter by menu item payment_method"
// @Param reservation_id query string false "Filter by restaurant reservation_id"
// @Param amount query string false "Filter by restaurant amount"
// @Param limit query int false "Number of items to return"
// @Param offset query int false "Offset for pagination"
// @Success 200 {object} genproto.PaymentsResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/payment [get]
func (h *Handler) GetAllPaymentHandler(ctx *gin.Context) {
	request := pb.GetAllPaymentRequest{}

	request.PaymentStatus = ctx.Query("payment_status")
	request.PaymentMethod = ctx.Query("payment_method")
	payment := ctx.Query("amount")
	if IsAmount(payment) {
		BadRequest(ctx, fmt.Errorf("error -> payment is validate"))
		return
	}
	if Parse(request.ReservationId) {
		BadRequest(ctx, fmt.Errorf("invalid reservation ID"))
		return
	}
	limit := ctx.Query("limit")
	limit1, err := IsLimitOffsetValidate(limit)
	if err != nil {
		h.Log.Error("error")
		BadRequest(ctx, err)
		return
	}

	if Parse(request.ReservationId) {
		h.Log.Error("error")
		BadRequest(ctx, fmt.Errorf("invalid reservation ID"))

	offset := ctx.Query("offset")
	offset1, err := IsLimitOffsetValidate(offset)
	if err != nil {
		BadRequest(ctx, err)
		return
	}
	request.LimitOffset.Offset = int64(offset1)
	request.LimitOffset.Limit = int64(limit1)

	resp, err := h.PaymentService.GetAllPayment(ctx, &request)
	if err != nil {
		h.Log.Error("error")
		InternalServerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
}