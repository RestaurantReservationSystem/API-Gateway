package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func OK(gn *gin.Context, err error) {
	gn.JSON(200, gin.H{
		"status":  http.StatusOK,
		"time":    time.Now(),
		"message": err.Error(),
		"success": true,
	})
	gn.Header("Content-Type", "application/json")

}
func Created(gn *gin.Context, err error) {
	gn.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"time":    time.Now(),
		"message": err,
		"success": true,
	})
	gn.Header("Content-Type", "application/json")

}
func InternalServerError(gn *gin.Context, err error) {
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

	return !(err==nil)
}
