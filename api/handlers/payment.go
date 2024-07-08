package handlers

import (
	pb "api_get_way/genproto"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePayment(gn gin.Context) {

	request := pb.CreatePaymentRequest{
		ReservationId: nil,
		Amount:        nil,
		PaymentMethod: nil,
		PaymentStatus: nil,
	}
	h.PaymentService.CreatePayment(gn)
}
