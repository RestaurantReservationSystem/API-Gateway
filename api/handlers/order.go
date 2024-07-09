package handlers

import (
	pb "api_get_way/genproto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateOrderHandler(ctx *gin.Context) {

	request := pb.CreateOrderRequest{}

	err := ctx.ShouldBind(&request)
	if err != nil {
		BadRequest(ctx, err)
	}

	if request.Quantity == "" {
		BadRequest(ctx, err)
	}
	Parse(ctx, request.ReservationId)
	Parse(ctx, request.MenuItemId)

	_, err = h.Order.CreateOrder(ctx, &request)

	if err != nil {
		InternalServerError(ctx, err)
	}

	Created(ctx, nil)

}

func (h *Handler) UpdateOrderHandler(ctx *gin.Context) {
	request := pb.UpdateOrderRequest{}

	err := ctx.ShouldBind(&request)
	if err != nil {
		BadRequest(ctx, err)
	}

	if request.Quantity == "" {
		BadRequest(ctx, err)
	}
	Parse(ctx, request.Id)
	Parse(ctx, request.ReservationId)
	Parse(ctx, request.MenuItemId)

	_, err = h.Order.UpdateOrder(ctx, &request)

	if err != nil {
		InternalServerError(ctx, err)
	}

	Created(ctx, nil)

}

func (h Handler) DeleteOrder(ctx *gin.Context) {
	id := ctx.Param("id")

	Parse(ctx, id)

	_, err := h.Order.DeleteOrder(ctx, &pb.IdRequest{Id: id})

	if err != nil {
		InternalServerError(ctx, err)
	}

	Created(ctx, nil)
}

func (h *Handler) GetByIdOrderHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	Parse(ctx, id)

	resp, err := h.Order.GetByIdOrder(ctx, &pb.IdRequest{Id: id})

	if err != nil {
		InternalServerError(ctx, err)
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) GetAllOrderHandler(ctx *gin.Context) {
	request := pb.GetAllOrderRequest{}

	err := ctx.ShouldBind(&request)
	if err != nil {
		BadRequest(ctx, err)
	}

	if request.Quantity == "" {
		BadRequest(ctx, err)
	}
	Parse(ctx, request.ReservationId)
	Parse(ctx, request.MenuItemId)

	resp, err := h.Order.GetAllOrder(ctx, &request)

	if err != nil {
		InternalServerError(ctx, err)
	}

	ctx.JSON(http.StatusOK, resp)
}
