package book

import (
	middleware "gin-framework/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/books", middleware.RateLimiterMiddleware(), NewBookController().GetBooks)
	r.GET("/books/:id", middleware.RateLimiterMiddleware(), NewBookController().GetBookByID)
	r.POST("/books", middleware.RateLimiterMiddleware(), NewBookController().CreateBook)
	r.PUT("/books/:id", middleware.RateLimiterMiddleware(), NewBookController().UpdateBook)
	r.DELETE("/books/:id", middleware.RateLimiterMiddleware(), NewBookController().DeleteBook)
}