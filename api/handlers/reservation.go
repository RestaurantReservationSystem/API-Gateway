package handlers

import (
	pb "api_get_way/genproto"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateReservationHandler handles the creation of a new reservation.
// @Summary Create Reservation
// @Description Create a new reservation
// @Tags Reservation
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param Create body genproto.CreateReservationRequest true "Create Reservation"
// @Success 200 {object} genproto.ReservationResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/reservation/create [post]
func (h *Handler) CreateReservationHandler(ctx *gin.Context) {
	request := pb.CreateReservationRequest{}
	err := ctx.ShouldBind(&request)
	if err != nil {
		h.log.Error("error")
		BadRequest(ctx, err)
		return
	}

	if _, err := uuid.Parse(request.UserId); err != nil {
		h.log.Error("error")
		BadRequest(ctx, err)
		return
	}

	if _, err := uuid.Parse(request.RestaurantId); err != nil {
		h.log.Error("error")
		BadRequest(ctx, err)
		return
	}

	if request.Status != "pending" && request.Status != "confirmed" && request.Status != "cancelled" {
		h.log.Error("error")
		BadRequest(ctx, fmt.Errorf("invalid status"))
		return
	}

	resp, err := h.ReservationService.CreateReservation(ctx, &request)
	if err != nil {
		h.log.Error("error")
		InternalServerError(ctx, err)
		return
	}
	h.log.Info("ishladi")
	ctx.JSON(http.StatusOK, resp)
}

// UpdateReservationHandler handles the update of a reservation.
// @Summary Update Reservation
// @Description Update an existing reservation
// @Tags Reservation
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "Reservation ID"
// @Param Update body genproto.UpdateReservationRequest true "Update Reservation"
// @Success 200 {object} genproto.ReservationResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/reservation/update/{id} [put]
func (h *Handler) UpdateReservationHandler(ctx *gin.Context) {
	request := pb.UpdateReservationRequest{}
	err := ctx.ShouldBind(&request)
	if err != nil {
		h.log.Error("error")
		BadRequest(ctx, err)
		return
	}

	if _, err := uuid.Parse(request.UserId); err != nil {
		h.log.Error("error")
		BadRequest(ctx, err)
		return
	}

	if _, err := uuid.Parse(request.RestaurantId); err != nil {
		h.log.Error("error")
		BadRequest(ctx, err)
		return
	}

	if request.Status != "" && request.Status != "pending" && request.Status != "confirmed" && request.Status != "cancelled" {
		h.log.Error("error")
		BadRequest(ctx, fmt.Errorf("invalid status"))
		return
	}

	resp, err := h.ReservationService.UpdateReservation(ctx, &request)
	if err != nil {
		h.log.Error("error")
		InternalServerError(ctx, err)
		return
	}
	h.log.Info("ishladi")
	ctx.JSON(http.StatusOK, resp)
}

// DeleteReservationHandler handles the deletion of a reservation.
// @Summary Delete Reservation
// @Description Delete an existing reservation
// @Tags Reservation
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "Reservation ID"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/reservation/delete/{id} [delete]
func (h *Handler) DeleteReservationHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		h.log.Error("error")
		BadRequest(ctx, err)
		return
	}

	_, err := h.ReservationService.DeleteReservation(ctx, &pb.IdRequest{Id: id})
	if err != nil {
		h.log.Error("error")
		InternalServerError(ctx, err)
		return
	}
	h.log.Info("ishladi")
	OK(ctx)
}

// GetByIdReservationHandler handles fetching a reservation by its ID.
// @Summary Get Reservation by ID
// @Description Get a reservation by its ID
// @Tags Reservation
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "Reservation ID"
// @Success 200 {object} genproto.ReservationResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/reservation/get_id/{id} [get]
func (h *Handler) GetByIdReservationHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		h.log.Error("error")
		BadRequest(ctx, err)
		return
	}

	resp, err := h.ReservationService.GetByIdReservation(ctx, &pb.IdRequest{Id: id})
	if err != nil {
		h.log.Error("error")
		InternalServerError(ctx, err)
		return
	}
	h.log.Info("ishladi")
	ctx.JSON(http.StatusOK, resp)
}

// GetAllReservationHandler handles fetching all reservations.
// @Summary Get All Reservations
// @Description Get all reservations
// @Tags Reservation
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param query query genproto.GetAllReservationRequest true "Get All Reservations"
// @Success 200 {object} genproto.ReservationsResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/reservation/get_all [get]
func (h *Handler) GetAllReservationHandler(ctx *gin.Context) {
	request := pb.GetAllReservationRequest{}
	err := ctx.ShouldBind(&request)
	if err != nil {
		h.log.Error("error")
		BadRequest(ctx, err)
		return
	}

	if request.UserId != "" {
		if _, err := uuid.Parse(request.UserId); err != nil {
			h.log.Error("error")
			BadRequest(ctx, err)
			return
		}
	}

	if request.RestaurantId != "" {
		if _, err := uuid.Parse(request.RestaurantId); err != nil {
			BadRequest(ctx, err)
			return
		}
	}

	if request.Status != "" && request.Status != "pending" && request.Status != "confirmed" && request.Status != "cancelled" {
		BadRequest(ctx, fmt.Errorf("invalid status"))
		return
	}

	resp, err := h.ReservationService.GetAllReservation(ctx, &request)
	if err != nil {
		h.log.Error("error")
		InternalServerError(ctx, err)
		return
	}
	h.log.Info("ishladi")
	ctx.JSON(http.StatusOK, resp)
}
