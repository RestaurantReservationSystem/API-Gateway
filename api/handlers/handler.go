package handlers

import (
	"api_get_way/genproto"
)

type Handler struct {
	PaymentService     genproto.PaymentServiceClient
	ReservationService genproto.ReservationServiceClient
	UsersService       genproto.UserServiceClient
}

func NewHandler(payment genproto.PaymentServiceClient, reservation genproto.ReservationServiceClient, user genproto.UserServiceClient) *Handler {
	return &Handler{
		PaymentService:     payment,
		ReservationService: reservation,
		UsersService:       user,
	}
}
