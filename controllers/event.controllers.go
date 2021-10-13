package controllers

import (
	"net/http"
	"slot/models"
	"slot/services"
	"slot/utils"

	"github.com/gin-gonic/gin"
)

type EventController struct{}

func (EventController) AvailableSlots(ctx *gin.Context) {
	var Query struct {
		Date     string `json:"date"`
		TimeZone string `json:"timezone"`
	}
	ctx.ShouldBindJSON(&Query)

	res, err := services.AvailableSlots(Query.Date, Query.TimeZone)
	if err != "" {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (EventController) BookedSlots(ctx *gin.Context) {
	var zone struct {
		TimeZone string `json:"timezone"`
	}
	ctx.ShouldBindJSON(&zone)

	res, err := services.BookedSlots(zone.TimeZone)
	if err != "" {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (EventController) CreateEvent(ctx *gin.Context) {
	var event models.Event

	if inpErr := ctx.ShouldBind(&event); inpErr != nil {
		ctx.JSON(http.StatusUnprocessableEntity, "Invalid input provided")
		return
	}

	//check duration limit
	if err := utils.CreateValidator(&event); err != "" {
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	res, err := services.CreateEvent(&event)
	if err != "" {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

func (EventController) GetOneEvent(ctx *gin.Context) {
	id := ctx.Params[0].Value

	res, err := services.GetOneEvent(id)
	if err != "" {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusAccepted, res.Value)
}

func (EventController) UpdateEvent(ctx *gin.Context) {
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

func (EventController) DeleteEvent(ctx *gin.Context) {
	id := ctx.Params[0].Value
	var event models.Event

	res, err := services.DeleteEvent(id, &event)
	if err != "" {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
