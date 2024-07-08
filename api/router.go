package api

import (
	"api_get_way/api/handlers"
	"api_get_way/api/middleware"

	"github.com/gin-gonic/gin"
)

func RouterApi(hand *handlers.Handler) {
	router := gin.Default()
	router.Use(middleware.MiddleWare())

	restaurant := router.Group("/api/restaurant")
	{
		restaurant.POST("create")
		restaurant.GET("/get_all")
		restaurant.GET("/get_by_id/:id")
		restaurant.PUT("update/:id")
		restaurant.PUT("delete/:id")
	}

	reservation := router.Group("/api/reservation")
	{
		reservation.POST("/create")
		reservation.GET("/get_all")
		reservation.GET("get_id/:id")
		reservation.PUT("/update/:id")
		reservation.DELETE("/delete/:id")
	}

	menu := router.Group("/api/menu")
	{
		menu.POST("/create")
		menu.GET("/get_all")
		menu.GET("/get_id/:id")
		menu.POST("/update/:id")
		menu.DELETE("/delete/:id")
	}

	payment := router.Group("/api/payment")
	{
		payment.POST("/create")
		payment.GET("/get_id/:id")
		payment.PUT("update/:id")
	}
}
