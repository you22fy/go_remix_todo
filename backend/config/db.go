package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB
var err error

func DatabaseInit() {
	host := "localhost"
	user := "root"
	password := "password"
	dbName := "app"
	port := "5432"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbName=%s port=%s", host, user, password, dbName, port)
	database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

func DB() *gorm.DB {
	return database
}
