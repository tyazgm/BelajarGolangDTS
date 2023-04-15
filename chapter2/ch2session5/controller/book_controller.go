package controller

import (
	"ch2session5/model"
	"ch2session5/repository"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	bookRepository repository.BookRepository
}

func NewBookController(bookRepository repository.BookRepository) *BookController {
	return &BookController{
		bookRepository: bookRepository,
	}
}

// CreateBook godoc
// @summary add book
// @description add book to the database
// @tags Books
// @produce json
// @accept json
// @param data body model.Book true "data is mandatory"
// @succes 200 {object} model.Book
// @router /books [POST]
func (bc *BookController) CreateBook(ctx *gin.Context) {
	var newBook model.Book

	err := ctx.ShouldBindJSON(&newBook)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	savedBook, err := bc.bookRepository.Save(newBook)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, savedBook)
}

// GetAllBook godoc
// @summary get all book
// @description get all book at database
// @tags Books
// @produce json
// @succes 200 {array} model.Book
// @router /books [GET]
func (bc *BookController) GetAllBook(ctx *gin.Context) {

	books, err := bc.bookRepository.Get()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, books)
}

// GetBookByID godoc
// @summary get book by ID
// @description get selected book from database by ID
// @tags Books
// @param bookID path string true "bookID you want to retrieve"
// @produce json
// @succes 200 {array} model.Book
// @router /books/{bookID} [GET]
func (bc *BookController) GetBookByID(ctx *gin.Context) {
	bookID := ctx.Param("bookID")

	bookData, err := bc.bookRepository.Get(bookID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, bookData)
}

// UpdateBook godoc
// @summary update book
// @description update book choosed by ID
// @tags Books
// @param bookID path string true "bookID you want to update"
// @accept json
// @param data body model.Book true "data is mandatory"
// @produce json
// @succes 200 {object} model.Book
// @router /books/{bookID} [PUT]
func (bc *BookController) UpdateBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	var bookUpdate model.Book

	err := ctx.ShouldBindJSON(&bookUpdate)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	updatedBook, err := bc.bookRepository.Update(bookUpdate, bookID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":       fmt.Sprintf("book with id %v has been succesfully uptaded", bookID),
		"bookDatas now": updatedBook,
	})
}

// DeleteBook godoc
// @summary delete book
// @description delete book by given ID
// @tags Books
// @param bookID path string true "bookID you want to delete"
// @produce json
// @succes 200 {object} model.Book
// @router /books/{bookID} [DELETE]
func (bc *BookController) DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")

	deletedBook, err := bc.bookRepository.Delete(bookID)

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
