package handlers

import (
	pb "api_get_way/genproto"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateMenuHandler(ctx *gin.Context) {
	request := pb.CreateMenuRequest{}

	err := ctx.ShouldBind(&request)

	if err != nil {
		BadRequest(ctx, err)
	}

	Parse(ctx, request.RestaurantId)

	if request.Description == "" || request.Name == "" {
		BadRequest(ctx, fmt.Errorf("malumot toliq emas"))
	}

	_, err = h.Menu.CreateMenu(ctx, &request)

	if err != nil {
		InternalServerError(ctx, err)
	}

	Created(ctx, nil)
}

func (h *Handler) UpdateMenuHandler(ctx *gin.Context) {

	request := pb.UpdateMenuRequest{}

	err := ctx.ShouldBind(&request)
	if err != nil {
		BadRequest(ctx, err)
	}

	Parse(ctx, request.RestaurantId)
	Parse(ctx, request.Id)

	if request.Name == "" || request.Description == "" {
		BadRequest(ctx, fmt.Errorf("maliumot toliq emas"))
	}

	_, err = h.Menu.UpdateMenu(ctx, &request)

	if err != nil {
		InternalServerError(ctx, err)
	}

	OK(ctx, nil)
}

func (h *Handler) DeleteMenuHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	Parse(ctx, id)

	_, err := h.Menu.DeleteMenu(ctx, &pb.IdRequest{Id: id})

	if err != nil {
		InternalServerError(ctx, err)
	}

	Created(ctx, nil)
}

func (h *Handler) GetByIdMenuHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	Parse(ctx, id)

	resp, err := h.Menu.GetByIdMenu(ctx, &pb.IdRequest{Id: id})

	if err != nil {
		InternalServerError(ctx, err)
	}

	ctx.JSON(http.StatusOK, resp)

}

func (h *Handler) GetAllMenuHandler(ctx *gin.Context) {

	request := pb.GetAllMenuRequest{}

	err := ctx.ShouldBind(&request)

	if err != nil {
		BadRequest(ctx, err)
	}

	Parse(ctx, request.RestaurantId)

	if request.Name == "" || request.Description == "" {
		BadRequest(ctx, fmt.Errorf("malumot toliq emas"))
	}

	resp, err := h.Menu.GetAllMenu(ctx, &request)

	if err != nil {
		InternalServerError(ctx, err)
	}

	ctx.JSON(http.StatusOK, resp)
}
