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

// CreateRestaurantHandler 		handles the creation of a new user
// @Summary 		Create Menu
// @Description 	Create page
// @Tags 			Menu
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param   		Create  body     pb.RegisterUserRequest  true   "Create"
// @Success 		200   {string}      "Create Successful"
// @Failure 		401   {string}   string    "Error while Created"
// @Router 			/api/restaurant/create [post]

func (h *Handler) CreateRestaurantHandler(ctx *gin.Context) {
	request := pb.CreateRestaurantRequest{}

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		h.log.Error("error")
		BadRequest(ctx, err)
		return
	}

	if request.Name == "" || request.Address == "" || request.Description == "" {
		h.log.Error("error")
		BadRequest(ctx, fmt.Errorf("fild lar to'liq toldirilmadi"))
		return
	}
	if len(request.PhoneNumber) == 16 {
		tel := strings.Split(request.PhoneNumber, "-")

		for _, v := range tel {
			_, err = strconv.Atoi(v)
			if err != nil {
				h.log.Error("error")
				BadRequest(ctx, err)
				return
			}

		}
	}

	_, err = h.ReservationService.CreateRestaurant(ctx, &request)
	if err != nil {
		h.log.Error("error")
		InternalServerError(ctx, err)
		h.log.Error("error")		
		return
	}
	h.log.Info("ishladi")
	Created(ctx)
}

// UpdateRestaurant 		handles the creation of a new user
// @Summary 		Update Menu
// @Description 	Update page
// @Tags 			Menu
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param   		Update  body     pb.RegisterUserRequest  true   "Update"
// @Success 		200   {string}      "Update Successful"
// @Failure 		401   {string}   string    "Error while Created"
// @Router 			/api/restaurant/update/:id [put]

func (h *Handler) UpdateRestaurant() {

}

// DeleteRestaurantHandler 		handles the creation of a new user
// @Summary 		Delete Menu
// @Description 	Delete page
// @Tags 			Menu
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param   		Delete  body     pb.RegisterUserRequest  true   "Delete"
// @Success 		200   {string}      "Delete Successful"
// @Failure 		401   {string}   string    "Error while Created"
// @Router 			/api/restaurant/Delete/:id [put]

func (h *Handler) DeleteRestaurantHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Println("+++++++++++++")

	_, err := uuid.Parse(id)

	if err != nil {
		h.log.Error("error")
		BadRequest(ctx, err)
		return
	}
	_, err = h.ReservationService.GetByIdRestaurant(ctx, &pb.IdRequest{Id: id})
	if err != nil {
		h.log.Error("error")
		BadRequest(ctx, fmt.Errorf("bu id oldin ochiriligan"))
	}
	_, err = h.ReservationService.DeleteRestaurant(ctx, &pb.IdRequest{Id: id})

	if err != nil {
		h.log.Error("error")
		InternalServerError(ctx, err)
	}
	h.log.Info("ishladi")
	OK(ctx)
}

// GetByIdRestaurantHandler 		handles the creation of a new user
// @Summary 		GetId Menu
// @Description 	GetId page
// @Tags 			Menu
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param   		GetId  body     pb.RegisterUserRequest  true   "GetId"
// @Success 		200   {string}      "GetId Successful"
// @Failure 		401   {string}   string    "Error while Created"
// @Router 			/api/restaurant/get_by_id/:id [get]

func (h *Handler) GetByIdRestaurantHandler(ctx *gin.Context) {

	id := ctx.Param("id")

	_, err := uuid.Parse(id)

	if err != nil {
		h.log.Error("error")
		BadRequest(ctx, err)
	}

	resp, err := h.ReservationService.GetByIdRestaurant(ctx, &pb.IdRequest{Id: id})

	if err != nil {
		h.log.Error("error")
		InternalServerError(ctx, err)
	}

	ctx.JSON(http.StatusOK, resp)
}

// GetAllRestaurantsHandler 		handles the creation of a new user
// @Summary 		GetAll Menu
// @Description 	GetAll page
// @Tags 			Menu
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param   		GetAll  body     pb.RegisterUserRequest  true   "GetAll"
// @Success 		200   {string}      "GetAll Successful"
// @Failure 		401   {string}   string    "Error while Created"
// @Router 			/api/restaurant/get_all [get]

func (h *Handler) GetAllRestaurantsHandler(ctx *gin.Context) {

	request := pb.GetAllRestaurantRequest{}
	request.Name = ctx.Param("name")
	if request.Name == "" {
		request.Name = ""
	}
	if request.PhoneNumber == "" {
		request.PhoneNumber = ""
	}
	if request.Address == "" {
		request.Address = ""
	}
	if request.Description == "" {
		request.Description = ""
	}
	limit := ctx.Param("limit")
	limit1, err := IsLimitValidate(limit)
	if err != nil {
		h.log.Error("error")
		BadRequest(ctx, err)
		return
	}
	request.LimitOffset.Limit = int64(limit1)
	fmt.Println("+++++++++", limit1)

	offset := ctx.Param("offset")
	offset1, err := IsOffsetValidate(offset)
	if err != nil {
		h.log.Error("error")
		BadRequest(ctx, err)
		return
	}
	request.LimitOffset.Offset = int64(offset1)
	fmt.Println("+++++++++", offset1)

	resp, err := h.ReservationService.GetAllRestaurants(ctx, &request)
	if err != nil {
		h.log.Error("error")
		InternalServerError(ctx, err)
		return
	}
	h.log.Info("ishladi")
	ctx.JSON(http.StatusOK, resp)
}
