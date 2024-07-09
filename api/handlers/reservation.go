package handlers

import (
	pb "api_get_way/genproto"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateReservationHandler 		handles the creation of a new user
// @Summary 		Create Menu
// @Description 	Create page
// @Tags 			Menu
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param   		Create  body     pb.RegisterUserRequest  true   "Create"
// @Success 		200   {string}      "Create Successful"
// @Failure 		401   {string}   string    "Error while Created"
// @Router 			/api/reservation/create [post]

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

// UpdateReservationHandler 		handles the creation of a new user
// @Summary 		Update Menu
// @Description 	Update page
// @Tags 			Menu
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param   		Update  body     pb.RegisterUserRequest  true   "Update"
// @Success 		200   {string}      "Update Successful"
// @Failure 		401   {string}   string    "Error while Created"
// @Router 			/api/reservation/update/:id [put]

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

// DeleteReservationHandler 		handles the creation of a new user
// @Summary 		Delete Menu
// @Description 	Delete page
// @Tags 			Menu
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param   		Delete  body     pb.RegisterUserRequest  true   "Delete"
// @Success 		200   {string}      "Delete Successful"
// @Failure 		401   {string}   string    "Error while Created"
// @Router 			/api/reservation/delete/:id [delete]

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

// GetByIdReservationHandler 		handles the creation of a new user
// @Summary 		GetId Menu
// @Description 	GetId page
// @Tags 			Menu
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param   		Get  body     pb.RegisterUserRequest  true   "Get"
// @Success 		200   {string}      "GetId Successful"
// @Failure 		401   {string}   string    "Error while Created"
// @Router 			/api/reservation/get_id/:id [get]

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

// GetAllReservationHandler 		handles the creation of a new user
// @Summary 		GetAll Menu
// @Description 	GetAll page
// @Tags 			Menu
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param   		Get  body     pb.RegisterUserRequest  true   "Get"
// @Success 		200   {string}      "GetAll Successful"
// @Failure 		401   {string}   string    "Error while Created"
// @Router 			/api/reservation/get_all [get]

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
