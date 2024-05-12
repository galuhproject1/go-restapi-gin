package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(postgres.Open("host=localhost user=postgres password=123456 dbname=trygoapi port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Product{})

	DB = database
}
