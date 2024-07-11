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
		h.log.Error("error")
		BadRequest(ctx, err)
		return
	}

	if Parse(payment.ReservationId) {
		h.log.Error("error")
		BadRequest(ctx, fmt.Errorf("invalid reservation ID"))
		return
	}

	if !isValidPaymentMethod(payment.PaymentMethod) {
		h.log.Error("error")
		BadRequest(ctx, fmt.Errorf("invalid payment method"))
		return
	}

	_, err = h.PaymentService.CreatePayment(ctx, &payment)
	if err != nil {
		h.log.Error("error")
		InternalServerError(ctx, err)
		return
	}
	h.log.Info("iahkadu")
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
		h.log.Error("error")
		BadRequest(ctx, err)
		return
	}

	payment.Id = ctx.Param("id")
	if Parse(payment.Id) {
		h.log.Error("error")
		BadRequest(ctx, fmt.Errorf("invalid payment ID"))
		return
	}

	if Parse(payment.ReservationId) {
		h.log.Error("error")
		BadRequest(ctx, fmt.Errorf("invalid reservation ID"))
		return
	}

	if !isValidPaymentMethod(payment.PaymentMethod) {
		h.log.Error("error")
		BadRequest(ctx, fmt.Errorf("invalid payment method"))
		return
	}

	_, err = h.PaymentService.UpdatePayment(ctx, &payment)
	if err != nil {
		h.log.Error("error")
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
		h.log.Error("error")
		BadRequest(ctx, fmt.Errorf("invalid payment ID"))
		return
	}

	_, err := h.PaymentService.DeletePayment(ctx, &pb.IdRequest{Id: id})
	if err != nil {
		h.log.Error("error")
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
		h.log.Error("error")
		BadRequest(ctx, fmt.Errorf("invalid payment ID"))
		return
	}

	resp, err := h.PaymentService.GetByIdPayment(ctx, &pb.IdRequest{Id: id})
	if err != nil {
		h.log.Error("error")
		InternalServerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// GetAllPaymentHandler handles fetching all payments.
// @Summary Get All Payments
// @Description Get all payments
// @Tags Payment
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param request query genproto.GetAllPaymentRequest true "Get All Payments"
// @Success 200 {object} genproto.PaymentsResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/payment/get_all [get]
func (h *Handler) GetAllPaymentHandler(ctx *gin.Context) {
	req := pb.GetAllPaymentRequest{}

	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		h.log.Error("error")
		BadRequest(ctx, err)
		return
	}

	if Parse(req.ReservationId) {
		h.log.Error("error")
		BadRequest(ctx, fmt.Errorf("invalid reservation ID"))
		return
	}

	resp, err := h.PaymentService.GetAllPayment(ctx, &req)
	if err != nil {
		h.log.Error("error")
		InternalServerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
