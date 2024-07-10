package handlers

import (
	pb "api_get_way/genproto"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateRestaurantHandler handles the creation of a new restaurant.
// @Summary Create Restaurant
// @Description Create a new restaurant
// @Tags Restaurant
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param Create body genproto.CreateRestaurantRequest true "Create Restaurant"
// @Success 200 {object} genproto.RestaurantResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/restaurant/create [post]
func (h *Handler) CreateRestaurantHandler(ctx *gin.Context) {
	request := pb.CreateRestaurantRequest{}

	if err := ctx.ShouldBind(&request); err != nil {
		BadRequest(ctx, err)
		return
	}

	if request.Name == "" || request.Address == "" || request.Description == "" {
		BadRequest(ctx, fmt.Errorf("fields are incomplete"))
		return
	}

	if len(request.PhoneNumber) == 16 {
		tel := strings.Split(request.PhoneNumber, "-")
		for _, v := range tel {
			if _, err := strconv.Atoi(string(v)); err != nil {
				BadRequest(ctx, err)
				return
			}
		}
	}

	resp, err := h.ReservationService.CreateRestaurant(ctx, &request)
	if err != nil {
		InternalServerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// UpdateRestaurantHandler handles the update of a restaurant.
// @Summary Update Restaurant
// @Description Update an existing restaurant
// @Tags Restaurant
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "Restaurant ID"
// @Param Update body genproto.UpdateRestaurantRequest true "Update Restaurant"
// @Success 200 {object} genproto.RestaurantResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/restaurant/update/{id} [put]
func (h *Handler) UpdateRestaurantHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		BadRequest(ctx, err)
		return
	}

	request := pb.UpdateRestaurantRequest{}
	if err := ctx.ShouldBind(&request); err != nil {
		BadRequest(ctx, err)
		return
	}

	resp, err := h.ReservationService.UpdateRestaurant(ctx, &pb.UpdateRestaurantRequest{
		Id:          id,
		Name:        request.Name,
		Address:     request.Address,
		Description: request.Description,
		PhoneNumber: request.PhoneNumber,
	})
	if err != nil {
		InternalServerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// DeleteRestaurantHandler handles the deletion of a restaurant.
// @Summary Delete Restaurant
// @Description Delete an existing restaurant
// @Tags Restaurant
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "Restaurant ID"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/restaurant/delete/{id} [delete]
func (h *Handler) DeleteRestaurantHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		BadRequest(ctx, err)
		return
	}

	if _, err := h.ReservationService.DeleteRestaurant(ctx, &pb.IdRequest{Id: id}); err != nil {
		InternalServerError(ctx, err)
		return
	}

	OK(ctx, nil)
}

// GetByIdRestaurantHandler handles fetching a restaurant by its ID.
// @Summary Get Restaurant by ID
// @Description Get a restaurant by its ID
// @Tags Restaurant
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "Restaurant ID"
// @Success 200 {object} genproto.RestaurantResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/restaurant/get-by-id/{id} [get]
func (h *Handler) GetByIdRestaurantHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		BadRequest(ctx, err)
		return
	}

	resp, err := h.ReservationService.GetByIdRestaurant(ctx, &pb.IdRequest{Id: id})
	if err != nil {
		InternalServerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// GetAllRestaurantsHandler handles the retrieval of all restaurants.
// @Summary Get All Restaurants
// @Description Retrieves a list of all restaurants
// @Tags Restaurant
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param Get query genproto.GetAllRestaurantRequest true "Get All Restaurants"
// @Success 200 {object} genproto.RestaurantsResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/restaurant/get-all [get]
func (h *Handler) GetAllRestaurantsHandler(ctx *gin.Context) {
	request := pb.GetAllRestaurantRequest{}
	if err := ctx.ShouldBind(&request); err != nil {
		BadRequest(ctx, err)
		return
	}

	if request.Address == "" || request.Description == "" || request.Name == "" {
		BadRequest(ctx, fmt.Errorf("fields are incomplete"))
		return
	}

	if len(request.PhoneNumber) == 16 {
		tel := strings.Split(request.PhoneNumber, "-")
		for _, v := range tel {
			if _, err := strconv.Atoi(string(v)); err != nil {
				BadRequest(ctx, err)
				return
			}
		}
	}

	resp, err := h.ReservationService.GetAllRestaurants(ctx, &request)
	if err != nil {
		InternalServerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
