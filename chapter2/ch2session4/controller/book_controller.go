package controller

import (
	"ch2session4/model"
	"ch2session4/repository"
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
