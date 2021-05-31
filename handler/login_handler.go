package handler

import (
	"github.com/HiBang15/jwt-golang/model"
	"github.com/HiBang15/jwt-golang/token"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func LoginHandler(context *gin.Context) {
	var loginObj model.LoginRequest

	if err := context.ShouldBindJSON(&loginObj); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
			Status:  http.StatusBadRequest,
			Message: "Cannot Login!",
			Error:   "Username or password fail with err",
		})
		return
	}

	var claims = &model.JwtClaims{}
	claims.CompanyId = "CompanyId"
	claims.Username = loginObj.Username
	claims.Roles = []int{1, 2, 3}
	claims.Audience = context.Request.Header.Get("Referer")

	var tokenCreationTime = time.Now().UTC()
	var expirationTime = tokenCreationTime.Add(time.Duration(10) + time.Minute)
	tokenString, err := token.GenerateToken(claims, expirationTime)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
			Status:  http.StatusBadRequest,
			Message: "Cannot Login!",
			Error:   "Username or password fail with err",
		})
		return
	}

	context.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
		Status:  http.StatusOK,
		Message: "Ok!",
		Data:    tokenString,
	})
}
