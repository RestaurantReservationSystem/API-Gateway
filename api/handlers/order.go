package handlers

import (
	pb "api_get_way/genproto"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateOrderHandler 		handles the creation of a new user
// @Summary 		Create Menu
// @Description 	Create page
// @Tags 			Menu
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param   		Create  body     pb.RegisterUserRequest  true   "Create"
// @Success 		200   {string}      "Create Successful"
// @Failure 		401   {string}   string    "Error while Created"
// @Router 			/api/order/create [post]

func (h *Handler) CreateOrderHandler(ctx *gin.Context) {

	request := pb.CreateOrderRequest{}

	err := ctx.ShouldBind(&request)
	if err != nil {
		BadRequest(ctx, err)
	}

	if request.Quantity == "" {
		BadRequest(ctx, err)
	}
	
	if Parse( request.ReservationId) || Parse( request.MenuItemId){
		BadRequest(ctx,fmt.Errorf("id hato"))
		return
	}

	_, err = h.Order.CreateOrder(ctx, &request)

	if err != nil {
		InternalServerError(ctx, err)
		return
	}

	Created(ctx, nil)

}

// UpdateOrderHandler 		handles the creation of a new user
// @Summary 		Update Menu
// @Description 	Update page
// @Tags 			Menu
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param   		Create  body     pb.RegisterUserRequest  true   "Update"
// @Success 		200   {string}      "Update Successful"
// @Failure 		401   {string}   string    "Error while Created"
// @Router 			/api/order/Update/:id [put]

func (h *Handler) UpdateOrderHandler(ctx *gin.Context) {
	request := pb.UpdateOrderRequest{}

	err := ctx.ShouldBind(&request)
	if err != nil {
		BadRequest(ctx, err)
		return
	}

	if request.Quantity == "" {
		BadRequest(ctx, err)
		return
	}

	if Parse( request.ReservationId) || Parse( request.MenuItemId) || Parse(request.Id){
		BadRequest(ctx,fmt.Errorf("id hato"))
		return
	}

	_, err = h.Order.UpdateOrder(ctx, &request)

	if err != nil {
		InternalServerError(ctx, err)
		return
	}

	Created(ctx, nil)

}

// DeleteOrder 		handles the creation of a new user
// @Summary 		Delete Menu
// @Description 	Delete page
// @Tags 			Menu
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param   		Create  body     pb.RegisterUserRequest  true   "Delete"
// @Success 		200   {string}      "Delete Successful"
// @Failure 		401   {string}   string    "Error while Created"
// @Router 			/api/order/delete/:id [delete]

func (h Handler) DeleteOrder(ctx *gin.Context) {
	id := ctx.Param("id")

	if Parse(id){
		BadRequest(ctx,fmt.Errorf("id hato"))
		return
	}

	_, err := h.Order.DeleteOrder(ctx, &pb.IdRequest{Id: id})

	if err != nil {
		InternalServerError(ctx, err)
		return
	}

	Created(ctx, nil)
}

// GetByIdOrderHandler 		handles the creation of a new user
// @Summary 		GetId Menu
// @Description 	GetId page
// @Tags 			Menu
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param   		Create  body     pb.RegisterUserRequest  true   "GetId"
// @Success 		200   {string}      "GetAll Successful"
// @Failure 		401   {string}   string    "Error while Created"
// @Router 			/api/order/get_id/:id [get]


func (h *Handler) GetByIdOrderHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if Parse(id){
		BadRequest(ctx,fmt.Errorf("id hato"))
		return
	}

	resp, err := h.Order.GetByIdOrder(ctx, &pb.IdRequest{Id: id})

	if err != nil {
		InternalServerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// GetAllOrderHandler 		handles the creation of a new user
// @Summary 		GetAll Menu
// @Description 	GetAll page
// @Tags 			Menu
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param   		Create  body     pb.RegisterUserRequest  true   "GetAll"
// @Success 		200   {string}      "GetAll Successful"
// @Failure 		401   {string}   string    "Error while Created"
// @Router 			/api/order/get_all [get]

func (h *Handler) GetAllOrderHandler(ctx *gin.Context) {
	request := pb.GetAllOrderRequest{}

	err := ctx.ShouldBind(&request)
	if err != nil {
		BadRequest(ctx, err)
	}

	if request.Quantity == "" {
		BadRequest(ctx, err)
	}

	if Parse(request.ReservationId) || Parse(request.MenuItemId){
		BadRequest(ctx,fmt.Errorf("id hato"))
		return
	}

	resp, err := h.Order.GetAllOrder(ctx, &request)

	if err != nil {
		InternalServerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

