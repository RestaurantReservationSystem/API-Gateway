package handlers

import (
	pb "api_get_way/genproto"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"strings"
)

func isValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func isValidLimit(limit string) (int, error) {
	if limit == "" {
		return 0, nil
	}
	return strconv.Atoi(limit)
}

func isValidOffset(offset string) (int, error) {
	if offset == "" {
		return 0, nil
	}
	return strconv.Atoi(offset)
}

// CreateUser handles the creation of a new user.
// @Summary Create User
// @Description Create a new user
// @Tags User
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param Create body genproto.RegisterUserRequest true "Create"
// @Success 200 {string} string "Create Successful"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/user/register [post]
func (h *Handler) CreateUser(gn *gin.Context) {
	request := pb.RegisterUserRequest{}
	if err := gn.ShouldBindJSON(&request); err != nil {
		BadRequest(gn, err)
		return
	}
	fmt.Println("++++++++++")

	if len(request.UserName) < 4 {
		BadRequest(gn, fmt.Errorf("username is not valid"))
		return
	}
	if len(request.Password) < 7 {
		BadRequest(gn, fmt.Errorf("password is not valid"))
		return
	}
	if len(request.Email) < 7 || !strings.Contains(request.Email, "@gmail.com") {
		BadRequest(gn, fmt.Errorf("email is not valid"))
		return
	}

	_, err := h.UsersService.CreateUser(gn, &request)
	if err != nil {
		fmt.Println("++++++++++", err)
		InternalServerError(gn, err)
		return
	}

	Created(gn)
}

// UpdateUser handles updating a user.
// @Summary Update User
// @Description Update an existing user
// @Tags User
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "User ID"
// @Param Update body genproto.UpdatedUserRequest true "Update"
// @Success 200 {string} string "Update Successful"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/user/update/{id} [put]
func (h *Handler) UpdateUser(gn *gin.Context) {
	id := gn.Param("id")
	if !isValidUUID(id) {
		BadRequest(gn, fmt.Errorf("ID is not valid"))
		return
	}

	request := pb.UpdatedUserRequest{}
	if err := gn.ShouldBind(&request); err != nil {
		BadRequest(gn, err)
		return
	}

	if len(request.UserName) < 4 {
		BadRequest(gn, fmt.Errorf("username is not valid"))
		return
	}
	if len(request.Password) < 7 {
		BadRequest(gn, fmt.Errorf("password is not valid"))
		return
	}
	if len(request.Email) < 7 || !strings.Contains(request.Email, "@gmail.com") {
		BadRequest(gn, fmt.Errorf("email is not valid"))
		return
	}
	request.Id = id
	_, err := h.UsersService.GetByIdUser(gn, &pb.IdRequest{})
	if err != nil {
		BadRequest(gn, fmt.Errorf("bu id database da yoq"))
		return
	}
	_, err = h.UsersService.UpdateUser(gn, &request)
	if err != nil {
		InternalServerError(gn, err)
		return
	}

	OK(gn)
}

// DeleteUser handles deleting a user.
// @Summary Delete User
// @Description Delete an existing user
// @Tags User
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {string} string "Delete Successful"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/user/delete/{id} [delete]
func (h *Handler) DeleteUser(gn *gin.Context) {
	id := gn.Param("id")
	if !isValidUUID(id) {
		BadRequest(gn, fmt.Errorf("ID is not valid"))
		return
	}
	_, err := h.UsersService.GetByIdUser(gn, &pb.IdRequest{})
	if err != nil {
		BadRequest(gn, fmt.Errorf("bu id database da yoq"))
		return
	}

	_, err = h.UsersService.DeleteUser(gn, &pb.IdRequest{Id: id})
	if err != nil {
		InternalServerError(gn, err)
		return
	}

	OK(gn)
}

// GetUserById handles fetching a user by ID.
// @Summary Get User by ID
// @Description Get a user by their ID
// @Tags User
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} genproto.UserResponse "Get User Successful"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/user/get_id/{id} [get]
func (h *Handler) GetUserById(gn *gin.Context) {
	id := gn.Param("id")
	if !isValidUUID(id) {
		BadRequest(gn, fmt.Errorf("ID is not valid"))
		return
	}

	response, err := h.UsersService.GetByIdUser(gn, &pb.IdRequest{Id: id})
	if err != nil {
		InternalServerError(gn, err)
		return
	}

	gn.JSON(http.StatusOK, response)
}

// GetAllUser handles fetching all users with optional filters.
// @Summary Get all users
// @Description Get all users with optional filtering
// @Tags User
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param user_name query string false "User Name"
// @Param password query string false "Password"
// @Param email query string false "Email"
// @Param limit query string false "Limit"
// @Param offset query string false "Offset"
// @Success 200 {object} genproto.GetAllUsers "Get All Users Successful"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/user/get_all [get]
func (h *Handler) GetAllUser(gn *gin.Context) {
	limitStr := gn.Query("limit")
	offsetStr := gn.Query("offset")

	limit, err := isValidLimit(limitStr)
	if err != nil {
		BadRequest(gn, err)
		return
	}

	offset, err := isValidOffset(offsetStr)
	if err != nil {
		BadRequest(gn, err)
		return
	}

	request := pb.GetAllUserRequest{
		UserName: gn.Query("user_name"),
		Password: gn.Query("password"),
		Email:    gn.Query("email"),
		Filter:   &pb.Filter{Limit: int64(limit), Offset: int64(offset)},
	}

	response, err := h.UsersService.GetAllUser(gn, &request)
	if err != nil {
		InternalServerError(gn, err)
		return
	}

	gn.JSON(http.StatusOK, response)
}

// LoginUser handles user login.
// @Summary Login User
// @Description Login a user
// @Tags User
// @Accept json
// @Produce json
// @Param Create body genproto.LoginRequest true "Login"
// @Success 200 {object} string "Login Successful"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/user/login [post]
func (h *Handler) LoginUser(gn *gin.Context) {
	request := pb.LoginRequest{}
	if err := gn.ShouldBindJSON(&request); err != nil {
		BadRequest(gn, err)
		return
	}

	if len(request.Email) < 7 && strings.Contains(request.Email, "@gmail.com") {
		BadRequest(gn, fmt.Errorf("email is not valid"))
		return
	}

	response, err := h.UsersService.LoginUser(gn, &request)
	if err != nil {
		InternalServerError(gn, err)
		return
	}
	accessToken, err := h.UsersService.GenerateToken(gn, response)
	if err != nil {
		InternalServerError(gn, err)
		return
	}

	gn.JSON(http.StatusOK, accessToken)
}
