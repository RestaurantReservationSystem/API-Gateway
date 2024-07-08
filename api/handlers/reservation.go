package handlers

import (
	pb "api_get_way/genproto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) CreateReservationHandler(ctx *gin.Context) {

	resquest := pb.CreateReservationRequest{}

	err := ctx.ShouldBind(&resquest)

	if err != nil {
		BadRequest(ctx, err)
		return
	}
	
	resp, err := h.ReservationService.CreateReservation(ctx, &resquest)

	if err != nil {
		InternalServerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)

}

func (h *Handler) UpdateReservationHandler(ctx *gin.Context) {


	request := pb.CreateReservationRequest{}

	err := ctx.ShouldBind(&request)

	if err != nil {
		BadRequest(ctx, err)
		return
	}

	resp, err := h.ReservationService.CreateReservation(ctx, &request)

	if err != nil {
		InternalServerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)

}

func (h *Handler) DeleteReservationHandler(ctx *gin.Context) {

	id := ctx.Param("id")

	_, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Id notogri formatda",
		})
		return
	}

	request := pb.IdRequest{}
	request.Id = id

	resp, err := h.ReservationService.DeleteReservation(ctx, &request)

	if err != nil {
		InternalServerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) GetByIdReservationHandler(ctx *gin.Context) {

	id := ctx.Param("id")

	_, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Id notogri formatda",
		})
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

	resp, err := h.ReservationService.GetAllReservation(ctx, &request)

	if err != nil {
		InternalServerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)

}
