package controller

import (
	"latihan/model"
	"latihan/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService service.OrderService
}

func NewOrderController(orderService service.OrderService) *OrderController {
	return &OrderController{
		orderService: orderService,
	}
}

func (oc *OrderController) CreateOrder(ctx *gin.Context) {
	var request model.OrderCreateRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: err.Error(),
		})
		return
	}

	userID, isExist := ctx.Get("user_id")
	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: model.ErrorInvalidToken.Err,
		})
		return
	}

	order, err := oc.orderService.Create(request, userID.(string))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

func (oc *OrderController) GetListOrder(ctx *gin.Context) {

}
func (oc *OrderController) GetOrder(ctx *gin.Context) {

}
