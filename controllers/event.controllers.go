package controllers

import (
	"net/http"
	"slot/models"
	"slot/services"

	"github.com/gin-gonic/gin"
)

func AvailableSlots(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Get available slots")
}

func BookedSlots(ctx *gin.Context) {
	res, err := services.BookedSlots()
	if err != "" {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, res.Value)
}

func CreateEvent(ctx *gin.Context) {
	var event models.Event

	if inpErr := ctx.ShouldBind(&event); inpErr != nil {
		ctx.JSON(http.StatusUnprocessableEntity, "Invalid input provided")
		return
	}

	// if err := utils.CreateValidator(&event); err != "" {
	// 	ctx.JSON(http.StatusUnprocessableEntity, err)
	// 	return
	// }

	res, err := services.CreateEvent(&event)
	if err != "" {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

func GetOneEvent(ctx *gin.Context) {
	id := ctx.Params[0].Value

	res, err := services.GetOneEvent(id)
	if err != "" {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusAccepted, res.Value)
}

func UpdateEvent(ctx *gin.Context) {
	id := ctx.Params[0].Value
	var event models.Event

	if inpErr := ctx.ShouldBindJSON(&event); inpErr != nil {
		ctx.JSON(http.StatusUnprocessableEntity, "Invalid input provided")
		return
	}

	res, err := services.UpdateEvent(id, &event)

	if err != "" {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

func DeleteEvent(ctx *gin.Context) {
	id := ctx.Params[0].Value
	var event models.Event

	res, err := services.DeleteEvent(id, &event)
	if err != "" {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
