package models

import (
	"fmt"
	"slot/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var err error

func ConnectToDb() {
	config.DB, err = gorm.Open("mysql", config.DbUrl(config.BuildDBConfig()))

	if err != nil {
		fmt.Printf("Status: %v\n", err)
		return
	}

	config.DB.AutoMigrate(&Event{})

	fmt.Println("Database connected!")
}
