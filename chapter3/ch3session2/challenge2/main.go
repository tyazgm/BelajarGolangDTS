package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"challenge2/controller"
	"challenge2/middleware"
	"challenge2/model"
	"challenge2/repository"
	"challenge2/service"
)

var db *gorm.DB

func main() {
	DatabaseInit()
	StartServer()
}

func DatabaseInit() {
	var err error

	db, err = gorm.Open(postgres.Open("host=localhost port=5432 user=postgres password=postgres dbname=ecommerce sslmode=disable"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		panic(err)
	}

	err = sqlDB.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println(db)

	db.AutoMigrate(model.User{}, model.Order{})
}

func StartServer() *gin.Engine {

	orderRepository := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(*orderRepository)
	orderController := controller.NewOrderController(*orderService)

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(*userRepository)
	userController := controller.NewUserController(*userService)

	router := gin.Default()

	router.POST("user/register", userController.Register)
	router.POST("user/login", userController.Login)

	orderGroup := router.Group("/order", middleware.AuthMiddleware)

	orderGroup.POST("/", orderController.CreateOrder)
	orderGroup.GET("/", orderController.GetListOrder)
	orderGroup.GET("/:id", orderController.GetOrder)

	router.Run(":8080")

	return router
}
