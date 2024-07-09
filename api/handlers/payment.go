package handlers

import (
	pb "api_get_way/genproto"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IsValidLimit(limit string) (int, error) {
	if limit == "" {
		limit += "0"
	}
	limit1, err := strconv.Atoi(limit)
	if err != nil {
		return 0, err
	}
	return limit1, nil
}

var allowedPaymentMethods = []string{"card", "payment", "cash"}

func isValidPaymentMethod(method string) bool {
	for _, m := range allowedPaymentMethods {
		if m == method {
			return true
		}
	}
	return false
}

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

// CreatePayment 		handles the creation of a new user
// @Summary 		Create Menu
// @Description 	Create page
// @Tags 			Menu
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param   		Create  body     pb.RegisterUserRequest  true   "Create"
// @Success 		200   {string}      "Create Successful"
// @Failure 		401   {string}   string    "Error while Created"
// @Router 			/api/order/get_id/:id [get]

func (h *Handler) CreatePayment(gn *gin.Context) {
	payment := pb.CreatePaymentRequest{}
	err := gn.ShouldBindJSON(&payment)
	if err != nil {
		BadRequest(gn, err)
		return
	}
	newUuid := Parse(payment.ReservationId)
	if newUuid {
		BadRequest(gn, err)
		return
	}

	if !isValidPaymentMethod(payment.PaymentMethod) {
		BadRequest(gn, err)
	}

	_, err = h.PaymentService.CreatePayment(gn, &payment)
	if err != nil {
		InternalServerError(gn, err)
		return
	}
	Created(gn, err)

}

// UpdatePayment 		handles the creation of a new user
// @Summary 		Update Menu
// @Description 	Update page
// @Tags 			Menu
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param   		Update  body     pb.RegisterUserRequest  true   "Update"
// @Success 		200   {string}      "Update Successful"
// @Failure 		401   {string}   string    "Error while Created"
// @Router 			/api/order/update/:id [put]

func (h *Handler) UpdatePayment(gn *gin.Context) {
	payment := pb.UpdatePaymentRequest{}
	err := gn.ShouldBindJSON(&payment)
	if err != nil {
		BadRequest(gn, err)
		return
	}
	payment.Id = gn.Param("id")

	Parse(payment.Id)

	newUuid := Parse(payment.ReservationId)
	if newUuid {
		BadRequest(gn, err)
		return
	}

	if !isValidPaymentMethod(payment.PaymentMethod) {
		BadRequest(gn, err)
		return
	}

	_, err = h.PaymentService.UpdatePayment(gn, &payment)
	if err != nil {
		InternalServerError(gn, err)
		return
	}

	OK(gn, nil)
}

// DeletePayment 		handles the creation of a new user
// @Summary 		Delete Menu
// @Description 	Delete page
// @Tags 			Menu
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param   		Delete  body     pb.RegisterUserRequest  true   "Delete"
// @Success 		200   {string}      "Update Successful"
// @Failure 		401   {string}   string    "Error while Created"
// @Router 			/api/payment/delete/:id [delete]

func (h *Handler) DeletePayment(ctx *gin.Context) {
	id := ctx.Param("id")

	if Parse(id) {
		BadRequest(ctx, fmt.Errorf("id hato"))
		return
	}

	_, err := h.PaymentService.DeletePayment(ctx, &pb.IdRequest{Id: id})

	if err != nil {
		InternalServerError(ctx, err)
		return
	}

	Created(ctx, nil)
}

// GetByIdPayment 		handles the creation of a new user
// @Summary 		GetId Menu
// @Description 	GetId page
// @Tags 			Menu
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param   		GetId  body     pb.RegisterUserRequest  true   "GetId"
// @Success 		200   {string}      "GetId Successful"
// @Failure 		401   {string}   string    "Error while Created"
// @Router 			/api/payment/get_id/:id [get]

func (h *Handler) GetByIdPayment(ctx *gin.Context) {
	id := ctx.Param("id")

	if Parse(id) {
		BadRequest(ctx, fmt.Errorf("id hato"))
		return
	}

	resp, err := h.PaymentService.GetByIdPayment(ctx, &pb.IdRequest{Id: id})

	if err != nil {
		InternalServerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// GetAllPaymentHandler 		handles the creation of a new user
// @Summary 		GetAll Menu
// @Description 	GetAll page
// @Tags 			Menu
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param   		GetAll  body     pb.RegisterUserRequest  true   "GetAll"
// @Success 		200   {string}      "GetAll Successful"
// @Failure 		401   {string}   string    "Error while Created"
// @Router 			/api/payment/get_all/ [get]

func (h *Handler) GetAllPaymentHandler(ctx *gin.Context) {
	req := pb.GetAllPaymentRequest{}

	err := ctx.ShouldBind(&req)

	if err != nil {
		BadRequest(ctx, err)
		return
	}

	if Parse(req.ReservationId) {
		BadRequest(ctx, fmt.Errorf("id hato"))
		return
	}

	resp, err := h.PaymentService.GetAllPayment(ctx, &req)

	if err != nil {
		InternalServerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
