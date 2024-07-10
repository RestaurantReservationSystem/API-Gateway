package handlers

import "api_get_way/genproto/auth"



type Handler struct {
	PaymentService     
	ReservationService rese
	UsersService       auth.UserServiceClient
}

func NewHandler(py genproto.PaymentServiceClient, rs genproto.ReservationServiceClient, user genproto.UserServiceClient) *Handler {
	return &Handler{
		PaymentService:     py,
		ReservationService: rs,
		UsersService:       user}
}
