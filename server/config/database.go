package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConncectToDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
