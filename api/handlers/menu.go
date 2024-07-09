package handlers

import (
	pb "api_get_way/genproto"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateMenuHandler 		handles the creation of a new user
// @Summary 		Create Menu
// @Description 	Create page
// @Tags 			Menu
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param   		Create  body     pb.RegisterUserRequest  true   "Create"
// @Success 		200   {string}      "Create Successful"
// @Failure 		401   {string}   string    "Error while Created"
// @Router 			/api/menu/create [post]

func (h *Handler) CreateMenuHandler(ctx *gin.Context) {
	request := pb.CreateMenuRequest{}

	err := ctx.ShouldBind(&request)

	if err != nil {
		BadRequest(ctx, err)
	}

	if Parse(request.RestaurantId){
		BadRequest(ctx,fmt.Errorf("id hato"))
		return
	}

	if request.Description == "" || request.Name == "" {
		BadRequest(ctx, fmt.Errorf("malumot toliq emas"))
		return
	}

	_, err = h.Menu.CreateMenu(ctx, &request)

	if err != nil {
		InternalServerError(ctx, err)
		return
	}

	Created(ctx, nil)
}

// UpdateMenuHandler 		handles the creation of a new user
// @Summary 		Update Menu
// @Description 	Update page
// @Tags 			Menu
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param   		Create  body     pb.RegisterUserRequest  true   "Update"
// @Success 		200   {string}      "Create Successful"
// @Failure 		401   {string}   string    "Error while Created"
// @Router 			/api/menu/update/:id [PUT]

func (h *Handler) UpdateMenuHandler(ctx *gin.Context) {

	request := pb.UpdateMenuRequest{}

	err := ctx.ShouldBind(&request)
	if err != nil {
		BadRequest(ctx, err)
		return
	}

	if Parse(request.RestaurantId) || Parse(request.Id){
		BadRequest(ctx,fmt.Errorf("id hato"))
		return
	}

	if request.Name == "" || request.Description == "" {
		BadRequest(ctx, fmt.Errorf("maliumot toliq emas"))
		return
	}

	_, err = h.Menu.UpdateMenu(ctx, &request)

	if err != nil {
		InternalServerError(ctx, err)
		return
	}

	OK(ctx, nil)
}

// DeleteMenuHandler 		handles the creation of a new user
// @Summary 		Delete Menu
// @Description 	Delete page
// @Tags 			Menu
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param   		Create  body     pb.RegisterUserRequest  true   "Delete"
// @Success 		200   {string}      "Delete Successful"
// @Failure 		401   {string}   string    "Error while Delete"
// @Router 			/api/menu/delete/:id [DELETE]

func (h *Handler) DeleteMenuHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if Parse(id){
		BadRequest(ctx,fmt.Errorf("id hato"))
		return
	}

	_, err := h.Menu.DeleteMenu(ctx, &pb.IdRequest{Id: id})

	if err != nil {
		InternalServerError(ctx, err)
		return
	}

	Created(ctx, nil)
}

// GetByIdMenuHandler 		handles the creation of a new user
// @Summary 		GetById Menu
// @Description 	GetById page
// @Tags 			Menu
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param   		Create  body     pb.RegisterUserRequest  true   "GET"
// @Success 		200   {string}      "Get Successful"
// @Failure 		401   {string}   string    "Error while Get"
// @Router 			/api/menu/get_id/:id [DELETE]

func (h *Handler) GetByIdMenuHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if Parse(id){
		BadRequest(ctx,fmt.Errorf("id hato"))
		return
	}
	resp, err := h.Menu.GetByIdMenu(ctx, &pb.IdRequest{Id: id})

	if err != nil {
		InternalServerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)

}

// GetAllMenuHandler 		handles the creation of a new user
// @Summary 		GetAll Menu
// @Description 	GetAll page
// @Tags 			Menu
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param   		Create  body     pb.RegisterUserRequest  true   "GET"
// @Success 		200   {string}      "Get Successful"
// @Failure 		401   {string}   string    "Error while Get"
// @Router 			/api/menu/get_all/ [DELETE]


func (h *Handler) GetAllMenuHandler(ctx *gin.Context) {

	request := pb.GetAllMenuRequest{}

	err := ctx.ShouldBind(&request)

	if err != nil {
		BadRequest(ctx, err)
	}

	if Parse(request.RestaurantId){
		BadRequest(ctx,fmt.Errorf("id hato"))
		return
	}
	if request.Name == "" || request.Description == "" {
		BadRequest(ctx, fmt.Errorf("malumot toliq emas"))
		return
	}

	resp, err := h.Menu.GetAllMenu(ctx, &request)

	if err != nil {
		InternalServerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
