package main

// create hello world app

import (
	"gin-framework/src/book"

	middleware "gin-framework/middleware"

	"github.com/gin-gonic/gin"
)

// route end point
func router(r *gin.Engine)  {
	r.GET("/", middleware.RateLimiterMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{ "message": "Hello World"})
	})
	book.SetupRouter(r)
}

func main() {
	port:="8000"
	r:= gin.Default()
	router(r)
	r.Run(":"+port)
}