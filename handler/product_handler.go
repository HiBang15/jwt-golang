package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Product struct {
	Id int
	Name string
}
func GetAll(context *gin.Context) {
	context.AbortWithStatusJSON(http.StatusOK, []Product{
		{
			Id: 1,
			Name: "Product 1",
		},
		{
			Id: 2,
			Name: "Product 2",
		},
		{
			Id: 3,
			Name: "Product 3",
		},
	})
}
