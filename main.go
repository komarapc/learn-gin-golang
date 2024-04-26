package main

import (
	"fmt"
	middleware "gin-framework/middleware"
	"gin-framework/src/book"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

// route end point
func router(r *gin.Engine) {
	r.GET("/", middleware.RateLimiterMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Learn RESTFull API with Gin"})
	})
	book.SetupRouter(r)
}

func init() {
	godotenv.Load()
	gin.SetMode(os.Getenv("GIN_MODE"))
	fmt.Println("port", os.Getenv("PORT"))
}

func main() {
	r := gin.Default()
	router(r)
	r.Run(":" + os.Getenv("PORT"))
}
