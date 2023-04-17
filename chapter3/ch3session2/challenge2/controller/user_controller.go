package controller

import (
	"challenge2/model"
	"challenge2/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(ctx *gin.Context) {

	var request model.UserRegisterRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MyError{
			Err: err.Error(),
		})
		return
	}

	id, err := uc.userService.Register(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MyError{
			Err: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.UserRegisterResponse{
		ID: id,
	})
}

func (uc *UserController) Login(ctx *gin.Context) {
	var request model.UserLoginRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MyError{
			Err: err.Error(),
		})
		return
	}

	token, err := uc.userService.Login(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MyError{
			Err: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.UserLoginResponse{
		Token: token,
	})
}

// func (oc *OrderController) GetListOrder(ctx *gin.Context) {

// }
// func (oc *OrderController) GetOrder(ctx *gin.Context) {

// }
