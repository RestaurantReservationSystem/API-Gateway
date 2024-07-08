package handlers

import (
	pb "api_get_way/genproto"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	_ "github.com/google/uuid"
	"strconv"
)

func isValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
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

func (h *Handler) CreatePayment(gn *gin.Context) {
	payment := pb.CreatePaymentRequest{}
	err := gn.ShouldBindJSON(&payment)
	if err != nil {
		BadRequest(gn, err)
		return
	}
	newUuid := isValidUUID(payment.ReservationId)
	if !newUuid {
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

func (h *Handler) UpdatePayment(gn *gin.Context) {
	payment := pb.UpdatePaymentRequest{}
	err := gn.ShouldBindJSON(&payment)
	if err != nil {
		BadRequest(gn, err)
		return
	}
	payment.Id = gn.Param("id")
	isValidUUID(payment.Id)
	newUuid := isValidUUID(payment.ReservationId)
	if !newUuid {
		BadRequest(gn, err)
		return
	}

	if !isValidPaymentMethod(payment.PaymentMethod) {
		BadRequest(gn, err)
	}

	_, err = h.PaymentService.UpdatePayment(gn, &payment)

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

func (h *Handler) CreatePayment(gn *gin.Context) {
	payment := pb.CreatePaymentRequest{}
	err := gn.ShouldBindJSON(&payment)
	if err != nil {
		BadRequest(gn, err)
		return
	}
	newUuid := isValidUUID(payment.ReservationId)
	if newUuid == false {
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

func (h *Handler) UpdatePayment(gn *gin.Context) {
	payment := pb.UpdatePaymentRequest{}
	err := gn.ShouldBindJSON(&payment)
	if err != nil {
		BadRequest(gn, err)
		return
	}
	payment.Id = gn.Param("id")
	isValidUUID(payment.Id)
	newUuid := isValidUUID(payment.ReservationId)
	if newUuid == false {
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
