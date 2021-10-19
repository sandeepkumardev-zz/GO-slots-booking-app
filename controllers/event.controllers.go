package controllers

import (
	"net/http"
	M "slot/middleware"
	"slot/models"
	"slot/services"

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
		ctx.JSON(http.StatusNotFound, &M.ResponseTransformer{Message: err, Result: nil, IsSuccess: false})
		return
	}

	ctx.JSON(http.StatusOK, &M.ResponseTransformer{Message: "Successfully fetched list of events.", Result: res, IsSuccess: true})
}

func (EventController) BookedSlots(ctx *gin.Context) {
	var zone struct {
		TimeZone string `json:"timezone"`
	}
	ctx.ShouldBindJSON(&zone)

	res, err := services.BookedSlots(zone.TimeZone)
	if err != "" {
		ctx.JSON(http.StatusNotFound, &M.ResponseTransformer{Message: err, Result: nil, IsSuccess: false})
		return
	}

	ctx.JSON(http.StatusOK, &M.ResponseTransformer{Message: "Successfully fetched list of events.", Result: res, IsSuccess: true})
}

func (EventController) CreateEvent(ctx *gin.Context) {
	var event models.Event

	if inpErr := ctx.ShouldBind(&event); inpErr != nil {
		ctx.JSON(http.StatusUnprocessableEntity, &M.ResponseTransformer{Message: "Invalid input provided.", Result: nil, IsSuccess: false})
		return
	}

	res, err := services.CreateEvent(&event)
	if err != "" {
		ctx.JSON(http.StatusBadRequest, &M.ResponseTransformer{Message: err, Result: nil, IsSuccess: false})
		return
	}

	ctx.JSON(http.StatusCreated, &M.ResponseTransformer{Message: "Successfully created event.", Result: res, IsSuccess: true})
}

func (EventController) GetOneEvent(ctx *gin.Context) {
	id := ctx.Params[0].Value

	res, err := services.GetOneEvent(id)
	if err != "" {
		ctx.JSON(http.StatusBadRequest, &M.ResponseTransformer{Message: err, Result: nil, IsSuccess: false})
		return
	}

	ctx.JSON(http.StatusAccepted, &M.ResponseTransformer{Message: "Successfully fetched event.", Result: res.Value, IsSuccess: true})
}

func (EventController) UpdateEvent(ctx *gin.Context) {
	id := ctx.Params[0].Value
	var event models.Event

	if inpErr := ctx.ShouldBindJSON(&event); inpErr != nil {
		ctx.JSON(http.StatusUnprocessableEntity, &M.ResponseTransformer{Message: "Invalid input provided.", Result: nil, IsSuccess: false})
		return
	}

	msg, err := services.UpdateEvent(id, &event)

	if err != "" {
		ctx.JSON(http.StatusBadRequest, &M.ResponseTransformer{Message: err, Result: nil, IsSuccess: false})
		return
	}

	ctx.JSON(http.StatusCreated, &M.ResponseTransformer{Message: msg, Result: nil, IsSuccess: true})
}

func (EventController) DeleteEvent(ctx *gin.Context) {
	id := ctx.Params[0].Value
	var event models.Event

	msg, err := services.DeleteEvent(id, &event)
	if err != "" {
		ctx.JSON(http.StatusBadRequest, &M.ResponseTransformer{Message: err, Result: nil, IsSuccess: false})
		return
	}

	ctx.JSON(http.StatusOK, &M.ResponseTransformer{Message: msg, Result: nil, IsSuccess: true})
}

func (EventController) UploadFile(ctx *gin.Context) {
	id := ctx.Params[0].Value

	msg, err := services.UploadFile(id, ctx)
	if err != "" {
		ctx.JSON(http.StatusNotFound, &M.ResponseTransformer{Message: err, Result: nil, IsSuccess: true})
		return
	}

	ctx.JSON(http.StatusOK, &M.ResponseTransformer{Message: msg, Result: nil, IsSuccess: true})
}
