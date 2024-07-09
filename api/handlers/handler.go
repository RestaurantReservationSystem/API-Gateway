package handlers

import (
	"api_get_way/genproto"
)

type Handler struct {
	PaymentService     genproto.PaymentServiceClient
	ReservationService genproto.ReservationServiceClient
	UsersService       genproto.UserServiceClient
	Restaran           genproto.RestaurantServiceClient
	Order              genproto.OrderServiceClient
	Menu               genproto.MenuServiceClient
}

func NewHandler(py genproto.PaymentServiceClient, rs genproto.ReservationServiceClient, user genproto.UserServiceClient) *Handler {
	return &Handler{
		PaymentService:     py,
		ReservationService: rs,
		UsersService:       user}
}
