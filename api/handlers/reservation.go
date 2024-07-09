package handlers

import (
	pb "api_get_way/genproto"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) CreateReservationHandler(ctx *gin.Context) {

	request := pb.CreateReservationRequest{}

	err := ctx.ShouldBind(&request)
	if err != nil {
		BadRequest(ctx, err)
		return
	}

	_, err = uuid.Parse(request.UserId)
	if err != nil {
		BadRequest(ctx, err)
		return
	}

	_, err = uuid.Parse(request.RestaurantId)
	if err != nil {
		BadRequest(ctx, err)
		return
	}

	if request.Status != "pending" && request.Status != "confirmed" && request.Status != "cancelled" {
		BadRequest(ctx, fmt.Errorf("malumot toliq emas"))
		return
	}

	_, err = h.ReservationService.CreateReservation(ctx, &request)

	if err != nil {
		InternalServerError(ctx, err)
		return
	}

	Created(ctx, nil)

}

func (h *Handler) UpdateReservationHandler(ctx *gin.Context) {

	request := pb.UpdateReservationRequest{}

	err := ctx.ShouldBind(&request)

	if err != nil {
		BadRequest(ctx, err)
		return
	}
	_, err = uuid.Parse(request.UserId)
	if err != nil {
		BadRequest(ctx, err)
		return
	}

	_, err = uuid.Parse(request.RestaurantId)
	if err != nil {
		BadRequest(ctx, err)
		return
	}

	if request.Status != "" && request.Status != "pending" && request.Status != "confirmed" && request.Status != "cancelled" {
		BadRequest(ctx, fmt.Errorf("malumot toliq emas"))
		return
	}

	_, err = h.ReservationService.UpdateReservation(ctx, &request)

	if err != nil {
		InternalServerError(ctx, err)
		return
	}

	OK(ctx, nil)

}

func (h *Handler) DeleteReservationHandler(ctx *gin.Context) {

	id := ctx.Param("id")

	_, err := uuid.Parse(id)
	if err != nil {
		BadRequest(ctx, err)
		return
	}

	request := pb.IdRequest{}
	request.Id = id

	_, err = h.ReservationService.DeleteReservation(ctx, &request)

	if err != nil {
		InternalServerError(ctx, err)
		return
	}

	Created(ctx, nil)
}

func (h *Handler) GetByIdReservationHandler(ctx *gin.Context) {

	id := ctx.Param("id")

	_, err := uuid.Parse(id)
	if err != nil {
		BadRequest(ctx, err)
		return
	}

	request := pb.IdRequest{}
	request.Id = id

	resp, err := h.ReservationService.GetByIdReservation(ctx, &request)

	if err != nil {
		InternalServerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)

}

func (h *Handler) GetAllReservationHandler(ctx *gin.Context) {

	request := pb.GetAllReservationRequest{}

	err := ctx.ShouldBind(&request)

	if err != nil {
		BadRequest(ctx, err)
		return
	}

	if request.UserId != "" {
		_, err := uuid.Parse(request.UserId)
		if err != nil {
			BadRequest(ctx, err)
			return
		}
	}
	if request.RestaurantId != "" {
		_, err := uuid.Parse(request.RestaurantId)
		if err != nil {
			BadRequest(ctx, err)
			return
		}
	}

	if request.Status != "" {
		if request.Status != "" && request.Status != "pending" && request.Status != "confirmed" && request.Status != "cancelled" {
			BadRequest(ctx, fmt.Errorf("malumot toliq emas"))
			return
		}
	}

	resp, err := h.ReservationService.GetAllReservation(ctx, &request)

	if err != nil {
		InternalServerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
