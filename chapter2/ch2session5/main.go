package main

import (
	"ch2session5/controller"
	"ch2session5/model"
	"ch2session5/repository"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	_ "ch2session5/docs"

	"gorm.io/driver/postgres"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var db *gorm.DB

//@title Books API
//@description Documentation of Books API
//@version 1.0
//@host localhost:8080
//@protocol http

func main() {
	DatabaseInit()
	StartServer()
}

func DatabaseInit() {
	var err error

	db, err = gorm.Open(postgres.Open("host=localhost port=5432 user=postgres password=postgres dbname=gormdb sslmode=disable"), &gorm.Config{})

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

	db.AutoMigrate(model.Author{}, model.Book{})
}

func StartServer() *gin.Engine {

	newBookRepository := repository.NewBookRepository(db)
	bookController := controller.NewBookController(*newBookRepository)

	router := gin.Default()

	router.POST("/books", bookController.CreateBook)
	router.GET("/books/", bookController.GetAllBook)
	router.GET("/books/:bookID", bookController.GetBookByID)
	router.PUT("/books/:bookID", bookController.UpdateBook)
	router.DELETE("/books/:bookID", bookController.DeleteBook)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")

	return router
}
