package handlers

import (
	"api_get_way/genproto"
)

type Handler struct {
	PaymentService     genproto.PaymentServiceClient
	ReservationService genproto.ReservationServiceClient
}

func NewHandler(py genproto.PaymentServiceClient, rs genproto.ReservationServiceClient) *Handler {
	return &Handler{PaymentService: py, ReservationService: rs}
}
