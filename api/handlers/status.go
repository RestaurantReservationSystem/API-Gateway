package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func OK(gn *gin.Context) {
	gn.JSON(200, gin.H{
		"status":  http.StatusOK,
		"time":    time.Now(),
		"success": true,
	})
	gn.Header("Content-Type", "application/json")

}
func Created(gn *gin.Context) {
	gn.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"time":    time.Now(),
		"success": true,
	})
	gn.Header("Content-Type", "application/json")

}
func InternalServerError(gn *gin.Context, err error) {
	fmt.Println("salom")
	gn.JSON(http.StatusInternalServerError, gin.H{
		"status":  http.StatusInternalServerError,
		"time":    time.Now(),
		"message": err.Error(),
		"success": false,
	})
	gn.Header("Content-Type", "application/json")

}
func BadRequest(gn *gin.Context, err error) {
	gn.JSON(http.StatusBadRequest, gin.H{
		"status":  http.StatusBadRequest,
		"time":    time.Now(),
		"message": err.Error(),
		"success": false,
	})
	gn.Header("Content-Type", "application/json")

}

func Parse(id string) bool {
	_, err := uuid.Parse(id)

	return !(err == nil)
}

func IsLimitValidate(limit string) (int, error) {
	if len(limit) == 0 {
		limit += "0"
	}
	limit1, err := strconv.Atoi(limit)
	if err != nil {
		return 0, err
	}
	return limit1, nil
}
func IsOffsetValidate(offset string) (int, error) {
	if len(offset) == 0 {
		offset += "0"
	}
	offset1, err := strconv.Atoi(offset)
	if err != nil {

		return 0, err
	}
	return offset1, nil
}
