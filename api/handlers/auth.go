package handlers

import (
	pb "api_get_way/genproto"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterHendler(ctx *gin.Context) {
	request := pb.RegisterUserRequest{}

	err := ctx.ShouldBind(&request)

	if err != nil {
		BadRequest(ctx, err)
		return
	}

	if !strings.Contains(request.Email, "@gmail.com") {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"ERROR": "gmail hato kiritdingiz!",
		})
		return
	}

	if len(request.Password) >= 7 {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"ERROR": "passworda hatilik",
		})
		return
	}

	resp, err := h.UsersService.CreateUser(ctx, &request)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"ERROR": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)

}

func (h *Handler) DeleteUserHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	Parse(ctx, id)

	request := pb.IdRequest{
		Id: id,
	}

	resp, err := h.UsersService.DeleteUser(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error,
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateUserHandler(ctx *gin.Context) {
	request := pb.UpdatedUserRequest{}

	err := ctx.ShouldBind(&request)
	if err != nil {
		BadRequest(ctx, err)
		return
	}

	resp, err := h.UsersService.UpdateUser(ctx, &request)

	if err != nil {
		InternalServerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)

}

func (h *Handler) GetUserByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	Parse(ctx, id)

	request := pb.IdRequest{
		Id: id,
	}

	resp, err := h.UsersService.GetByIdUser(ctx, &request)

	if err != nil {
		InternalServerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) GetAllUserHandler(ctx *gin.Context) {

	request := pb.GetAllUserRequest{}

	err := ctx.ShouldBind(&request)

	if err != nil {
		BadRequest(ctx, err)
		return
	}

	if !strings.Contains(request.Email, "@gmail.com") {
		BadRequest(ctx,fmt.Errorf("email hato"))
		return
	}

	if len(request.Password) >= 7 {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"ERROR": "passworda hatilik",
		})
		return
	}

	resp, err := h.UsersService.GetAllUser(ctx, &request)

	if err != nil {
		InternalServerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)

}
