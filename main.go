package main

import (
	"fmt"
	db "slot/models"
	"slot/routes"
)

func main() {
	fmt.Println("Mini Slot Booking Project")

	db.ConnectToDb()

	router := routes.RouterSetup()
	err := router.Run(":3000")

	if err != nil {
		fmt.Println("Something went wrong with the router!")
	}
}
