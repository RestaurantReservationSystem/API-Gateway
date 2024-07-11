package handlers

import (
	pb "api_get_way/genproto"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateOrderHandler handles the creation of a new order.
// @Summary Create Order
// @Description Create a new order
// @Tags Order
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param Create body genproto.CreateOrderRequest true "Create Order"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/order/create [post]
func (h *Handler) CreateOrderHandler(ctx *gin.Context) {
	request := pb.CreateOrderRequest{}

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		BadRequest(ctx, err)
		h.log.Error("error")
		return
	}

	if request.Quantity == "" {
		BadRequest(ctx, fmt.Errorf("quantity is required"))
		h.log.Error("error")
		return
	}

	if Parse(request.ReservationId) || Parse(request.MenuItemId) {
		BadRequest(ctx, fmt.Errorf("id hato"))
		h.log.Error("error")
		return
	}
	if Parse(request.ReservationId) || Parse(request.MenuItemId) {
		BadRequest(ctx, fmt.Errorf("id hato"))
		h.log.Error("error")
		return
	}
	_, err = h.ReservationService.GetByIdMenu(ctx, &pb.IdRequest{Id: request.MenuItemId})
	if err != nil {
		BadRequest(ctx, fmt.Errorf("menu id mavjud emas bazada"))
		return
	}
	_, err = h.ReservationService.GetByIdReservation(ctx, &pb.IdRequest{Id: request.ReservationId})
	if err != nil {
		BadRequest(ctx, fmt.Errorf("reservation id mavjud emas bazada"))
		return
	}

	_, err = h.ReservationService.CreateOrder(ctx, &request)
	if err != nil {
		InternalServerError(ctx, err)
		h.log.Error("error")
		return
	}
	h.log.Info("ishladi")
	Created(ctx)
}

// UpdateOrderHandler handles the update of an order.
// @Summary Update Order
// @Description Update an existing order
// @Tags Order
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "Order ID"
// @Param Update body genproto.UpdateOrderRequest true "Update Order"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/order/update/{id} [put]
func (h *Handler) UpdateOrderHandler(ctx *gin.Context) {
	request := pb.UpdateOrderRequest{}

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		BadRequest(ctx, err)
		h.log.Error("error")
		return
	}

	request.Id = ctx.Param("id")

	if request.Quantity == "" {
		BadRequest(ctx, fmt.Errorf("quantity is required"))
		h.log.Error("error")
		return
	}

	if Parse(request.ReservationId) || Parse(request.MenuItemId) || Parse(request.Id) {
		BadRequest(ctx, fmt.Errorf("id hato"))
		h.log.Error("error")
		return
	}
	_, err = h.ReservationService.GetByIdOrder(ctx, &pb.IdRequest{Id: request.Id})
	if err != nil {
		BadRequest(ctx, fmt.Errorf("error -> order id mavjud emas"))
		return
	}
	_, err = h.ReservationService.GetByIdMenu(ctx, &pb.IdRequest{Id: request.MenuItemId})
	if err != nil {
		BadRequest(ctx, fmt.Errorf("menu id mavjud emas bazada"))
		return
	}
	_, err = h.ReservationService.GetByIdReservation(ctx, &pb.IdRequest{Id: request.ReservationId})
	if err != nil {
		BadRequest(ctx, fmt.Errorf("reservation id mavjud emas bazada"))
		return
	}

	_, err = h.ReservationService.UpdateOrder(ctx, &request)
	if err != nil {
		InternalServerError(ctx, err)
		h.log.Error("error")
		return
	}
	h.log.Info("ishladi")
	OK(ctx)
}

