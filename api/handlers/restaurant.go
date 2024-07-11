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
// @Success 200 {string} string "Create Successful"
// @Failure 400 {string} string "Error while Creating"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/restaurant/create [post]
func (h *Handler) CreateRestaurantHandler(ctx *gin.Context) {
	request := pb.CreateRestaurantRequest{}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		BadRequest(ctx, err)
		return
	}

	if request.Name == "" || request.Address == "" || request.Description == "" {
		BadRequest(ctx, fmt.Errorf("fild lar to'liq toldirilmadi"))
		return
	}
	if len(request.PhoneNumber) == 16 {
		tel := strings.Split(request.PhoneNumber, "-")
		for _, v := range tel {
			_, err = strconv.Atoi(v)
			if err != nil {
				BadRequest(ctx, err)
				return
			}
		}
	}

	_, err = h.ReservationService.CreateRestaurant(ctx, &request)
	if err != nil {
		InternalServerError(ctx, err)
		return
	}

	Created(ctx)
}

// UpdateRestaurantHandler handles the updating of an existing restaurant.
// @Summary Update Restaurant
// @Description Update an existing restaurant
// @Tags Restaurant
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "Restaurant ID"
// @Param Update body genproto.UpdateRestaurantRequest true "Update Restaurant"
// @Success 200 {string} string "Update Successful"
// @Failure 400 {string} string "Error while Updating"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/restaurant/update/{id} [put]
func (h *Handler) UpdateRestaurantHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	_, err := uuid.Parse(id)
	if err != nil {
		BadRequest(ctx, err)
		return
	}

	request := pb.UpdateRestaurantRequest{}
	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		BadRequest(ctx, err)
		return
	}
	byId := pb.IdRequest{
		Id: request.Id,
	}
	_, err = h.ReservationService.GetByIdRestaurant(ctx, &byId)
	if err != nil {
		BadRequest(ctx, fmt.Errorf("error is bu id yoq"))
		return
	}
	request.Id = id
	_, err = h.ReservationService.UpdateRestaurant(ctx, &request)
	if err != nil {
		fmt.Println("+++++++++", err)
		InternalServerError(ctx, err)
		return
	}

	OK(ctx)
}

// DeleteRestaurantHandler handles the deletion of a restaurant.
// @Summary Delete Restaurant
// @Description Delete an existing restaurant
// @Tags Restaurant
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "Restaurant ID"
// @Success 200 {string} string "Delete Successful"
// @Failure 400 {string} string "Error while Deleting"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/restaurant/delete/{id} [delete]
func (h *Handler) DeleteRestaurantHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	_, err := uuid.Parse(id)
	if err != nil {
		BadRequest(ctx, err)
		return
	}

	_, err = h.ReservationService.GetByIdRestaurant(ctx, &pb.IdRequest{Id: id})
	if err != nil {
		BadRequest(ctx, fmt.Errorf("bu id oldin ochirilgan"))
		return
	}

	_, err = h.ReservationService.DeleteRestaurant(ctx, &pb.IdRequest{Id: id})
	if err != nil {
		InternalServerError(ctx, err)
		return
	}

	OK(ctx)
}

// GetByIdRestaurantHandler retrieves a restaurant by its ID.
// @Summary Get Restaurant by ID
// @Description Retrieve a restaurant by its ID
// @Tags Restaurant
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "Restaurant ID"
// @Success 200 {object} genproto.RestaurantResponse
// @Failure 400 {string} string "Error while Retrieving"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/restaurant/get_by_id/{id} [get]
func (h *Handler) GetByIdRestaurantHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	_, err := uuid.Parse(id)
	if err != nil {
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

func validateLimitOffset(value string, defaultValue int) (int, error) {
	if value == "" {
		return defaultValue, nil
	}

	return IsLimitOffsetValidate(value)
}

// GetAllRestaurantsHandler retrieves a list of restaurants with optional filtering and pagination.
// @Summary Get All Restaurants
// @Description Retrieve a list of restaurants with optional filtering and pagination
// @Tags Restaurant
// @Accept json
// @Produce json
// @Param name query string false "Filter by restaurant name"
// @Param phone_number query string false "Filter by restaurant phone number"
// @Param address query string false "Filter by restaurant address"
// @Param description query string false "Filter by restaurant description"
// @Param limit query int false "Number of items to return"
// @Param offset query int false "Offset for pagination"
// @Success 200 {object} genproto.RestaurantsResponse
// @Failure 400 {string} string "Error while Retrieving"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/restaurant/get_all [get]
func (h *Handler) GetAllRestaurantsHandler(ctx *gin.Context) {
	request := pb.GetAllRestaurantRequest{
		Name:        ctx.Query("name"),
		PhoneNumber: ctx.Query("phone_number"),
		Address:     ctx.Query("address"),
		Description: ctx.Query("description"),
	}

	limit, err := validateLimitOffset(ctx.Query("limit"), 10) // Default limit is 10
	if err != nil {
		BadRequest(ctx, err)
		return
	}

	offset, err := validateLimitOffset(ctx.Query("offset"), 0) // Default offset is 0
	if err != nil {
		BadRequest(ctx, err)
		return
	}

	request.LimitOffset = &pb.Filter{
		Offset: int64(offset),
		Limit:  int64(limit),
	}

	resp, err := h.ReservationService.GetAllRestaurants(ctx, &request)
	if err != nil {
		InternalServerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
