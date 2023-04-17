package middleware

import (
	"latihan/helper"
	"latihan/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(ctx *gin.Context) {
	auth := ctx.GetHeader("Authorization")

	token := strings.Split(auth, " ")[1]

	jwtToken, err := helper.VerifyToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MyError{
			Err: err.Error(),
		})
		return
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: err.Error(),
		})
		return
	}

	ctx.Set("user_id", claims["user_id"])
	ctx.Next()
}
