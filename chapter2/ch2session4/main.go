package main

import (
	"ch2session4/controller"
	"ch2session4/model"
	"ch2session4/repository"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

var db *gorm.DB

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

	router.Run(":8080")

	return router
}
