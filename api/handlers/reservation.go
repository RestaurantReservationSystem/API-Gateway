package handlers

import (
	pb "api_get_way/genproto"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
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
	// time input this example type 2024-10-14T23:34:34Z
	if err != nil {
		h.log.Error("error")
		BadRequest(ctx, fmt.Errorf("failed to bind request: %v", err))
		return
	}

	if _, err := uuid.Parse(request.UserId); err != nil {
		h.log.Error("error")
		BadRequest(ctx, fmt.Errorf("invalid user ID: %v", err))
		return
	}

	if _, err := uuid.Parse(request.RestaurantId); err != nil {
		h.log.Error("error")
		
		BadRequest(ctx, fmt.Errorf("invalid restaurant ID: %v", err))
		return
	}

	if request.Status != "pending" && request.Status != "confirmed" && request.Status != "cancelled" {
		h.log.Error("error")
		BadRequest(ctx, fmt.Errorf("invalid status: %v", request.Status))
		return
	}

	time1, err := time.Parse(time.RFC3339, request.ReservationTime)
	if err != nil {
		h.log.Error("error")
		InternalServerError(ctx, err)
		return
	}
	h.log.Info("ishladi")
	ctx.JSON(http.StatusOK, resp)
		
	}

	request.ReservationTime = time1.Format(time.RFC3339)
	//_, err = h.ReservationService.GetByIdRestaurant(ctx, &pb.IdRequest{Id: request.RestaurantId})
	//if err != nil {
	//	BadRequest(ctx, fmt.Errorf("Restaurant bazada mavjud emas"))
	//	return
	//}
	//_, err = h.UsersService.GetByIdUser(ctx, &pb.IdRequest{Id: request.UserId})
	//if err != nil {
	//	BadRequest(ctx, fmt.Errorf("User bazada mavjud emas"))
	//	return
	//}
	_, err = h.ReservationService.CreateReservation(ctx, &request)
	if err != nil {
		InternalServerError(ctx, fmt.Errorf("failed to create reservation: %v", err))
		return
	}
	OK(ctx)

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
		BadRequest(ctx, fmt.Errorf("failed to bind request: %v", err))
		return
	}

	if _, err := uuid.Parse(request.UserId); err != nil {
		h.log.Error("error")
		BadRequest(ctx, fmt.Errorf("invalid user ID: %v", err))
		return
	}

	if _, err := uuid.Parse(request.RestaurantId); err != nil {
		h.log.Error("error")
		BadRequest(ctx, fmt.Errorf("invalid restaurant ID: %v", err))
		return
	}

	if request.Status != "" && request.Status != "pending" && request.Status != "confirmed" && request.Status != "cancelled" {
		h.log.Error("error")
		BadRequest(ctx, fmt.Errorf("invalid status: %v", request.Status))
		return
	}
	_, err = h.ReservationService.GetByIdReservation(ctx, &pb.IdRequest{Id: request.Id})
	if err != nil {
		h.log.Error("error")
		InternalServerError(ctx, err)
		return
	}
	h.log.Info("ishladi")
	ctx.JSON(http.StatusOK, resp)
=======
		BadRequest(ctx, fmt.Errorf("Bu id Databaseda yoq"))
		return
	}
	//_, err = h.ReservationService.GetByIdRestaurant(ctx, &pb.IdRequest{Id: request.RestaurantId})
	//if err != nil {
	//	BadRequest(ctx, fmt.Errorf("Restaurant bazada mavjud emas"))
	//	return
	//}
	//_, err = h.UsersService.GetByIdUser(ctx, &pb.IdRequest{Id: request.UserId})
	//if err != nil {
	//	BadRequest(ctx, fmt.Errorf("User bazada mavjud emas"))
	//	return
	//}
	_, err = h.ReservationService.GetByIdReservation(ctx, &pb.IdRequest{Id: request.UserId})
	if err != nil {
		BadRequest(ctx, fmt.Errorf("resevation bazada mavjud emas"))
		return
	}
	_, err = h.ReservationService.UpdateReservation(ctx, &request)
	if err != nil {
		InternalServerError(ctx, fmt.Errorf("failed to update reservation: %v", err))
		return
	}

	Created(ctx)
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
	_, err := h.ReservationService.GetByIdReservation(ctx, &pb.IdRequest{})
	if err != nil {
		BadRequest(ctx, fmt.Errorf("Bu id Databaseda yoq"))
		return
	}
	_, err = h.ReservationService.DeleteReservation(ctx, &pb.IdRequest{Id: id})
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

// GetAllReservationHandler retrieves reservations with optional filtering and pagination.
// @Summary Get All Reservation
// @Description Retrieve reservations with optional filtering and pagination.
// @Tags Reservation
// @Accept json
// @Produce json
// @Param status query string false "Filter by reservation status"
// @Param user_id query string false "Filter by user ID"
// @Param reservation_time query string false "Filter by reservation time"
// @Param restaurant_id query string false "Filter by restaurant ID"
// @Param limit query int false "Number of items to return"
// @Param offset query int false "Offset for pagination"
// @Success 200 {object} genproto.ReservationsResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/reservation/get_all [get]
func (h *Handler) GetAllReservationHandler(ctx *gin.Context) {
	request := pb.GetAllReservationRequest{}
	limit := ctx.Query("limit")
	limit1, err := IsLimitOffsetValidate(limit)
	if err != nil {
		h.log.Error("error")
		BadRequest(ctx, fmt.Errorf("invalid limit: %v", err))
		return
	}
	request.LimitOffset = &pb.Filter{
		Limit: int64(limit1),
	}

	if request.UserId != "" {
		if _, err := uuid.Parse(request.UserId); err != nil {
			h.log.Error("error")
			BadRequest(ctx, err)
	offset := ctx.Query("offset")
	offset1, err := IsLimitOffsetValidate(offset)
	if err != nil {
		BadRequest(ctx, fmt.Errorf("invalid offset: %v", err))
		return
	}
	request.LimitOffset.Offset = int64(offset1)
	userID := ctx.Query("user_id")
	if userID != "" {
		if _, err := uuid.Parse(userID); err != nil {
			BadRequest(ctx, fmt.Errorf("invalid user ID: %v", err))
			return
		}
		request.UserId = userID
	}

	// Validate restaurant_id if provided
	restaurantID := ctx.Query("restaurant_id")
	if restaurantID != "" {
		if _, err := uuid.Parse(restaurantID); err != nil {
			BadRequest(ctx, fmt.Errorf("invalid restaurant ID: %v", err))
			return
		}
		request.RestaurantId = restaurantID
	}

	request.Status = ctx.Query("status")

	resp, err := h.ReservationService.GetAllReservation(ctx, &request)
	if err != nil {

		h.log.Error("error")
		InternalServerError(ctx, fmt.Errorf("failed to get reservations: %v", err))
		return
	}
	h.log.Info("ishladi")
	ctx.JSON(http.StatusOK, resp)
}