// DeleteOrderHandler handles the deletion of an order.
// @Summary Delete Order
// @Description Delete an existing order
// @Tags Order
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/order/delete/{id} [delete]
func (h *Handler) DeleteOrderHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if Parse(id) {
		BadRequest(ctx, fmt.Errorf("id hato"))
		h.log.Error("error")
		return
	}
	_, err := h.ReservationService.GetByIdOrder(ctx, &pb.IdRequest{})
	if err != nil {
		BadRequest(ctx, fmt.Errorf("error -> order id mavjud emas"))
		return
	}

	_, err = h.ReservationService.DeleteOrder(ctx, &pb.IdRequest{Id: id})
	if err != nil {
		InternalServerError(ctx, err)
		h.log.Error("error")
		return
	}
	h.log.Info("ishladi")
	OK(ctx)
}

// GetByIdOrderHandler handles fetching an order by its ID.
// @Summary Get Order by ID
// @Description Get an order by its ID
// @Tags Order
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} genproto.OrderResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/order/get_id/{id} [get]
func (h *Handler) GetByIdOrderHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if Parse(id) {
		BadRequest(ctx, fmt.Errorf("id hato"))
		h.log.Error("error")
		return
	}

	resp, err := h.ReservationService.GetByIdOrder(ctx, &pb.IdRequest{Id: id})

	if err != nil {
		InternalServerError(ctx, err)
		h.log.Error("error")
		return
	}
	h.log.Info("ishladi")
	ctx.JSON(http.StatusOK, resp)
}

// GetAllOrderHandler retrieves orders with optional filtering and pagination.
// @Summary Get All Order
// @Description Retrieve orders with optional filtering and pagination.
// @Tags Order
// @Accept json
// @Produce json
// @Param quantity query string false "Filter by order item quantity"
// @Param menu_item_id query string false "Filter by menu item ID"
// @Param reservation_id query string false "Filter by reservation ID"
// @Param limit query int false "Number of items to return"
// @Param offset query int false "Offset for pagination"
// @Success 200 {object} genproto.OrdersResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/order/get_all [get]
func (h *Handler) GetAllOrderHandler(ctx *gin.Context) {
	request := pb.GetAllOrderRequest{}

	limit := ctx.Query("limit")
	limit1, err := IsLimitOffsetValidate(limit)
	if err != nil {
		BadRequest(ctx, err)
		h.log.Error("error")
		return
	}
	offset := ctx.Query("offset")
	offset1, err := IsLimitOffsetValidate(offset)
	if err != nil {
		BadRequest(ctx, err)
		return
	}
	request.LimitOffset = &pb.Filter{
		Limit:  int64(limit1),
		Offset: int64(offset1),
	}
	request.Quantity = ctx.Query("quantity")
	if request.Quantity != "" {
		if _, err := strconv.ParseFloat(request.Quantity, 32); err != nil {
			BadRequest(ctx, err)
			return
		}
	}
	request.MenuItemId = ctx.Query("menu_item_id")
	if request.MenuItemId != "" {
		if !Parse(request.MenuItemId) {
			BadRequest(ctx, fmt.Errorf("invalid menu_item_id format"))
			return
		}
		_, err := h.ReservationService.GetByIdMenu(ctx, &pb.IdRequest{Id: request.MenuItemId})
		if err != nil {
			BadRequest(ctx, fmt.Errorf("menu id does not exist in the database"))
			return
		}
	}
	request.ReservationId = ctx.Query("reservation_id")
	if request.ReservationId != "" {
		if !Parse(request.ReservationId) {
			BadRequest(ctx, fmt.Errorf("invalid reservation_id format"))
			return
		}
		_, err := h.ReservationService.GetByIdReservation(ctx, &pb.IdRequest{Id: request.ReservationId})
		if err != nil {
			BadRequest(ctx, fmt.Errorf("reservation id does not exist in the database"))
			return
		}
	}

	resp, err := h.ReservationService.GetAllOrder(ctx, &request)
	if err != nil {
		InternalServerError(ctx, err)
		h.log.Error("error")
		return
	}
	h.log.Info("ishladi")

	ctx.JSON(http.StatusOK, resp)
}
