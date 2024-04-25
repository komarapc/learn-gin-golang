package book

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/books", NewBookController().GetBooks)
	r.GET("/books/:id", NewBookController().GetBookByID)
	r.POST("/books", NewBookController().CreateBook)
	r.PUT("/books/:id", NewBookController().UpdateBook)
	r.DELETE("/books/:id", NewBookController().DeleteBook)
}