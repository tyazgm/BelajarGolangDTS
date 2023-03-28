package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	StartServer()
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

type Book struct {
	BookID string
	Title  string
	Author string
	Price  int
}

var BookDatas = []Book{}

func CreateBook(ctx *gin.Context) {
	var newBook Book

	err := ctx.ShouldBindJSON(&newBook)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newBook.BookID = fmt.Sprintf("c%d", len(BookDatas)+1)
	BookDatas = append(BookDatas, newBook)

	ctx.JSON(http.StatusOK, BookDatas)
}

func UpdateBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	found := false
	var updatedBook Book

	err := ctx.ShouldBindJSON(&updatedBook)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, book := range BookDatas {
		if bookID == book.BookID {
			found = true
			BookDatas[i] = updatedBook
			BookDatas[i].BookID = bookID
			break
		}
	}

	if !found {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":       fmt.Sprintf("book with id %v has been succesfully uptaded", bookID),
		"bookDatas now": BookDatas,
	})
}

func GetAllBook(ctx *gin.Context) {
	bookData := BookDatas

	ctx.JSON(http.StatusOK, bookData)
}

func GetBookByID(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	bookData := BookDatas
	found := false

	for i, car := range bookData {
		if car.BookID == bookID {
			found = true
			ctx.JSON(http.StatusOK, bookData[i])
			break
		}
	}

	if !found {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}
}

func DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	found := false
	var bookIdx int

	for i, book := range BookDatas {
		if book.BookID == bookID {
			found = true
			bookIdx = i
			break
		}
	}

	if !found {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}

	copy(BookDatas[bookIdx:], BookDatas[bookIdx+1:])
	BookDatas[len(BookDatas)-1] = Book{}
	BookDatas = BookDatas[:len(BookDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("book with id %v has been succesfully removed", bookID),
	})
}
