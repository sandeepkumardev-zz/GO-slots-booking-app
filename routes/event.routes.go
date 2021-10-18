package routes

import (
	"net/http"
	ctrl "slot/controllers"

	"github.com/gin-gonic/gin"
)

func RouterSetup() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Events booking app!")
	})

	router.GET("/slots", ctrl.EventController{}.AvailableSlots)
	router.GET("/events", ctrl.EventController{}.BookedSlots)
	router.POST("/event", ctrl.EventController{}.CreateEvent)
	router.GET("/event/:eventId", ctrl.EventController{}.GetOneEvent)
	router.PUT("/event/:eventId", ctrl.EventController{}.UpdateEvent)
	router.DELETE("/event/:eventId", ctrl.EventController{}.DeleteEvent)

	router.POST("/upload/:eventId", ctrl.EventController{}.UploadFile)

	return router
}
