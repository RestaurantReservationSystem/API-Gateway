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

func (h *Handler) CreateRestaurantHandler(ctx *gin.Context) {
	request := pb.CreateRestaurantRequest{}

	err := ctx.ShouldBind(&request)
	if err != nil {
		BadRequest(ctx, err)
	}

	if request.Name == "" || request.Address == "" || request.Description == "" {
		BadRequest(ctx, fmt.Errorf("fild lar to'liq toldirilmadi"))
		return
	}
	if len(request.PhoneNumber) == 16 {
		tel := strings.Split(request.PhoneNumber, "-")

		for _, v := range tel {
			_, err = strconv.Atoi(string(v))
			if err != nil {
				BadRequest(ctx, err)
				return
			}

		}
	}

	_, err = h.Restaran.CreateRestaurant(ctx, &request)
	if err != nil {
		InternalServerError(ctx, err)
		return
	}

	Created(ctx, nil)
}

func (h *Handler) UpdateRestaurant() {

}

func (h *Handler) DeleteRestaurantHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	_, err := uuid.Parse(id)

	if err != nil {
		BadRequest(ctx, err)
		return
	}
	_, err = h.Restaran.DeleteRestaurant(ctx, &pb.IdRequest{Id: id})

	if err != nil {
		InternalServerError(ctx, err)
	}

	OK(ctx, nil)
}

func (h *Handler) GetByIdRestaurantHandler(ctx *gin.Context) {

	id := ctx.Param("id")

	_, err := uuid.Parse(id)

	if err != nil {
		BadRequest(ctx, err)
	}

	resp,err:=h.Restaran.GetByIdRestaurant(ctx,&pb.IdRequest{Id: id,})

	if err != nil {
		InternalServerError(ctx,err)
	}

	ctx.JSON(http.StatusOK,resp)
}

func (h *Handler) GetAllRestaurantsHandler(ctx *gin.Context) {
	
	request:=pb.GetAllRestaurantRequest{}

	err:=ctx.ShouldBind(&request)

	if err != nil {
		BadRequest(ctx,err)
	}

	if request.Address=="" || request.Description=="" || request.Name==""{
		BadRequest(ctx,fmt.Errorf("malumot tioliq emas"))
	}

	if len(request.PhoneNumber) == 16 {
		tel := strings.Split(request.PhoneNumber, "-")

		for _, v := range tel {
			_, err = strconv.Atoi(string(v))
			if err != nil {
				BadRequest(ctx, err)
				return
			}

		}
	}

	resp,err:=h.Restaran.GetAllRestaurants(ctx,&request)

	if err != nil {
		InternalServerError(ctx,err)
	}

	ctx.JSON(http.StatusOK,resp)
}
