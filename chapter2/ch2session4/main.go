package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	// _ "github.com/lib/pq"
	"gorm.io/driver/postgres"
)

type Book struct {
	BookID   string `gorm:"primary key" json:"book_id"`
	Title    string `gorm:"not null;unique;varchar(255)" json:"title"`
	Price    int    `gorm:"not null" json:"price"`
	AuthorID int    `json:"author_id"`
}

type Author struct {
	ID    int    `gorm:"primary key" json:"author_id"`
	Name  string `gorm:"not null;varchar(255)" json:"author_name"`
	Books []Book `json:"author_books"`
}

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

	db.AutoMigrate(Author{}, Book{})
}

func StartServer() *gin.Engine {

	router := gin.Default()

	router.POST("/books", CreateBook)
	router.PUT("/books/:bookID", UpdateBook)
	router.GET("/books/", GetAllBook)
	router.GET("/books/:bookID", GetBookByID)
	router.DELETE("/books/:bookID", DeleteBook)

	router.Run(":8080")

	return router
}

func CreateBook(ctx *gin.Context) {
	var newBook Book

	err := ctx.ShouldBindJSON(&newBook)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	tx := db.Create(&newBook)

	if tx.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": tx.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, newBook)
}

func UpdateBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	var bookUpdate Book

	err := ctx.ShouldBindJSON(&bookUpdate)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	tx := db.Clauses(clause.Returning{
		Columns: []clause.Column{
			{
				Name: "book_id",
			},
		}}).
		Where("book_id = ?", bookID).
		Updates(&bookUpdate)

	if tx.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": tx.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":       fmt.Sprintf("book with id %v has been succesfully uptaded", bookID),
		"bookDatas now": bookUpdate,
	})
}

func GetAllBook(ctx *gin.Context) {
	tx := db.Find(&Book{})

	if tx.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": tx.Error.Error(),
		})
	}

	rows, err := tx.Rows()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	books := make([]Book, 0)

	for rows.Next() {
		var book Book
		err = rows.Scan(&book.BookID, &book.Title, &book.Price, &book.AuthorID)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		books = append(books, book)
	}

	ctx.JSON(http.StatusOK, books)
}

func GetBookByID(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	var bookData Book

	tx := db.Find(&bookData, bookID)

	if tx.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": tx.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, bookData)
}

func DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")

	var deletedBook Book

	tx := db.Clauses(clause.Returning{}).Delete(&deletedBook, bookID)

	if tx.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": tx.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":      fmt.Sprintf("book with id %v has been succesfully removed", bookID),
		"deleted book": deletedBook,
	})
}
