package handlers

import (
	pb "api_get_way/genproto"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateMenuHandler handles the creation of a new menu item.
// @Summary Create Menu
// @Description Create a new menu item
// @Tags Menu
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param Create body genproto.CreateMenuRequest true "Create Menu"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/menu/create [post]

func (h *Handler) CreateMenuHandler(ctx *gin.Context) {
	// Initialize a protobuf request structure
	request := pb.CreateMenuRequest{}

	// Bind JSON request body to protobuf structure
	if err := ctx.ShouldBindJSON(&request); err != nil {
		BadRequest(ctx, err)
		h.log.Error("error")

		if Parse(request.RestaurantId) {
			BadRequest(ctx, fmt.Errorf("id hato"))
			h.log.Error("error")
			return
		}

		// Perform additional validation if needed
		if request.Description == "" || request.Name == "" {
			BadRequest(ctx, fmt.Errorf("malumot toliq emas"))
			h.log.Error("error")
			return
		}

		if Parse(request.RestaurantId) {
			BadRequest(ctx, fmt.Errorf("id hato"))
			h.log.Error("error")
			return
		}
		if request.Price <= 0 {
			BadRequest(ctx, fmt.Errorf("hatolik price"))
			h.log.Error("error")
			return
		}
		_, err = h.ReservationService.CreateMenu(ctx, &request)
		if err != nil {
			InternalServerError(ctx, err)
			h.log.Error("error")
			return
		}

		// Call the service method to create the menu item
		_, err := h.ReservationService.CreateMenu(ctx, &request)
		if err != nil {
			InternalServerError(ctx, err)
			h.log.Error("error")
			return
		}

		// Respond with success message
		Created(ctx)
	}

	// UpdateMenuHandler handles the update of a menu item.
	// @Summary Update Menu
	// @Description Update an existing menu item
	// @Tags Menu
	// @Accept json
	// @Security BearerAuth
	// @Produce json
	// @Param id path string true "Menu ID"
	// @Param Update body genproto.UpdateMenuRequest true "Update Menu"
	// @Success 200 {object} string
	// @Failure 400 {object} string
	// @Failure 500 {object} string
	// @Router /api/menu/update/{id} [put]
}
func (h *Handler) UpdateMenuHandler(ctx *gin.Context) {
	request := pb.UpdateMenuRequest{}

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		BadRequest(ctx, err)
		h.log.Error("error")
		return
	}
	request.Id = ctx.Param("id")

	if Parse(request.RestaurantId) || Parse(request.Id) {
		BadRequest(ctx, fmt.Errorf("id hato"))
		h.log.Error("error")
		return
	}

	if request.Name == "" || request.Description == "" {
		BadRequest(ctx, fmt.Errorf("malumot toliq emas"))
		h.log.Error("error")
		return
	}

	_, err = h.ReservationService.UpdateMenu(ctx, &request)

	if err != nil {
		fmt.Println("++++++++++++++")

		InternalServerError(ctx, err)
		h.log.Error("error")
		return
	}

	OK(ctx)
}

// DeleteMenuHandler handles the deletion of a menu item.
// @Summary Delete Menu
// @Description Delete an existing menu item
// @Tags Menu
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "Menu ID"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/menu/delete/{id} [delete]

func (h *Handler) DeleteMenuHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if Parse(id) {
		h.log.Error("error")
		return
	}
	_, err := h.ReservationService.GetByIdReservation(ctx, &pb.IdRequest{Id: id})
	if err != nil {
		BadRequest(ctx, err)
		h.log.Error("error")
		return
	}

	_, err = h.ReservationService.DeleteMenu(ctx, &pb.IdRequest{Id: id})

	if err != nil {
		fmt.Println("++++++++++++", err)
		InternalServerError(ctx, err)
		h.log.Error("error")
		return
	}
	h.log.Info("ishladi")
	OK(ctx)
}

// GetByIdMenuHandler handles the request to fetch a menu item by its ID.
// @Summary Get Menu by ID
// @Description Get a menu item by its ID
// @Tags Menu
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "Menu ID"
// @Success 200 {object} genproto.MenuResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/menu/get_id/{id} [get]

func (h *Handler) GetByIdMenuHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if Parse(id) {
		h.log.Error("error")
		return
	}
	resp, err := h.ReservationService.GetByIdMenu(ctx, &pb.IdRequest{Id: id})

	if err != nil {
		InternalServerError(ctx, err)
		h.log.Error("error")
		return
	}
	h.log.Info("ishladi")
	ctx.JSON(http.StatusOK, resp)
}

// GetAllMenuHandler handles the request to fetch all menu items.
// @Summary Get All Menus
// @Description Get all menu items
// @Tags Menu
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param request query genproto.GetAllMenuRequest true "Get All Menus"
// @Success 200 {object} genproto.MenusResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/menu/get_all [get]

func (h *Handler) GetAllMenuHandler(ctx *gin.Context) {
	request := pb.GetAllMenuRequest{}

	err := ctx.ShouldBind(&request)

	if err != nil {
		BadRequest(ctx, err)
		h.log.Error("error")
		return
	}

	if Parse(request.RestaurantId) {
		BadRequest(ctx, fmt.Errorf("id hato"))
		h.log.Error("error")
		return
	}
	if request.Name == "" || request.Description == "" {
		BadRequest(ctx, fmt.Errorf("malumot toliq emas"))
		h.log.Error("error")
		return
	}
	if len(request.RestaurantId) > 0 {

		if Parse(request.RestaurantId) {
			BadRequest(ctx, fmt.Errorf("id hato"))
			h.log.Error("error")
			return
		}
	}
	resp, err := h.ReservationService.GetAllMenu(ctx, &request)

	if err != nil {
		InternalServerError(ctx, err)
		h.log.Error("error")
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
