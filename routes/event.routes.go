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

	router.GET("/slots", ctrl.AvailableSlots)
	router.GET("/events", ctrl.BookedSlots)
	router.POST("/event", ctrl.CreateEvent)
	router.GET("/event/:eventId", ctrl.GetOneEvent)
	router.PUT("/event/:eventId", ctrl.UpdateEvent)
	router.DELETE("/event/:eventId", ctrl.DeleteEvent)

	return router
}
