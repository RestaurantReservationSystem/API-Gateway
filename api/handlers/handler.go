package handlers

import (
	"api_get_way/genproto"
	"go.uber.org/zap"

)

type Handler struct {
	PaymentService     genproto.PaymentServiceClient
	ReservationService genproto.ReservationServiceClient
	UsersService       genproto.UserServiceClient
	Log                *zap.Logger
}

func NewHandler(payment genproto.PaymentServiceClient, reservation genproto.ReservationServiceClient, user genproto.UserServiceClient, l *zap.Logger) *Handler {
	return &Handler{
		PaymentService:     payment,
		ReservationService: reservation,
		UsersService:       user,
		Log: l,
		
	}
}
