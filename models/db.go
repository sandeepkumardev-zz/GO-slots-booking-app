package models

import (
	"fmt"
	"slot/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var err error

func ConnectToDb() {
	slotConfig, fileConfig := config.SlotDBConfig()

	config.SlotDB, err = gorm.Open("mysql", config.DbUrl(slotConfig))
	config.FileDB, err = gorm.Open("mysql", config.DbUrl(fileConfig))

	if err != nil {
		fmt.Printf("Status: %v\n", err)
		return
	}

	config.SlotDB.AutoMigrate(&Event{})
	config.FileDB.AutoMigrate(&File{})

	fmt.Println("Database connected!")
}
