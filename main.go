package main

import (
	"fmt"
	db "slot/models"
	"slot/routes"
	"time"

	"github.com/gin-contrib/cors"
)

func main() {
	fmt.Println("Mini Slot Booking Project")

	db.ConnectToDb()

	router := routes.RouterSetup()
	err := router.Run(":3000")

	// cors
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.MaxAge = 12 * time.Hour
	router.Use(cors.New(config))

	if err != nil {
		fmt.Println("Something went wrong with the router!")
	}
}
