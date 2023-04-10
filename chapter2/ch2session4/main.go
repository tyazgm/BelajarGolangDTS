package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

type Book struct {
	BookID string
	Title  string
	Author string
	Price  int
}

var db *sql.DB

func main() {
	StartServer()
}

func StartServer() *gin.Engine {
	var err error

	db, err = sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=bookDB sslmode=disable")

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println(db)

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
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	query := "insert into book values($1, $2, $3, $4) returning *"

	row := db.QueryRow(query, newBook.BookID, newBook.Title, newBook.Author, newBook.Price)

	err = row.Scan(&newBook.BookID, &newBook.Title, &newBook.Author, &newBook.Price)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
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

	query := "update book set title = $1, author = $2, price = $3 where bookid = $4 returning *"

	rows, err := db.Query(query, &bookUpdate.Title, &bookUpdate.Author, &bookUpdate.Price, &bookUpdate.BookID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	books := make([]Book, 0)

	for rows.Next() {
		var book Book
		err = rows.Scan(&book.BookID, &book.Title, &book.Author, &book.Price)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		books = append(books, book)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":       fmt.Sprintf("book with id %v has been succesfully uptaded", bookID),
		"bookDatas now": books,
	})
}

func GetAllBook(ctx *gin.Context) {
	query := "select * from book"

	rows, err := db.Query(query)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	books := make([]Book, 0)

	for rows.Next() {
		var book Book
		err = rows.Scan(&book.BookID, &book.Title, &book.Author, &book.Price)

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

	query := "select * from book where bookid=$1 returning *"

	row := db.QueryRow(query, bookID)

	err := row.Scan(&bookData.BookID, &bookData.Title, &bookData.Author, &bookData.Price)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, bookData)
}

func DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")

	var deletedBook Book

	query := "delete from book where bookid=$1 returning *"

	row := db.QueryRow(query, bookID)

	err := row.Scan(&deletedBook.BookID, &deletedBook.Title, &deletedBook.Author, &deletedBook.Price)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":      fmt.Sprintf("book with id %v has been succesfully removed", bookID),
		"deleted book": deletedBook,
	})
}
