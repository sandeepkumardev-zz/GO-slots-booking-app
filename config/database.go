package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var SlotDB *gorm.DB
var FileDB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func SlotDBConfig() (*DBConfig, *DBConfig) {
	slotDbConfig := DBConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		Password: "12345678",
		DBName:   "slot_booking_app",
	}
	fileDbConfig := DBConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		Password: "12345678",
		DBName:   "file_uploding_app",
	}

	return &slotDbConfig, &fileDbConfig
}

func DbUrl(db *DBConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", db.User, db.Password, db.Host, db.Port, db.DBName)
}
