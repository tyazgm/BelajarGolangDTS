package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	StartServer()
}

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/cars", CreateCar)
	router.PUT("/cars/:carID", UpdateCar)
	router.GET("/cars/", GetAllCar)
	router.GET("/cars/:carID", GetCarByID)
	router.DELETE("/cars/:carID", DeleteCar)

	router.Run(":8080")

	return router
}

type Car struct {
	CarID string
	Brand string
	Model string
	Price int
}

var CarDatas = []Car{}

func CreateCar(ctx *gin.Context) {
	var newCar Car

	err := ctx.ShouldBindJSON(&newCar)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newCar.CarID = fmt.Sprintf("c%d", len(CarDatas)+1)
	CarDatas = append(CarDatas, newCar)

	ctx.JSON(http.StatusOK, CarDatas)
}

func UpdateCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	found := false
	var updatedCar Car

	err := ctx.ShouldBindJSON(&updatedCar)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, car := range CarDatas {
		if carID == car.CarID {
			found = true
			CarDatas[i] = updatedCar
			CarDatas[i].CarID = carID
			break
		}
	}

	if !found {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("car with id %v not found", carID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":      fmt.Sprintf("car with id %v has been succesfully uptaded", carID),
		"carDatas now": CarDatas,
	})
}

func GetAllCar(ctx *gin.Context) {
	carData := CarDatas

	ctx.JSON(http.StatusOK, carData)
}

func GetCarByID(ctx *gin.Context) {
	carID := ctx.Param("carID")
	carData := CarDatas
	found := false

	for i, car := range carData {
		if car.CarID == carID {
			found = true
			ctx.JSON(http.StatusOK, carData[i])
			break
		}
	}

	if !found {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("car with id %v not found", carID),
		})
		return
	}
}

func DeleteCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	found := false
	var carIdx int

	for i, car := range CarDatas {
		if car.CarID == carID {
			found = true
			carIdx = i
			break
		}
	}

	if !found {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("car with id %v not found", carID),
		})
		return
	}

	copy(CarDatas[carIdx:], CarDatas[carIdx+1:])
	CarDatas[len(CarDatas)-1] = Car{}
	CarDatas = CarDatas[:len(CarDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("car with id %v hasx been succesfully removed", carID),
	})
}
