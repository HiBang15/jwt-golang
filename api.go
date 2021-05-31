package main

import (
	"github.com/HiBang15/jwt-golang/handler"
	"github.com/HiBang15/jwt-golang/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/login")
	router.GET("/api/v1/product", middleware.ValidateToken(), handler.GetAll)
	router.Run("localhost:8080")
}
