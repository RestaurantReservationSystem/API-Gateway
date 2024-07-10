package api

import (
	_ "api_get_way/api/docs"
	"api_get_way/api/handlers"
	genproto "api_get_way/genproto"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/files" // Import swaggo files handler
	files "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger" // Import gin-swagger middleware
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
)

// @tite APi service
// @version 1.0
// @description APi service
// @host localhost:8080
// @BasePath /
func RouterApi(con1 *grpc.ClientConn, con2 *grpc.ClientConn, con3 *grpc.ClientConn) *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))
	paymentCon := genproto.NewPaymentServiceClient(con1)
	reservationCon := genproto.NewReservationServiceClient(con2)
	userCon := genproto.NewUserServiceClient(con3)
	h := handlers.NewHandler(paymentCon, reservationCon, userCon)
	restaurant := router.Group("/api/restaurant")
	{
		restaurant.POST("/create", h.CreateRestaurantHandler)
		restaurant.GET("/get_all", h.GetAllRestaurantsHandler)
		restaurant.GET("/get_by_id/:id", h.GetByIdRestaurantHandler)
		restaurant.PUT("/update/:id")
		restaurant.DELETE("/delete/:id", h.DeleteRestaurantHandler)

	}

	restaurant := router.Group("/api/restaurant")
	{
		restaurant.POST("create", h.CreateRestaurantHandler)
		restaurant.GET("/get_all", h.GetAllReservationHandler)
		restaurant.GET("/get_by_id/:id", h.GetByIdReservationHandler)
		restaurant.PUT("update/:id")
		restaurant.PUT("delete/:id", h.DeleteRestaurantHandler)
	}

	reservation := router.Group("/api/reservation")
	{
		reservation.POST("/create", h.CreateReservationHandler)
		reservation.GET("/get_all", h.GetAllReservationHandler)
		reservation.GET("/get_id/:id", h.GetByIdReservationHandler)
		reservation.PUT("/update/:id", h.UpdateReservationHandler)
		reservation.DELETE("/delete/:id", h.DeleteReservationHandler)
	}

	menu := router.Group("/api/menu")
	{
		menu.POST("/create", h.CreateMenuHandler)
		menu.GET("/get_all", h.GetAllMenuHandler)
		menu.GET("/get_id/:id", h.GetByIdMenuHandler)
		menu.PUT("/update/:id", h.UpdateMenuHandler)
		menu.DELETE("/delete/:id", h.DeleteMenuHandler)
		menu.DELETE("/delete/:id", h.DeletePaymentHandler)
	}

	order := router.Group("/api/order")
	{
		order.POST("/create", h.CreateOrderHandler)
		order.GET("/get_all", h.GetAllOrderHandler)
		order.GET("/get_id/:id", h.GetByIdOrderHandler)
		order.PUT("/update/:id", h.UpdateOrderHandler)
		order.DELETE("/delete/:id", h.DeleteOrderHandler)
	}

	payment := router.Group("/api/payment")
	{
		payment.POST("/create", h.CreatePaymentHandler)
		payment.GET("/get_id/:id", h.Get)
		payment.PUT("/update/:id", h.UpdatePaymentHandler)
		payment.DELETE("/delete/:id", h.DeletePaymentHandler)
		payment.GET("/get_all", h.GetAllPaymentHandler)
	}

}
