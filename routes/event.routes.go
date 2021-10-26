package routes

import (
	"net/http"
	ctrl "slot/controllers"
	M "slot/middleware"

	"github.com/gin-gonic/gin"
)

func RouterSetup() *gin.Engine {
	router := gin.Default()

	router.Use(M.Cors())
	router.Use(M.LogMiddleware)
	router.Use(M.RateLimit())

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Events booking app!")
	})

	secure := router.Group("/auth", M.EnsureLoggedIn())
	{
		secure.GET("/", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "Secure Page!")
		})
	}

	router.GET("/slots", ctrl.EventController{}.AvailableSlots)
	router.GET("/events", ctrl.EventController{}.BookedSlots)
	router.POST("/event", ctrl.EventController{}.CreateEvent)
	router.GET("/event/:eventId", ctrl.EventController{}.GetOneEvent)
	router.PUT("/event/:eventId", ctrl.EventController{}.UpdateEvent)
	router.DELETE("/event/:eventId", ctrl.EventController{}.DeleteEvent)

	router.POST("/upload/:Id", ctrl.EventController{}.UploadFile)

	return router
}
