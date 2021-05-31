package middleware

import (
	"github.com/HiBang15/jwt-golang/model"
	"github.com/HiBang15/jwt-golang/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReturnUnauthorized(context *gin.Context) {
	context.AbortWithStatusJSON(http.StatusUnauthorized, model.Response{
		Status:  http.StatusUnauthorized,
		Message: "Unauthorized access!",
		Error:   "You are not authorized to access this path",
	})
}


func ValidateToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.Request.Header.Get("apiKey")
		referer := context.Request.Header.Get("Referer")
		valid, claims := token.VerifyToken(tokenString, referer)
		if !valid {
			ReturnUnauthorized(context)
		}
		if len(context.Keys) == 0 {
			context.Keys = make(map[string]interface{})
		}
		
		context.Keys["CompanyId"] = claims.CompanyId
		context.Keys["Username"] = claims.Username
		context.Keys["Roles"] = claims.Roles
	}
}