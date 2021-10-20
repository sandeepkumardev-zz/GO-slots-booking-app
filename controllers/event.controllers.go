package controllers

import (
	"net/http"
	M "slot/middleware"
	"slot/models"
	"slot/services"
	V "slot/validators"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EventController struct{}

func (EventController) AvailableSlots(ctx *gin.Context) {
	var Query struct {
		Date     string `json:"date"`
		TimeZone string `json:"timezone"`
	}
	ctx.ShouldBindJSON(&Query)

	res := services.AvailableSlots(Query.Date, Query.TimeZone)
	if !res.Success {
		ctx.JSON(http.StatusBadRequest, &M.ResponseTransformer{Message: res.Message, Data: res.Data, Success: res.Success})
		return
	}

	ctx.JSON(http.StatusOK, &M.ResponseTransformer{Message: res.Message, Data: res.Data, Success: res.Success})
}

func (EventController) BookedSlots(ctx *gin.Context) {
	var zone struct {
		TimeZone string `json:"timezone"`
	}
	ctx.ShouldBindJSON(&zone)

	res := services.BookedSlots(zone.TimeZone)
	if !res.Success {
		ctx.JSON(http.StatusBadRequest, &M.ResponseTransformer{Message: res.Message, Data: res.Data, Success: res.Success})
		return
	}

	ctx.JSON(http.StatusOK, &M.ResponseTransformer{Message: res.Message, Data: res.Data, Success: res.Success})
}

func (EventController) CreateEvent(ctx *gin.Context) {
	var event models.Event

	if inpErr := ctx.ShouldBind(&event); inpErr != nil {
		ctx.JSON(http.StatusUnprocessableEntity, &M.ResponseTransformer{Message: "Invalid input provided.", Data: nil, Success: false})
		return
	}

	//validation
	errors := V.Validation{}.IsValid(&event)
	if errors != nil {
		ctx.JSON(http.StatusBadRequest, &M.ResponseTransformer{Message: "Validation error!", Data: errors, Success: false})
		return
	}

	res := services.CreateEvent(&event)
	if !res.Success {
		ctx.JSON(http.StatusBadRequest, &M.ResponseTransformer{Message: res.Message, Data: res.Data, Success: res.Success})
		return
	}

	ctx.JSON(http.StatusCreated, &M.ResponseTransformer{Message: res.Message, Data: res.Data, Success: res.Success})
}

func (EventController) GetOneEvent(ctx *gin.Context) {
	id := ctx.Params[0].Value

	res := services.GetOneEvent(id)
	if !res.Success {
		ctx.JSON(http.StatusBadRequest, &M.ResponseTransformer{Message: res.Message, Data: res.Data, Success: res.Success})
		return
	}

	ctx.JSON(http.StatusOK, &M.ResponseTransformer{Message: res.Message, Data: res.Data, Success: res.Success})
}

func (EventController) UpdateEvent(ctx *gin.Context) {
	//convert params id to int
	id, _ := strconv.Atoi(ctx.Params[0].Value)
	var event models.Event

	if inpErr := ctx.ShouldBindJSON(&event); inpErr != nil {
		ctx.JSON(http.StatusUnprocessableEntity, &M.ResponseTransformer{Message: "Invalid input provided.", Data: nil, Success: false})
		return
	}

	res := services.UpdateEvent(id, &event)

	if !res.Success {
		ctx.JSON(http.StatusBadRequest, &M.ResponseTransformer{Message: res.Message, Data: res.Data, Success: res.Success})
		return
	}

	ctx.JSON(http.StatusCreated, &M.ResponseTransformer{Message: res.Message, Data: res.Data, Success: res.Success})
}

func (EventController) DeleteEvent(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Params[0].Value)
	var event models.Event

	res := services.DeleteEvent(id, &event)
	if !res.Success {
		ctx.JSON(http.StatusBadRequest, &M.ResponseTransformer{Message: res.Message, Data: res.Data, Success: res.Success})
		return
	}

	ctx.JSON(http.StatusOK, &M.ResponseTransformer{Message: res.Message, Data: res.Data, Success: res.Success})
}

func (EventController) UploadFile(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Params[0].Value)

	res := services.UploadFile(id, ctx)
	if !res.Success {
		ctx.JSON(http.StatusBadRequest, &M.ResponseTransformer{Message: res.Message, Data: res.Data, Success: res.Success})
		return
	}

	ctx.JSON(http.StatusOK, &M.ResponseTransformer{Message: res.Message, Data: res.Data, Success: res.Success})
}
