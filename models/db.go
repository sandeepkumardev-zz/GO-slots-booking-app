package models

import (
	"fmt"
	"slot/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var err error

func ConnectToDb() error {
	slotConfig, userConfig := config.SlotDBConfig()

	config.SlotDB, err = gorm.Open("mysql", config.DbUrl(slotConfig))
	config.UserDB, err = gorm.Open("mysql", config.DbUrl(userConfig))

	if err != nil {
		fmt.Printf("Status: %v\n", err)
		return err
	}

	config.SlotDB.AutoMigrate(&Event{})
	config.SlotDB.AutoMigrate(&File{})
	//2nd db
	config.UserDB.AutoMigrate(&User{})

	fmt.Println("Database connected!")
	return nil
}
