package book

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type BookController struct {
	BookService BookService
}

func NewBookController() *BookController {
	return &BookController{
		BookService: BookService{},
	}
}

func (c *BookController) GetBooks(ctx *gin.Context) {
	books := c.BookService.GetBooks()
	ctx.JSON(http.StatusOK, books)
}
func (c *BookController) GetBookByID(ctx *gin.Context) {
	id,_ :=  strconv.Atoi(ctx.Param("id"))
	book := c.BookService.GetBookByID(id)
	if book == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	ctx.JSON(http.StatusOK, book)
}

func (c *BookController) CreateBook(ctx *gin.Context) {
	var book Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := c.BookService.CreateBook(book)
	ctx.JSON(http.StatusCreated, result)
}

func (c *BookController) UpdateBook(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var book Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := c.BookService.UpdateBook(id, book)
	if result == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (c *BookController) DeleteBook(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	result := c.BookService.DeleteBook(id)
	if !result {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}