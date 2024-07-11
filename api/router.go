package api

import (
	_ "api_get_way/api/docs"
	"api_get_way/api/handlers"
	"api_get_way/api/middleware"
	genproto "api_get_way/genproto"

	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/files"
	files "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// RouterApi @title API Service
// @version 1.0
// @description API service
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func RouterApi(con1 *grpc.ClientConn, con2 *grpc.ClientConn, con3 *grpc.ClientConn, logger *zap.Logger) *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

	paymentCon := genproto.NewPaymentServiceClient(con1)
	reservationCon := genproto.NewReservationServiceClient(con2)
	userCon := genproto.NewUserServiceClient(con3)
	h := handlers.NewHandler(paymentCon, reservationCon, userCon, logger)

	user := router.Group("/api/user")
	{
		user.POST("/register", h.CreateUser)
		user.POST("/login", h.LoginUser)
	}

	authRoutes := router.Group("/")
	authRoutes.Use(middleware.MiddleWare())
	{
		user := authRoutes.Group("/api/user")
		{
			user.GET("/get_id/:id", h.GetUserById)
			user.PUT("/update/:id", h.UpdateUser)
			user.DELETE("/delete/:id", h.DeleteUser)
			user.GET("/get_all", h.GetAllUser)
		}

		restaurant := authRoutes.Group("/api/restaurant")
		{
			restaurant.POST("create", h.CreateRestaurantHandler)
			restaurant.GET("/get_all", h.GetAllRestaurantsHandler)
			restaurant.GET("/get_by_id/:id", h.GetByIdRestaurantHandler)
			restaurant.PUT("update/:id", h.UpdateRestaurantHandler)
			restaurant.DELETE("delete/:id", h.DeleteRestaurantHandler)
		}

		reservation := authRoutes.Group("/api/reservation")
		{
			reservation.POST("/create", h.CreateReservationHandler)
			reservation.GET("/get_all", h.GetAllReservationHandler)
			reservation.GET("/get_id/:id", h.GetByIdReservationHandler)
			reservation.PUT("/update/:id", h.UpdateReservationHandler)
			reservation.DELETE("/delete/:id", h.DeleteReservationHandler)
		}

		menu := authRoutes.Group("/api/menu")
		{
			menu.POST("/create", h.CreateMenuHandler)
			menu.GET("/get_all", h.GetAllMenuHandler)
			menu.GET("/get_id/:id", h.GetByIdMenuHandler)
			menu.PUT("/update/:id", h.UpdateMenuHandler)
			menu.DELETE("/delete/:id", h.DeleteMenuHandler)
		}

		order := authRoutes.Group("/api/order")
		{
			order.POST("/create", h.CreateOrderHandler)
			order.GET("/get_all", h.GetAllOrderHandler)
			order.GET("/get_id/:id", h.GetByIdOrderHandler)
			order.PUT("/update/:id", h.UpdateOrderHandler)
			order.DELETE("/delete/:id", h.DeleteOrderHandler)
		}

		payment := authRoutes.Group("/api/payment")
		{
			payment.POST("/create", h.CreatePaymentHandler)
			payment.GET("/get_id/:id", h.GetByIdPaymentHandler)
			payment.PUT("/update/:id", h.UpdatePaymentHandler)
			payment.DELETE("/delete/:id", h.DeletePaymentHandler)
			payment.GET("/get_all", h.GetAllPaymentHandler)
		}
	}

	return router
}
